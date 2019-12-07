#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <sys/types.h>
#include <unistd.h>

int main()
{
	setreuid(1000,1000);
	setregid(1000, 1000);	
	setvbuf(stdout, NULL, _IONBF, 0);
	char arg[64] ;
	char command[64];
	while(1)
	{
		puts("provide argument to ls : ");
		scanf("%s",&arg);
		char * found = strstr( arg, "sh" );
		char * found1 = strstr( arg, "vi" );
		char * found2 = strstr( arg, "ex" );
		char * found3 = strstr( arg, "ed" );
		
		if(strlen(arg) > 3 || found != NULL | found1 != NULL | found2 != NULL | found3 != NULL )
		{
			puts("argument not allowed exiting ...");
			return 1;
		}
		else
		{
			strcpy(command,"ls ");
			strcat(command , arg);
			printf("executing %s \n",command);
			system(command);
			return 0;
		}
	}
}
