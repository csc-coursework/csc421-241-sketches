
#include<stdio.h>
#include<stdlib.h>
#include<unistd.h>
#include<pthread.h>
#include<string.h>
#include<sys/time.h>
#include<assert.h>

/*
 * accumulate: a program in C to show thread programing
 * and a mutex
 *
 * last-update:
 *	23 sep 2023 bjr- created
 *
 */

#define N_THREADS 5 

pthread_mutex_t lock_g = PTHREAD_MUTEX_INITIALIZER;
int accumulator_g ;


void * accumulate(void * the_args) {
	int i ;
	pthread_mutex_lock(&lock_g) ;
	i = accumulator_g ;
	printf("thread sleeps\n") ;	
	sleep(1) ;	
	accumulator_g = i + 1 ;	
	pthread_mutex_unlock(&lock_g) ;
	return NULL ;
}


int main(int argc, char * argv[]) {

	int i, rc ; 
	void * status ;
	pthread_t thread_id[N_THREADS] ;

	for (i=0; i<N_THREADS; i++) {
		if (pthread_create( thread_id+i, NULL, 
		accumulate, NULL  )) {
			perror("pthread_create") ;
			exit(-1) ;
		}
	}

	// let's get it started
	pthread_mutex_unlock(&lock_g) ;

	for (i=0; i<N_THREADS; i++ ) {
		if (pthread_join(thread_id[i], &status)) {
			perror("pthread_join") ;
			exit(-1) ;
		}
	}
	printf("the final value of accumulator is %d\n", accumulator_g) ;
	return 0 ;
}

