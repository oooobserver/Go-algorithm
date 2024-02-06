test:
	mkdir -p tmp
	cd tmp && touch h.txt
	echo "hello" > tmp/h.txt
	cat tmp/h.txt

clean:
	rm -rf tmp

.PHONY: test clean


