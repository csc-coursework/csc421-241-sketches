

public class Accumulate implements Runnable {

	public static int accumulator = 0 ;
	public static Object lock = new Object() ;

	public synchronized void run() {
		int i ;
		i = Accumulate.accumulator ;
		try {
			Thread.sleep(2000) ;
		} catch ( InterruptedException e) {
			System.out.println("thread sleep interrupted") ;
		}
		System.out.println("adding one") ;
		synchronized (Accumulate.lock) {
			i = Accumulate.accumulator ;
			Accumulate.accumulator = i+1 ;
		}
	}


	public static void main(String args[]) {
		for (int i = 0; i< 5; i++) {
	        	(new Thread(new Accumulate())).start();
		}
		try {
			Thread.sleep(4000) ;
		} catch ( InterruptedException e) {
	                       System.out.println("thread sleep interrupted") ;
	        }
		System.out.println("accumulator = " + Accumulate.accumulator) ;
    	}

}

