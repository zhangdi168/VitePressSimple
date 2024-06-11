- config.ts 入口文件
- vpsimple.ts 存放来自vpsimple的配置,请勿动！
- customer.ts 存放一些自定义的配置，或者需要覆盖vpsimple.ts的配置，比如第三方插件的配置等
> customer config优先级更大,如果两个对象有相同的属性，customer的属性值会覆盖前面的。