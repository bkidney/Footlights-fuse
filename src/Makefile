BITSTRING_DIR = $(shell ocamlfind query bitstring)
BITSTRING_MODULE = $(BITSTRING_DIR)/bitstring.cma
BITSTRING_DEPS = $(BITSTRING_MODULE) $(BITSTRING_DIR)/bitstring_persistent.cma $(BITSTRING_DIR)/pa_bitstring.cmo

OCAML_OPTS = -pp 'camlp4o $(BITSTRING_DEPS)'

OBJECTS = link block
PROG = distorage

all: $(PROG)

$(PROG): $(OBJECTS)

link: 
	ocamlfind ocamlc $(OCAML_OPTS) -linkpkg -package bitstring link.ml -o link

block: 
	ocamlfind ocamlc $(OCAML_OPTS) -linkpkg -package bitstring block.ml -o block

# Clean
clean:
	rm -f core *~ *.cmi *.cmo *.cmx *.cma *.cmxa *.cmp *.a *.o *.so
