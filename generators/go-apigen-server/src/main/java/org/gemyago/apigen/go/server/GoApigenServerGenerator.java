package org.gemyago.apigen.go.server;

import java.io.File;
import java.io.FileFilter;
import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.util.List;
import java.util.Set;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

import org.openapitools.codegen.*;
import org.openapitools.codegen.languages.*;
import org.openapitools.codegen.templating.mustache.CamelCaseAndSanitizeLambda;
import org.openapitools.codegen.templating.mustache.IndentedLambda;

import com.google.common.collect.ImmutableMap;
import com.samskivert.mustache.Mustache;

import io.swagger.v3.oas.models.Operation;
import io.swagger.v3.oas.models.parameters.Parameter;
import io.swagger.v3.oas.models.servers.Server;

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
  }

  @Override
  protected ImmutableMap.Builder<String, Mustache.Lambda> addMustacheLambdas() {
    ImmutableMap.Builder<String, Mustache.Lambda> lambdas = super.addMustacheLambdas();
    lambdas.put("tab_indented_2", new IndentedLambda(1, "\t", "", false, false));
    return lambdas;
  }

  @Override
  public CodegenOperation fromOperation(String path, String httpMethod, Operation operation, List<Server> servers) {
    // TODO: Make sure this is customizable, just default is set
    operation.addExtension("x-codegen-request-body-name", "payload");
    return super.fromOperation(path, httpMethod, operation, servers);
  }

  @Override
  public CodegenParameter fromParameter(Parameter parameter, Set<String> imports) {
    CodegenParameter codegenParameter = super.fromParameter(parameter, imports);
    codegenParameter.vendorExtensions.put("x-codegen-param-in", parameter.getIn());
    return codegenParameter;
  }
}
