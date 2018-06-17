
import java.io.*;

public class MooMain {

	public static void dumpToFile(Moo moo, String name) {
		try {
			ObjectOutputStream oos = new ObjectOutputStream(new FileOutputStream(name));
			oos.writeObject(moo);
		} catch (IOException e) {
			e.printStackTrace();
		}
	}

	public static Moo restoreFromFile(String name) {
		try {
			ObjectInputStream ois = new ObjectInputStream(new FileInputStream(name));
			return (Moo)ois.readObject();
		} catch (Exception e) {
			e.printStackTrace();
		}

		return null;
	}

	public static void main(String ...args) {
		Moo moo = new Moo(0x12345678);
		MooMain.dumpToFile(moo, "moo.hex");
		Moo moo2 = MooMain.restoreFromFile("moo.hex");
		moo2.print();
	}

	
}
