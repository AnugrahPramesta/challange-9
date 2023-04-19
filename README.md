# challange-9

Membuat API dengan Multi-Lever users : Admin & User

- Diagram Alur
  ![image](/images/diagramalur.PNG)

### Dokumentasi API

Berikut merupakan dokumentasi API yang telah dibuat:

- User
  | Method | Endpoint | Description |
  | ------ | --------------- | ------------- |
  | POST | /users/register | Register User |
  | POST | /users/login | Login User |

- Product
  | Method | Endpoint | Description |
  |--------|----------|-------------|
  | GET | /products | Get all products |
  | GET | /products/:productId | Get product by id |
  | POST | /products | Create new product |
  | PUT | /products/:productId | Update product by id |
  | DELETE | /products/:productId | Delete product by id |
