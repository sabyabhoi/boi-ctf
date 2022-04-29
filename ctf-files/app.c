#include <stdio.h>
#include <signal.h>

const char* FLAG = "REDACTED";

void segHandler() {
	puts(FLAG);
}

int main() {


	signal(SIGSEGV, segHandler);
	signal(SIGABRT, segHandler);

	puts("What's your name? ");
	fflush(stdout);
	char name[16];

	gets(name);
	puts("Hi");
	puts(name);

	return 0;
}
