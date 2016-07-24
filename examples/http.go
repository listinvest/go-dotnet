package main

import (
	"github.com/matiasinsaurralde/go-dotnet"

	"net/http"

	"fmt"
	"os"
	// "time"
)

//export the_callback
func the_callback() {
	fmt.Println("Hi, I'm the callback!")
}

func main() {
	fmt.Println("This is main, I'll initialize the .NET runtime.")

	properties := map[string]string{
		"TRUSTED_PLATFORM_ASSEMBLIES":   "/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/mscorlib.ni.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Private.CoreLib.ni.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/Microsoft.CodeAnalysis.CSharp.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/Microsoft.CodeAnalysis.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/Microsoft.CodeAnalysis.VisualBasic.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/Microsoft.CSharp.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/Microsoft.VisualBasic.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/Microsoft.Win32.Primitives.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/Microsoft.Win32.Registry.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/mscorlib.ni.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.AppContext.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Buffers.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Collections.Concurrent.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Collections.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Collections.Immutable.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.ComponentModel.Annotations.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.ComponentModel.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Console.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Diagnostics.Debug.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Diagnostics.DiagnosticSource.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Diagnostics.FileVersionInfo.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Diagnostics.Process.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Diagnostics.StackTrace.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Diagnostics.Tools.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Diagnostics.Tracing.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Dynamic.Runtime.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Globalization.Calendars.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Globalization.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Globalization.Extensions.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.IO.Compression.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.IO.Compression.ZipFile.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.IO.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.IO.FileSystem.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.IO.FileSystem.Primitives.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.IO.FileSystem.Watcher.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.IO.MemoryMappedFiles.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.IO.UnmanagedMemoryStream.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Linq.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Linq.Expressions.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Linq.Parallel.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Linq.Queryable.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Net.Http.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Net.NameResolution.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Net.Primitives.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Net.Requests.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Net.Security.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Net.Sockets.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Net.WebHeaderCollection.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Numerics.Vectors.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.ObjectModel.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Private.CoreLib.ni.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Private.Uri.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Reflection.DispatchProxy.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Reflection.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Reflection.Emit.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Reflection.Emit.ILGeneration.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Reflection.Emit.Lightweight.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Reflection.Extensions.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Reflection.Metadata.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Reflection.Primitives.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Reflection.TypeExtensions.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Resources.Reader.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Resources.ResourceManager.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Runtime.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Runtime.Extensions.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Runtime.Handles.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Runtime.InteropServices.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Runtime.InteropServices.RuntimeInformation.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Runtime.Loader.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Runtime.Numerics.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Security.Claims.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Security.Cryptography.Algorithms.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Security.Cryptography.Cng.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Security.Cryptography.Csp.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Security.Cryptography.Encoding.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Security.Cryptography.OpenSsl.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Security.Cryptography.Primitives.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Security.Cryptography.X509Certificates.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Security.Principal.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Security.Principal.Windows.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Text.Encoding.CodePages.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Text.Encoding.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Text.Encoding.Extensions.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Text.RegularExpressions.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Threading.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Threading.Overlapped.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Threading.Tasks.Dataflow.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Threading.Tasks.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Threading.Tasks.Extensions.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Threading.Tasks.Parallel.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Threading.Thread.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Threading.ThreadPool.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Threading.Timer.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Xml.ReaderWriter.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Xml.XDocument.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Xml.XmlDocument.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Xml.XPath.dll:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0/System.Xml.XPath.XDocument.dll",
		"APP_PATHS":                     "/Users/matias/dev/dotnet/cdotnet/lib/HelloWorld",
		"NATIVE_DLL_SEARCH_DIRECTORIES": "/Users/matias/dev/dotnet/cdotnet/lib/HelloWorld:/usr/local/share/dotnet/shared/Microsoft.NETCore.App/1.0.0",
	}

	err, runtime := dotnet.NewRuntime(dotnet.RuntimeParams{
		Properties:                  properties,
		ManagedAssemblyAbsolutePath: "HelloWorldMain.Exe",
	})

	if err != nil {
		fmt.Println("Something bad happened! :(")
		os.Exit(1)
	}

	SayHello := runtime.CreateDelegate("HelloWorld", "HelloWorld.HelloWorld", "Hello")

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Calling delegate")
		SayHello()
		w.Write([]byte("???"))
	})

	http.ListenAndServe(":5000", nil)

	runtime.Shutdown()
}