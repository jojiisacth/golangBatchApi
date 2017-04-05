# Golang based BatchApi
## Goal 
Implement a basic version of batch api  on go language   

##  Contents   
1. Batch API 
A restful api which will accept multiple requests combined as a batch request.  
Individual requests are executed simultaneously. Once all requests are processed the combined result are returned  
2. Product  API 
A restful api which will server  product information
3. Price  API 
A restful api which will server  price information
4. Promotion  API 
A restful api which will server  promotion information

##  Steps  to run
1. clone the repository

       git clone  https://github.com/jojiisacth/golangBatchApi.git
2. start the individual apis as follows   

        cd golangBatchApi
        cd priceapi && ./setup.sh  
        cd ./../productapi && ./setup.sh 
        cd ./../promotionapi && ./setup.sh  
        cd .. 
        
3. start the batch api as follows

        cd batchApi && ./setup.sh

     Now the bath api is ready     
4. Test 
    send  bact requests as follow

         curl -vX POST http://localhost:8080/batch/ -d @request.json --header "Content-Type: application/json"
    
