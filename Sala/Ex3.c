#include <stdio.h>

void main(){
	
	int qtdAluno, i;
	float mediaTotal = 0, mediaAluno;
	
	printf("Quantos alunos? ");
	scanf("%d", &qtdAluno);
	
	for(i = 0; i< qtdAluno; i++){
		printf("Nota do aluno %d: ", (i + 1));
		scanf("%f", &mediaAluno);
		mediaTotal += mediaAluno;
	}
	mediaTotal/=qtdAluno;
	printf("Media da turma: %.2f", mediaTotal);
	
}
