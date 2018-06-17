public class MyTask implements Task {
	private int status = 0;

	@Override
	public Object execute() {
		status = 1;
		return new Long(System.currentTimeMillis());
	}
	
	public int getStatus() {
		return status;
	}
	
	

}
