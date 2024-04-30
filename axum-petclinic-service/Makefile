watch-src-change:
	cargo watch -q -c -w src/ -x run   # Run the program on file change

check-src-change:
	cargo watch -q -c -w src/ -x check # Check the program on file change

watch-test-change:
	cargo watch -q -c -w tests/ -x "test -q quick_dev -- --nocapture"     # Run tests on file change
