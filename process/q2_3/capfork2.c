#include "capfork2.h"

int main(void) {
    pid_t pid;

    int pp[2];      // pipe
    int pp_2[2];
    char buf[256];

    pipe(pp);       // make pipe
    pipe(pp_2);
    pid = fork();  // make child process
    
    if (pid == 0) {
        // (2) read pipe from parent
        p_message *m = (p_message *)malloc(sizeof(p_message));
        char read_buf[256];
        m->buf = read_buf;
        m->pp = pp;
        m->type = "child";
        read_from_pipe(m);
        capitalize(read_buf);

        // (3) write capitalized string to pipe to parent
        m->buf = read_buf;
        m->pp = pp_2;
        m->type = "child";
        write_to_pipe(m);
    } else {
        // (1) send message
        printf("$ type your message >>> ");
        fgets(buf, 256, stdin);
        p_message *m = (p_message *)malloc(sizeof(p_message));
        m->buf = buf;
        m->pp = pp;
        m->type = "parent";
        write_to_pipe(m);

        // (4) read pipe from child
        char read_buf[256];
        m->buf = read_buf;
        m->pp = pp_2;
        m->type = "parent";
        read_from_pipe(m);
    }
    
}


char * capitalize_str(char *c) {
    char *p;
    for(p=c; *p; p++) {
        *p = (char)toupper((int)*p);
    }
    return c;
}

// capitalize
void capitalize(void *str) {
    char *buf = (char *)str;
    buf = capitalize_str(buf);
}

int write_to_pipe(p_message *m) {
    close(m->pp[0]);
    printf("[%s] write message : %s\n", m->type, m->buf);
    write(m->pp[1], m->buf, strlen(m->buf) + 1);
    return close(m->pp[1]);
}

void read_from_pipe(p_message *m) {
    close(m->pp[1]);
    read(m->pp[0], m->buf, 256);
    printf("[%s] read message from pipe: %s\n", m->type, m->buf);
    close(m->pp[0]);
}