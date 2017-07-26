#!/bin/bash

JSON="{ \"name\" : \"Pisco Sour\", \"ingredients\" : [ \"blanc d'oeuf\", \"pisco\", \"citron\" ] }"

curl --data "$JSON" http://localhost:8080/cocktail


