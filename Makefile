add-license:
	go tool nwa add -l mit -c "Gurkan" -y 2025 "**/*.go"

check-license:
	go tool nwa check -l mit -c "Gurkan" -y 2025 "**/*.go"

update-license:
	go tool nwa update -l mit -c "Gurkan" -y 2025 "**/*.go"