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
import org.openapitools.codegen.languages.*;
import org.openapitools.codegen.model.ModelMap;
import org.openapitools.codegen.model.OperationsMap;
import org.openapitools.codegen.templating.mustache.IndentedLambda;

import com.google.common.base.CaseFormat;
import com.google.common.collect.ImmutableMap;
import com.samskivert.mustache.Mustache;

import io.swagger.v3.oas.models.parameters.Parameter;
import io.swagger.v3.oas.models.parameters.RequestBody;

import static org.openapitools.codegen.utils.StringUtils.*;
import static org.openapitools.codegen.utils.CamelizeOption.LOWERCASE_FIRST_LETTER;

public class GoApigenServerGenerator extends AbstractGoCodegen {

  // source folder where to write the files
  protected String sourceFolder = "src";
  protected String apiVersion = "1.0.0";

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
    modelTemplateFiles.put(
        "model_validation.mustache", // the template to use
        "_validation.go"); // the extension for each file to write
    templateOutputDirs.put("model_validation.mustache", "internal");

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
     * Model Package. Optional, if needed, this can be used in templates
     */
    modelPackage = "models";

    /**
     * Additional Properties. These values can be passed to the templates and
     * are available in models, apis, and supporting files
     */
    additionalProperties.put("apiVersion", apiVersion);
    additionalProperties.put("invokerPackage", "Hello World");

    // additionalProperties.put(CodegenConstants.GENERATE_ALIAS_AS_MODEL, "true");

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

    typeMapping.put("date", "time.Time");
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

    if (files.length == 1) {
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
    super.processOpts();

    File outputFolderFile = new File(outputFolder);
    Path outputFolderFilePath = outputFolderFile.toPath();
    File goModFile = resolveGoMod(outputFolderFile);
    if (goModFile == null) {
      throw new RuntimeException("Can not find go.mod in a project hierarchy");
    }
    Path goModFilePath = goModFile.toPath();
    String moduleName = extractModule(goModFilePath);
    additionalProperties.put(
        CodegenConstants.INVOKER_PACKAGE,
        moduleName + "/" + goModFilePath.getParent().relativize(outputFolderFilePath).toString());

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
  public OperationsMap postProcessOperationsWithModels(OperationsMap objs, List<ModelMap> allModels) {
    OperationsMap operationsMap = super.postProcessOperationsWithModels(objs, allModels);
    List<Map<String, String>> imports = objs.getImports();
    if (imports == null)
      return objs;

    operationsMap.put("hasImportedModel", false);
    Iterator<Map<String, String>> iterator = imports.iterator();
    List<String> importedModels = new ArrayList<>();
    while (iterator.hasNext()) {
      String _import = iterator.next().get("import");
      if (_import.startsWith(modelPackage())) {
        operationsMap.put("hasImportedModel", true);
        importedModels.add(_import);
      }
    }
    operationsMap.put("importedModels", importedModels);

    return operationsMap;
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
}
