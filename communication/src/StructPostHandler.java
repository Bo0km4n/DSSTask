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

public class StructPostHandler extends Middleware implements HttpHandler {
    @Override
    public void handle(HttpExchange ex) throws IOException {
        // リクエストのヘッダなどを表示
        logRequest(ex);
        InputStream reqBody = ex.getRequestBody();
        
        try {
            ObjectInputStream ois = new ObjectInputStream(reqBody);
            Person p = (Person)ois.readObject();
            Task t = new Task();
            t.hello(p);
        } catch (Exception e) {
            e.printStackTrace();
        }
        byte[] b = "ok".getBytes(StandardCharsets.UTF_8);
        respond(ex, b);
    }
}