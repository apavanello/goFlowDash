db.getCollection("boxes").insertMany([{"_id":"1","boxType":"default","label":"Fluxo 1 \n processo inicial unico","Column":1,"position":{"x":300,"y":100},"extras":{"status":"done"}},
{"_id":"2","boxType":"default","label":"Fluxo 2","Column":2,"position":{"x":500,"y":100},"extras":{"status":"running"}},
{"_id":"3","boxType":"default","label":"Fluxo 3","Column":2,"position":{"x":500,"y":100},"extras":{"status":"done"}},
{"_id":"4","boxType":"default","label":"Fluxo 4","Column":2,"position":{"x":500,"y":100},"extras":{"status":"running"}},
{"_id":"5","boxType":"default","label":"Fluxo 5","Column":3,"position":{"x":700,"y":100},"extras":{"status":""}},
{"_id":"6","boxType":"default","label":"Fluxo 6","Column":3,"position":{"x":700,"y":100},"extras":{"status":"running"}},
{"_id":"7","boxType":"default","label":"Fluxo 7","Column":4,"position":{"x":900,"y":100},"extras":{"status":""}},
{"_id":"8","boxType":"default","label":"Fluxo 8","Column":4,"position":{"x":900,"y":100},"extras":{"status":""}},
{"_id":"9","boxType":"default","label":"Fluxo 9 - final unico","Column":5,"position":{"x":1100,"y":100},"extras":{"status":""}}
])