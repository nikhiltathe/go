#!/bin/bash


go tool cover -html=./$(BaseFolder)/$(CoverAll).out -o ./$(BaseFolder)/$(CoverAll).html
go tool cover -func=./$(BaseFolder)/$(CoverAll).out
