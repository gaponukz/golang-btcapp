# GSES2 BTC Application

<p align="center" width="100%">
    <img width="25%" src="https://github.com/gaponukz/golang-btcapp/assets/49754258/a68fe726-1067-4dcd-ae07-0973d7437ddb">
</p>

Golang version of [gaponukz/GSES2-BTC-application](https://github.com/gaponukz/GSES2-BTC-application)

## API interface
| Method | Description |
| :--- | :--- |
| GET `/rate` | Get the current rate of btc to Hryvnia |
| POST `/subscribe` | Subscribe `gmail` to the newsletter with the BTC price |
| POST `/sendEmails` | Send the BTC price to all subscribers of the newsletter. |

## Some logic explanation
![image](https://github.com/gaponukz/GSES2-BTC-application/assets/49754258/474fd9cd-2d01-4642-aa65-18cb55323e9d)

## Settings
Before usage you need to create `.env` file:
```env
gmail=user@gmail.com
gmailPassword=userpassowrf123
```

## Deploying
### From git (locally)
```bash
git clone https://github.com/gaponukz/golang-btcapp.git
cd golang-btcapp
go mod download
go build
./btcapp
```

### From docker
```bash
docker pull gaponukz/gobtcapp
docker run -d -p 8080:8080 --rm --env-file .env gaponukz/gobtcapp 
```
