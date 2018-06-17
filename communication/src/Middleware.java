import com.sun.net.httpserver.HttpServer;
import com.sun.net.httpserver.HttpExchange;
import com.sun.net.httpserver.HttpHandler;
import java.io.IOException;
import java.net.InetSocketAddress;
import java.nio.charset.StandardCharsets;
import java.util.concurrent.Executors;
import java.io.*;
import java.util.*;
import java.lang.Class;
import java.lang.reflect.Method;
public class Middleware {

    public static void logRequest(HttpExchange ctx) {
       // 「GET / HTTP1.0」などと表示
       System.out.println(ctx.getRequestMethod() + " / " + ctx.getProtocol());
       // リクエストのヘッダを表示
       ctx.getRequestHeaders().forEach((k,v) -> System.out.println(k + ": "+v));
       System.out.println(); 
    }

    public static void respond(HttpExchange ctx, byte[] responseBody) {
        ctx.getResponseHeaders().add("Content-Type", "text/html; charset=UTF-8");
        try {
            ctx.sendResponseHeaders(200, responseBody.length);  // 明示的に返す必要あり
            ctx.getResponseBody().write(responseBody);

        }catch(IOException ex){
            ex.printStackTrace();
        } 
    }
}