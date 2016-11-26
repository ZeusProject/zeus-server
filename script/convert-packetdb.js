#!/usr/bin/env node

let fs = require('fs')
let file = process.argv[2]
let version = process.argv[3]

let data = fs.readFileSync(file)
let lines = data.toString().split('\n')

let keys = []
let packets = {}

for (let line of lines) {
	line = line.trim()

	if (line == '') {
		continue
	}

	let byColon = line.split(':', 2)
	let cmd = byColon[0].trim()
	let value = (byColon[1] || '').trim()

	if (cmd == 'packet_ver') {
		if (parseInt(value) > parseInt(version)) {
			break
		}
	} else if (cmd == 'packet_keys') {
		keys = value.split(',').map(parseNumber)
	} else {
		let cols = cmd.split(',')
		let packet = parseNumber(cols[0])
		let size = parseNumber(cols[1])
		let name = cols[2]

		packets[packet] = {
			packet: packet,
			size: size,
			name: name
		}
	}
}

console.log('%s:', version)
console.log('  packets:')

for (let p in packets) {
	let packet = packets[p]

	console.log('    0x%s:', packet.packet.toString(16))
	console.log('      packet: %s', packet.name || 'SS_NONE')
	console.log('      size: %s', packet.size)
}

function parseNumber (x) {
	let base = 10

	if (x.startsWith('0x')) {
		base = 16
		x = x.slice(2)
	}

	return parseInt(x, base)
}

