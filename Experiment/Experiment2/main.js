const app = new Vue({
    el: '#app',
    data: {
        books: [
            {
                id: 1,
                name: '《算法导论》',
                date: '2006-9',
                price: 85.00,
                count: 1
            },
            {
                id: 2,
                name: '《UNIX编程艺术》',
                date: '2006-2',
                price: 59.00,
                count: 1
            },
            {
                id: 3,
                name: '《编程珠玑》',
                date: '2008-10',
                price: 39.00,
                count: 1
            },
            {
                id: 4,
                name: '《代码大全》',
                date: '2006-3',
                price: 128.00,
                count: 1
            },
        ]
    },
    methods: {

        add:function(i){
            this.books[i].count++
        },
        sub:function(i){
            this.books[i].count--;
        },
        remove:function(i){
            this.books.splice(i,1);
        }


    },
    filters:{//过滤器
        getFinalPrice:function(price){
            return '￥'+price.toFixed(2);
        }
    },
    computed: {
        finalPrice:function(){
            let sum=0;
            for(let i=0;i<this.books.length;i++){
                sum=sum+this.books[i].price*this.books[i].count
            }
            return sum
        }
    }
})





