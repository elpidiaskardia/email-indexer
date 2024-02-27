# email-indexer

This project indexes emails using Zinsearch as a database, where you can search for a word or phrase within the content of emails using a search bar. Golang was used for the backend, and Vue was used for the frontend.


# Screenshots
## Web
![screenshot web](https://raw.githubusercontent.com/elpidiaskardia/email-indexer/main/screenshots/web.PNG)

## Movil
![movil web](https://raw.githubusercontent.com/elpidiaskardia/email-indexer/main/screenshots/movil.PNG)

# Tecnologies Used
- **Backend -  Golang/go-chi**
  
  The backend is constructed using the Go programming language along with the Go-chi library. It functions as a service layer responsible for interacting with the ZincSearch database, fetching data, refining pertinent information, and     transmitting it to the frontend. Additionally, there is a profiler (pprof) accessible to scrutinize the backend's performance and pinpoint potential areas for enhancement.

  go version: 1.22.0

  go-chi version: v1.5.5

- **Database:** Zincsearch

  ZincSearch is a Go-based alternative to ElasticSearch, and you can explore its official website at ZincSearch. To populate ZincSearch with the email db folder, a Go script has been developed, conveniently located in the data folder. Following data ingestion, the backend establishes communication with ZincSearch through its endpoints to fetch the required information.

  Documentation zincsearch:  https://zincsearch-docs.zinc.dev/quickstart/

  The Data Set can be downloaded in http://www.cs.cmu.edu/~enron/enron_mail_20110402.tgz

- **Interface:** Vue 3

  
  The user interface for Mamuro email is constructed using Vue3 alongside the TypeScript Composition API, providing a codebase that is not only more maintainable but also scalable. Tailwind CSS is employed to style the application. The frontend is responsible for managing pagination and enables users to easily access the complete email content with just a single click.

  npm version: 10.2.4

  node version: 18.17.0

