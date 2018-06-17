import java.io.*;

public class Moo implements Serializable {
	public int v;
	public String str1 = "test1";
	//	public Object obj1 = "test2";
	public String str2 = "test2";
	public byte x = 1;
	public Moo(){}
	public Moo(int v) { this.v = v; }
	public void print() {
		System.out.printf("v = %x\n", v);
	}

	public void test() {
		System.out.printf("hello");
	}

	/*
	private void writeObject(ObjectOutputStream out) throws IOException {
		out.defaultWriteObject();
	}
	*/
	
}
