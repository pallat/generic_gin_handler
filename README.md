# Generic experiment

พยายามทดลองสร้าง generic router

## What I've learned

- func(request) (response,error) รูปแบบนี้ไม่เหมาะ เพราะในการใช้งานจริง มีการนำเข้า request หลายรูปแบบมากเกินไป และในขณะเดียวกัน response เองก็มีความหลากหลายมากเช่นกัน
- พบว่ารูปแบบการใช้ interface เป็น request น่าจะเหมาะสมกว่า ส่วน response ก็มีความหลากหลายเช่น status code, response, no-content
