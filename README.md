# 生成图片给客户端当作mock用的

启动之后，用下面这个链接试试？

> [http://localhost:8080/draw?dsl=rect:w=500,h=200,color=ff0000:](http://localhost:8080/draw?dsl=rect:w=500,h=200,color=ff0000:)

目前的规则很简单，只有矩形和颜色，以后有新需求再增加.

```
rect是矩形的意思。
w是宽
h是高
color是16进制色，参考 https://sunpma.com/other/rgb/
```

---

### dsl格式说明

画许多贝塞尔曲线：
```
beziers:w=500,h=200,color=ff0000:square=100,row=5,column=2
```

画圆形：
```
circle:w=500,h=500,color=FF6EB4:x=250,y=250,radius=250
```

许多横竖线小格子：
```
crisp:w=500,h=500,color=FF6EB4:minor=10,major=50
```

cubic:
```
crisp:w=500,h=500,color=FF6EB4:
```

许多椭圆形成的花:
```
ellipse:w=500,h=500,color=FF6EB4:
```

图片中画出字体：
```
gofont:w=500,h=500,color=FF6EB4:size=24,char=iofdfa fdsafd,color=203fff
```

一力我滴giao

