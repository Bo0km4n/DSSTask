import java.rmi.registry.Registry;
import java.rmi.registry.LocateRegistry;
import java.rmi.RemoteException;
import java.rmi.server.UnicastRemoteObject;
	
public class Server implements Hello {
	
    public Server() {}

    public String sayHello() {
    	return "Hello, world!";
    }
    
    @Override
	public Object execute(Task t) {
    	System.out.println("invoke remote execute");
    	Object r = t.execute();
    	System.out.println("server status " + t.getStatus());
		return r;
	}
	
    public static void main(String args[]) {
	
	try {
		
	    Server obj = new Server();
	    Hello stub = (Hello) UnicastRemoteObject.exportObject(obj, 0);

	    // Bind the remote object's stub in the registry
	    Registry registry = LocateRegistry.createRegistry(Registry.REGISTRY_PORT); // new style. don't need rmiregistry
	    //Registry registry = LocateRegistry.getRegistry(); // old style
	    registry.bind("Hello", stub);

	    System.err.println("Server ready");
	} catch (Exception e) {
	    System.err.println("Server exception: " + e.toString());
	    e.printStackTrace();
	}
    }	
}
