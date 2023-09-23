//
// Accumulate and AccumulateSync
// last-update: 
//	22 sep 2023 bjr:
//

public class AccumulateSync implements Runnable {

	public static int accumulator = 0 ;
	public static Object lock = new Object() ;
	public static long time_now = System.currentTimeMillis() ;

	public static String timer_string() {
		double t ;
		t = ((double)(System.currentTimeMillis()-AccumulateSync.time_now))/1000.0d ;
		return "["+String.format("%,.4f", t)+"]" ;
	}

	public void run() {
		int i ; 
		try {
		synchronized(AccumulateSync.lock) {
		       	System.out.println(this.timer_string()+" reading i") ;
			i = AccumulateSync.accumulator ;
			Thread.sleep(2000) ;
		       	System.out.println(this.timer_string()+" writing i") ;
			AccumulateSync.accumulator = i+1 ;
		}
		} catch (InterruptedException e) {
		}
	       	System.out.println(this.timer_string()+" thread exits") ;
	}


	public static void main(String args[]) {
		Thread [] threads = new Thread[5] ;
		for (int i = 0; i< 5; i++) {
			threads[i] = new Thread(new AccumulateSync()) ;
			threads[i].start() ;
		}
		try {
			for (int i = 0; i< 5; i++) threads[i].join() ;
		} catch (InterruptedException e) {
		}
		System.out.println("accumulator = " + AccumulateSync.accumulator) ;
    	}

}

