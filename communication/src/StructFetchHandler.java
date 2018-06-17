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

public class StructFetchHandler extends Middleware implements HttpHandler {
    @Override
    public void handle(HttpExchange ex) throws IOException {
        // リクエストのヘッダなどを表示
        logRequest(ex);
        Task p = new Task();
        ObjectOutput out = null;
        ByteArrayOutputStream bos = new ByteArrayOutputStream();
        out = new ObjectOutputStream(bos);
        out.writeObject(p);
        out.flush();
        byte[] body = bos.toByteArray();
        // レスポンスを返す
        respond(ex, body);
    }
}