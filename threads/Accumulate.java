//
// Accumulate and AccumulateSync
// last-update:
// 	22 sep 2023 bjr:
//

public class Accumulate implements Runnable {

	public static int accumulator = 0 ;
	public static long time_now = System.currentTimeMillis() ;

	public static String timer_string() {
		double t ;
		t = ((double)(System.currentTimeMillis()-Accumulate.time_now))/1000.0d ;
		return "["+String.format("%,.4f", t)+"]" ;
	}

	public void run() {
		int i ; 
		try {
		    System.out.println(this.timer_string()+" reading i") ;
			i = Accumulate.accumulator ;
			Thread.sleep(2000) ;
		    System.out.println(this.timer_string()+" writing i") ;
			Accumulate.accumulator = i+1 ;
		} catch (InterruptedException e) { }
	    System.out.println(this.timer_string()+" thread exits") ;
	}


	public static void main(String args[]) {
		Thread [] threads = new Thread[5] ;
		for (int i = 0; i< 5; i++) {
			threads[i] = new Thread(new Accumulate()) ;
			threads[i].start() ;
		}
		try {
			for (int i = 0; i< 5; i++) threads[i].join() ;
		} catch (InterruptedException e) { }
		System.out.println("accumulator = " + Accumulate.accumulator) ;
    	}

}

