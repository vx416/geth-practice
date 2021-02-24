

abigen:
	abigen --abi=./client/abi/${f}.abi --pkg=abi --out=./client/abi/${f}.go

test:
	echo testmake