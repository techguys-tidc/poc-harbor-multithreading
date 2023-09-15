FROM alpine:latest

RUN dd if=/dev/zero of=test.img bs=1 count=0 seek=100M

RUN echo "$(date +%s)000" 

CMD ["sh"]
