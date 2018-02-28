#!/bin/sh                                                                                                                                                                                                                                     

CONTAINER_IMAGE=XXX
PROFILE_PATH=/root/XXX                                                                                                                                                                                                      
BIN_PATH=XXX                                                                                                                                                                                                                        
                                                                                                                                                                                                                                              
docker cp "${CONTAINER_IMAGE}:${PROFILE_PATH}" .                                                                                                                                                                                                       
go tool pprof "${BIN_PATH}" $(basename "${PROFILE_PATH}") 
