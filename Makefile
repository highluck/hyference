build-fastText:
	rm -rf ml/lib/fastText
	cd ml/lib/ && git clone https://github.com/facebookresearch/fastText.git
	cd ml/lib/fastText && make
	rm -rf ml/fasttext_wrapper/fastText/obj
	mkdir ml/fasttext_wrapper/fastText/obj
	cp ml/lib/fastText/*.o ml/fasttext_wrapper/fastText/obj/
	cd ml/fasttext_wrapper/fastText && make

build-tensorflow:
	rm -rf ml/lib/tensorflow
	mkdir ml/lib/tensorflow
	curl -L "https://storage.googleapis.com/tensorflow/libtensorflow/libtensorflow-cpu-darwin-x86_64-2.5.0.tar.gz" | tar -C ml/lib/tensorflow -xz


build: build-fastText
	go build
