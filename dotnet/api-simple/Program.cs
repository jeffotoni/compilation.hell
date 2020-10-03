using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.Hosting;
using Microsoft.Extensions.Logging;
using Microsoft.AspNetCore.Builder;
using Microsoft.Extensions.DependencyInjection;
using Microsoft.AspNetCore.Http;
Host.CreateDefaultBuilder(args).ConfigureWebHostDefaults(webBuilder => webBuilder.Configure(app =>
    app.UseRouting().UseEndpoints(endpoints =>
        endpoints.MapGet("/", async context => await context.Response.WriteAsync("Hello world!"))))
).Build().Run();
