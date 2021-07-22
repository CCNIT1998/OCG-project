const formatMoney = function(value){
    let val = (value/1)
    return val.toString().replace(/\B(?=(\d{3})+(?!\d))/g, ".")
}   


export default formatMoney;