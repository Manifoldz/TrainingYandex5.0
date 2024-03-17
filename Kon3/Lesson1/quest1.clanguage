#include "stdio.h"
#include "stdlib.h"
#include "string.h"

typedef struct {
    char *key;
    int numPeople;
} KeyValue;

typedef struct {
    KeyValue *elements;
    int size;
    int capacity;
} Dictionary;

void initialize(Dictionary *MyDict);
void addOrUpdate(char *key, Dictionary *MyDict);
int compare(const void *a, const void *b);
void cleanup(Dictionary *dictionary);

int main() {
    int n = 0;
    if (scanf("%d\n", &n) != 1 || n < 1 || n > 2000000) {
        printf("Read-mistake1 n=%d", n);
        return 1;
    }
    Dictionary MyDict;
    initialize(&MyDict);
    if (MyDict.elements == NULL) {
        printf("MistakeOfAllocate#1");
        return 2;
    }

    for (int i = 0; i < n; i++) {
        int numSongs = 0;
        if (scanf("%d\n", &numSongs) != 1) {
            printf("Read-mistake2 numSongs=%d", numSongs);
            return 3;
        }

        for (int j = 0; j < numSongs; j++) {
            char NewSong[100];
            if (scanf("%s", NewSong) != 1) {
                printf("Read-mistake3 j=%d", j);
                return 4;
            }
            addOrUpdate(NewSong, &MyDict);
        }
    }

    qsort(MyDict.elements, MyDict.size, sizeof(KeyValue), compare);
    for (int i = 0; i < MyDict.size; i++) {
        if (MyDict.elements[i].numPeople == n) {
            printf(" %s", MyDict.elements[i].key);
        }
    }
    cleanup(&MyDict);
    return 0;
}

void addOrUpdate(char *key, Dictionary *MyDict) {
    int addedFlag = 0;

    if (MyDict->size >= MyDict->capacity) {
        MyDict->capacity *= 2;
        KeyValue *temp = realloc(MyDict->elements, MyDict->capacity * sizeof(KeyValue));
        if (temp == NULL) {
            printf("MistakeOfAllocate#2");
            return;
        }
        MyDict->elements = temp;
    }

    for (int i = 0; i < MyDict->size && !addedFlag; i++) {
        if (strcmp(MyDict->elements[i].key, key) == 0) {
            MyDict->elements[i].numPeople++;
            addedFlag = 1;
        }
    }

    if (!addedFlag) {
        MyDict->elements[MyDict->size].key = strdup(key);
        MyDict->elements[MyDict->size].numPeople++;
        MyDict->size++;
    }
}

int compare(const void *a, const void *b) { return strcmp(((KeyValue *)a)->key, ((KeyValue *)b)->key); }

void initialize(Dictionary *MyDict) {
    MyDict->size = 0;
    MyDict->capacity = 10;
    MyDict->elements = malloc(MyDict->capacity * sizeof(KeyValue));
}

void cleanup(Dictionary *dictionary) {
    for (int i = 0; i < dictionary->size; i++) {
        free(dictionary->elements[i].key);
    }
    free(dictionary->elements);
}
