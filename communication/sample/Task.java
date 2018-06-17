
import java.io.Serializable;

//public interface Task {
public interface Task extends Serializable {	
	
	Object execute();
	int getStatus();

}
