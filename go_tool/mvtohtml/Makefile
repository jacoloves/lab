MAKE=go build -o mvtohtml
MAIN=main.go
GEN=generate.go
EXE=execute.go
REG=reg.go
PARA=paragraph.go
CSS=css.go


mvtohtml:
	$(MAKE) $(MAIN) $(EXE) $(GEN) $(CSS) $(REG) $(PARA)

test:
	./mvtohtml test.md

clean:
	rm -rf mvtohtml
