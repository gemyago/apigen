package org.gemyago.apigen.go.server;

import java.io.File;
import java.io.FileFilter;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.ArrayList;
import java.util.Iterator;
import java.util.List;
import java.util.Locale;
import java.util.Map;
import java.util.Set;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

import org.apache.commons.lang3.StringUtils;
import org.openapitools.codegen.*;
import org.openapitools.codegen.config.GlobalSettings;
import org.openapitools.codegen.languages.*;
import org.openapitools.codegen.model.ModelMap;
import org.openapitools.codegen.model.OperationMap;
import org.openapitools.codegen.model.OperationsMap;
import org.openapitools.codegen.templating.mustache.IndentedLambda;
import org.openapitools.codegen.utils.CamelizeOption;
import org.openapitools.codegen.utils.ModelUtils;

import com.google.common.base.CaseFormat;
import com.google.common.collect.ImmutableMap;
import com.samskivert.mustache.Mustache;

import io.swagger.v3.oas.models.OpenAPI;
import io.swagger.v3.oas.models.Operation;
import io.swagger.v3.oas.models.media.Schema;
import io.swagger.v3.oas.models.parameters.Parameter;
import io.swagger.v3.oas.models.parameters.RequestBody;
import io.swagger.v3.oas.models.servers.Server;

import static org.openapitools.codegen.utils.StringUtils.*;
import static org.openapitools.codegen.utils.CamelizeOption.LOWERCASE_FIRST_LETTER;

public class GoApigenServerGenerator extends AbstractGoCodegen {

  // source folder where to write the files
  protected String sourceFolder = "src";
  protected String apiVersion = "1.0.0";

  private Boolean modelsOnly = false;
  private Boolean apisOnly = false;

  /**
   * Configures the type of generator.
   *
   * @return the CodegenType for this generator
   * @see org.openapitools.codegen.CodegenType
   */
  public CodegenType getTag() {
    return CodegenType.SERVER;
  }

  /**
   * Configures a friendly name for the generator. This will be used by the
   * generator
   * to select the library with the -g flag.
   *
   * @return the friendly name for the generator
   */
  public String getName() {
    return "go-apigen-server";
  }

  /**
   * Returns human-friendly help for the generator. Provide the consumer with help
   * tips, parameters here
   *
   * @return A string value for the help message
   */
  public String getHelp() {
    return "Generates a go-apigen-server server library.";
  }

  public GoApigenServerGenerator() {
    super();

    apiNameSuffix = "";
    apiNamePrefix = "";

    // set the output folder here
    outputFolder = "generated-code/go-apigen-server";

    boolean generateModels = GlobalSettings.getProperty(CodegenConstants.MODELS) != null;
    boolean generateApis = GlobalSettings.getProperty(CodegenConstants.APIS) != null;
    boolean generateSupportingFiles = GlobalSettings.getProperty(CodegenConstants.SUPPORTING_FILES) != null;

    // if all are false, this means no options are provided (or all three), so we
    // consider all as true
    if (!generateModels && !generateApis && !generateSupportingFiles) {
      generateModels = true;
      generateApis = true;
      generateSupportingFiles = true;
    }

    // supporting files must always be enable
    modelsOnly = generateModels && generateSupportingFiles && !generateApis;
    apisOnly = generateApis && generateSupportingFiles && !generateModels;

    // We have to generate validators for models in APIs only mode so we have to trick generator to do that
    // and below is the way to do it for now.
    if (apisOnly) {
      GlobalSettings.setProperty(CodegenConstants.MODELS, "");
    }

    if (generateModels) {
      /**
       * Models. You can write model files using the modelTemplateFiles map.
       * if you want to create one template for file, you can do so here.
       * for multiple files for model, just put another entry in the
       * `modelTemplateFiles` with
       * a different extension
       */
      modelTemplateFiles.put(
          "model.mustache", // the template to use
          ".go"); // the extension for each file to write
    }

    // We only generate models validation if we're generating apis
    if (generateApis) {
      modelTemplateFiles.put(
          "model_validation.mustache", // the template to use
          "_validation.go"); // the extension for each file to write
      templateOutputDirs.put("model_validation.mustache", "internal");
    }

    /**
     * Api classes. You can write classes for each Api file with the
     * apiTemplateFiles map.
     * as with models, add multiple entries with different extensions for multiple
     * files per
     * class
     */
    apiTemplateFiles.put(
        "controller.mustache",
        "_controller.go");
    apiTemplateFiles.put(
        "controller_params.mustache",
        "_params.go");

    /**
     * Template Location. This is the location which templates will be read from.
     * The generator
     * will use the resource stream to attempt to read the templates.
     */
    templateDir = "go-apigen-server";

    /**
     * Api Package. Optional, if needed, this can be used in templates
     */
    apiPackage = "handlers";

    /**
     * Model Package. Optional, if needed, this can be used in templates.
     */
    modelPackage = "models";

    /**
     * Additional Properties. These values can be passed to the templates and
     * are available in models, apis, and supporting files
     */
    additionalProperties.put("apiVersion", apiVersion);

    // additionalProperties.put(CodegenConstants.GENERATE_ALIAS_AS_MODEL, "true");

    // Common handlers related stuff is only generated if apis are generated
    if (generateApis) {
      /**
       * Supporting Files. You can write single files for the generator with the
       * entire object tree available. If the input file has a suffix of `.mustache
       * it will be processed by the template engine. Otherwise, it will be copied
       */
      supportingFiles.add(new SupportingFile(
          "handlers.mustache",
          apiPackage,
          "handlers.go"));

      supportingFiles.add(new SupportingFile(
          "validators.mustache",
          "internal",
          "validators.go"));
    }

    typeMapping.put("date", "time.Time");
  }

  @Override
  public String modelFileFolder() {
    // In modelsOnly mode we're generating to the root folder to keep it flat
    if (modelsOnly) {
      return outputFolder;
    }
    return super.modelFileFolder();
  }

  @Override
  public void preprocessOpenAPI(OpenAPI openAPI) {
    super.preprocessOpenAPI(openAPI);

    // We are generating request parameters as model with all parameters as fields
    // so we need to "simulate" and inject them to schemas
    List<Operation> allOperations = openAPI.getPaths().values().stream()
        .flatMap(path -> path.readOperations().stream()).toList();

    for (Operation operation : allOperations) {
      Schema<?> parametersModel = new Schema<>()
          .type("object")
          .description("Parameters for the " + operation.getOperationId() + " operation.");
      parametersModel.addExtension("x-apigen-operation-params-model", true);

      if (operation.getParameters() != null) {
        for (Parameter parameter : operation.getParameters()) {
          parametersModel.addProperty(parameter.getName(), parameter.getSchema());
          if (parameter.getRequired() != null && parameter.getRequired()) {
            parametersModel.addRequiredItem(parameter.getName());
          }
        }
      }

      if (operation.getRequestBody() != null) {
        RequestBody requestBody = operation.getRequestBody();
        Schema<?> bodySchema = requestBody.getContent().values().iterator().next().getSchema();
        parametersModel.addProperty("payload", bodySchema);
        if (requestBody.getRequired() != null) {
          parametersModel.addRequiredItem("payload");
        }
      }

      if (parametersModel.getProperties() == null || parametersModel.getProperties().isEmpty()) {
        continue;
      }

      String modelName = camelize(operation.getOperationId(), CamelizeOption.UPPERCASE_FIRST_CHAR) + "Params";
      openAPI.getComponents().getSchemas().put(modelName, parametersModel);
      operation.addExtension("x-apigen-params-model", modelName);
    }
  }

  @Override
  public String toApiFilename(String name) {
    return super.toApiFilename(name).replace("api_", "");
  }

  @Override
  public String toModelFilename(String name) {
    return super.toModelFilename(name).replace("model_", "");
  }

  private File resolveGoMod(File folder) {
    if (folder == null) {
      return null;
    }

    File[] files = folder.listFiles(new FileFilter() {
      @Override
      public boolean accept(File pathname) {
        return pathname.getName().equals("go.mod");
      }
    });

    if (files != null && files.length == 1) {
      return files[0];
    }

    return resolveGoMod(folder.getParentFile());
  }

  private String extractModule(Path filePath) {
    Pattern pattern = Pattern.compile("module (?<module>.*)", Pattern.MULTILINE);
    String contents;
    try {
      contents = Files.readString(filePath);
    } catch (IOException e) {
      throw new RuntimeException(e);
    }
    Matcher matcher = pattern.matcher(contents);
    if (matcher.find()) {
      return matcher.group("module");
    }
    throw new RuntimeException("Could not find module in go.mod file");
  }

  @Override
  public void processOpts() {
    if (!additionalProperties.containsKey(CodegenConstants.ENUM_CLASS_PREFIX)) {
      additionalProperties.put(CodegenConstants.ENUM_CLASS_PREFIX, true);
    }

    super.processOpts();

    File outputFolderFile = new File(outputFolder);
    Path outputFolderFilePath = outputFolderFile.toPath();
    File goModFile = resolveGoMod(outputFolderFile);
    if (goModFile == null) {
      throw new RuntimeException("Can not find go.mod in a project hierarchy");
    }
    Path goModFilePath = goModFile.toPath();
    String moduleName = extractModule(goModFilePath);
    String invokerPackage = moduleName + "/" + goModFilePath.getParent().relativize(outputFolderFilePath).toString();
    additionalProperties.put(
        CodegenConstants.INVOKER_PACKAGE,
        invokerPackage);

    // In apisOnly it is expected that the caller provides model package in a full form
    additionalProperties.put(
        "modelFullPackage",
        apisOnly ? modelPackage() : invokerPackage + "/" + modelPackage());

    if (additionalProperties.containsKey("generatedCodeComment")) {
      if (StringUtils.isEmpty(additionalProperties.get("generatedCodeComment").toString())) {
        additionalProperties.remove("generatedCodeComment");
      }
    } else {
      additionalProperties.put("generatedCodeComment", "Code generated by apigen DO NOT EDIT");
    }
  }

  @Override
  protected ImmutableMap.Builder<String, Mustache.Lambda> addMustacheLambdas() {
    ImmutableMap.Builder<String, Mustache.Lambda> lambdas = super.addMustacheLambdas();
    lambdas.put("tab_indented_2", new IndentedLambda(1, "\t", "", false, false));
    lambdas.put("tab_indented_3", new IndentedLambda(3, "\t", "", false, false));
    lambdas.put("tab_indented_4", new IndentedLambda(4, "\t", "", false, false));
    return lambdas;
  }

  @Override
  public CodegenParameter fromParameter(Parameter parameter, Set<String> imports) {
    CodegenParameter codegenParameter = super.fromParameter(parameter, imports);

    // fromParameter will not dereference the schema to set some properties
    Schema<?> schema = ModelUtils.getReferencedSchema(this.openAPI, parameter.getSchema());
    if (schema.getNullable() != null) {
      codegenParameter.isNullable = schema.getNullable();
    }

    appendParamVendorExtensions(codegenParameter, parameter.getIn());
    return codegenParameter;
  }

  @Override
  public CodegenParameter fromRequestBody(RequestBody body, Set<String> imports, String bodyParameterName) {
    if (StringUtils.isEmpty(bodyParameterName)) {
      bodyParameterName = "payload";
    }

    CodegenParameter codegenParameter = super.fromRequestBody(body, imports, bodyParameterName);
    appendParamVendorExtensions(codegenParameter, "body");

    codegenParameter.nameInCamelCase = camelize(codegenParameter.paramName, LOWERCASE_FIRST_LETTER);
    codegenParameter.nameInPascalCase = camelize(codegenParameter.paramName);
    codegenParameter.nameInSnakeCase = CaseFormat.UPPER_CAMEL.to(CaseFormat.UPPER_UNDERSCORE,
        codegenParameter.nameInPascalCase);
    codegenParameter.nameInLowerCase = codegenParameter.paramName.toLowerCase(Locale.ROOT);
    return codegenParameter;
  }

  @Override
  public String toRegularExpression(String pattern) {
    if (StringUtils.isEmpty(pattern)) {
      return pattern;
    }
    // Pattern is considered platform specific, so we're using as it (escaping only)
    return pattern.replace("\\", "\\\\");
  }

  @Override
  public void postProcessModelProperty(CodegenModel model, CodegenProperty property) {
    super.postProcessModelProperty(model, property);

    if (property.isEnum) {
      property.datatypeWithEnum = model.classname + property.nameInPascalCase;
    }
  }

  @Override
  public OperationsMap postProcessOperationsWithModels(OperationsMap objs, List<ModelMap> allModels) {
    OperationsMap operationsMap = super.postProcessOperationsWithModels(objs, allModels);
    operationsMap.put("operations.className", objs.getOperations().get("classname"));

    OperationMap operations = operationsMap.getOperations();
    for (CodegenOperation op : operations.getOperation()) {
      for (CodegenParameter param : op.allParams) {
        if (param.isEnum) {
          param.datatypeWithEnum = op.operationId + "Params"
              + camelize(param.datatypeWithEnum.replace("_", "-").toLowerCase());
        }
      }
    }

    List<Map<String, String>> imports = objs.getImports();
    if (imports == null)
      return objs;

    operationsMap.put("hasImportedModel", false);
    Iterator<Map<String, String>> iterator = imports.iterator();
    List<String> importedModels = new ArrayList<>();
    while (iterator.hasNext()) {
      String _import = iterator.next().get("import");
      int modelPkgStartsAt = _import.indexOf(modelPackage());
      if (modelPkgStartsAt >= 0) {
        operationsMap.put("hasImportedModel", true);
        importedModels.add(_import.substring(modelPkgStartsAt + modelPackage().length() + 1));
      }
    }
    operationsMap.put("importedModels", importedModels);

    return operationsMap;
  }

  @Override
  public CodegenOperation fromOperation(String path, String httpMethod, Operation operation, List<Server> servers) {
    CodegenOperation op = super.fromOperation(path, httpMethod, operation, servers);
    if (op.hasParams || !StringUtils.isEmpty(op.returnType)) {
      op.vendorExtensions.put("x-apigen-has-params-or-return", true);
    }
    return op;
  }

  private String titleCase(final String input) {
    if (input == null || "".equals(input)) {
      return "";
    }

    String firstLetter = input.substring(0, 1).toUpperCase(Locale.ROOT);
    if (input.length() == 1) {
      return firstLetter;
    }
    return firstLetter + input.substring(1);
  }

  private void appendParamVendorExtensions(CodegenParameter codegenParameter, String location) {
    codegenParameter.vendorExtensions.put("x-apigen-param-location", location);
    codegenParameter.vendorExtensions.put("x-apigen-param-location-tc", titleCase(location));
  }

  @Override
  public String sanitizeName(String name, String removeCharRegEx, ArrayList<String> exceptionList) {
    String result = super.sanitizeName(name, removeCharRegEx, exceptionList);

    // this is one of the few places where we can normalize abbreviations globally
    return normalizeAbbreviations(result);
  }

  @Override
  public String toModelName(String name) {
    return normalizeAbbreviations(super.toModelName(name));
  }

  @Override
  public String toVarName(String name) {
    return normalizeAbbreviations(super.toVarName(name));
  }

  @Override
  public String toApiName(String name) {
    return normalizeAbbreviations(super.toApiName(name));
  }

  private String normalizeAbbreviations(String name) {
    /**
     * Golang has a convention to upper case abbreviations in names. Most typical
     * case is "ID" instead of "Id".
     * This is the only place that we can do this transformation globally. We may
     * need to parametrize this in the future.
     */

    if (name.contains("Id")) {
      return name.replace("Id", "ID");
    }
    return name;
  }
}
