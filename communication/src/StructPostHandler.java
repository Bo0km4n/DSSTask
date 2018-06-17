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
        byte[] body = "hello".getBytes();
        InputStream reqBody = ex.getRequestBody();
        System.out.println(inputStreamToString(reqBody));
        // レスポンスを返す
        respond(ex, body);
    }

    private static String inputStreamToString(InputStream is) throws IOException {
        BufferedReader reader = null;
        reader = new BufferedReader(new InputStreamReader(is));
        StringBuilder sb = new StringBuilder();
        String b = null;
        while ((b=reader.readLine()) !=null){
            sb.append(b);
        }
        return sb.toString();
    }

    private static byte[] inputStreamToByteArray(InputStream is) throws IOException {
        ByteArrayOutputStream bout = new ByteArrayOutputStream();
        byte [] buffer = new byte[1024];
        while(true) {
            int len = is.read(buffer);
            if(len < 0) {
                break;
            }
            bout.write(buffer, 0, len);
        }
        return bout.toByteArray();
    }
}