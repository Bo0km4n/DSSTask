import com.sun.net.httpserver.HttpServer;
import com.sun.net.httpserver.HttpExchange;

import java.io.IOException;
import java.net.InetSocketAddress;
import java.nio.charset.StandardCharsets;
import java.util.concurrent.Executors;
import java.io.*;
import java.util.*;
import java.lang.Class;
import java.lang.reflect.Method;

public class Server {

    public static void main(String... args)throws IOException{

        HttpServer server = HttpServer.create(new InetSocketAddress(18888), 0);

        // Handlers
        StructFetchHandler sfh = new StructFetchHandler();
        StructPostHandler sph = new StructPostHandler();


        server.setExecutor(Executors.newCachedThreadPool());  // Executor の設定
        server.createContext("/api/v1/struct", sfh);
        server.createContext("/api/v1/post_struct", sph);
        server.start();
        System.out.println("start http listening :18888");
    }
}