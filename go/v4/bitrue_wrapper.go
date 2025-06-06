package ccxt

type Bitrue struct {
   *bitrue
   Core *bitrue
}

func NewBitrue(userConfig map[string]interface{}) Bitrue {
   p := &bitrue{}
   p.Init(userConfig)
   return Bitrue{
       bitrue: p,
       Core:  p,
   }
}

// PLEASE DO NOT EDIT THIS FILE, IT IS GENERATED AND WILL BE OVERWRITTEN:
// https://github.com/ccxt/ccxt/blob/master/CONTRIBUTING.md#how-to-contribute-code


/**
 * @method
 * @name bitrue#fetchStatus
 * @description the latest known information on the availability of the exchange API
 * @see https://github.com/Bitrue-exchange/Spot-official-api-docs#test-connectivity
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {object} a [status structure]{@link https://docs.ccxt.com/#/?id=exchange-status-structure}
 */
func (this *Bitrue) FetchStatus(params ...interface{}) (map[string]interface{}, error) {
    res := <- this.Core.FetchStatus(params...)
    if IsError(res) {
        return map[string]interface{}{}, CreateReturnError(res)
    }
    return res.(map[string]interface{}), nil
}
/**
 * @method
 * @name bitrue#fetchTime
 * @description fetches the current integer timestamp in milliseconds from the exchange server
 * @see https://github.com/Bitrue-exchange/Spot-official-api-docs#check-server-time
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {int} the current integer timestamp in milliseconds from the exchange server
 */
func (this *Bitrue) FetchTime(params ...interface{}) ( int64, error) {
    res := <- this.Core.FetchTime(params...)
    if IsError(res) {
        return -1, CreateReturnError(res)
    }
    return (res).(int64), nil
}
/**
 * @method
 * @name bitrue#fetchCurrencies
 * @description fetches all available currencies on an exchange
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {object} an associative dictionary of currencies
 */
func (this *Bitrue) FetchCurrencies(params ...interface{}) (Currencies, error) {
    res := <- this.Core.FetchCurrencies(params...)
    if IsError(res) {
        return Currencies{}, CreateReturnError(res)
    }
    return NewCurrencies(res), nil
}
/**
 * @method
 * @name bitrue#fetchMarkets
 * @description retrieves data on all markets for bitrue
 * @see https://github.com/Bitrue-exchange/Spot-official-api-docs#exchangeInfo_endpoint
 * @see https://www.bitrue.com/api-docs#current-open-contract
 * @see https://www.bitrue.com/api_docs_includes_file/delivery.html#current-open-contract
 * @param {object} [params] extra parameters specific to the exchange api endpoint
 * @returns {object[]} an array of objects representing market data
 */
func (this *Bitrue) FetchMarkets(params ...interface{}) ([]MarketInterface, error) {
    res := <- this.Core.FetchMarkets(params...)
    if IsError(res) {
        return nil, CreateReturnError(res)
    }
    return NewMarketInterfaceArray(res), nil
}
/**
 * @method
 * @name bitrue#fetchBalance
 * @description query for balance and get the amount of funds available for trading or funds locked in orders
 * @see https://github.com/Bitrue-exchange/Spot-official-api-docs#account-information-user_data
 * @see https://www.bitrue.com/api-docs#account-information-v2-user_data-hmac-sha256
 * @see https://www.bitrue.com/api_docs_includes_file/delivery.html#account-information-v2-user_data-hmac-sha256
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @param {string} [params.type] 'future', 'delivery', 'spot', 'swap'
 * @param {string} [params.subType] 'linear', 'inverse'
 * @returns {object} a [balance structure]{@link https://docs.ccxt.com/#/?id=balance-structure}
 */
func (this *Bitrue) FetchBalance(params ...interface{}) (Balances, error) {
    res := <- this.Core.FetchBalance(params...)
    if IsError(res) {
        return Balances{}, CreateReturnError(res)
    }
    return NewBalances(res), nil
}
/**
 * @method
 * @name bitrue#fetchOrderBook
 * @description fetches information on open orders with bid (buy) and ask (sell) prices, volumes and other data
 * @see https://github.com/Bitrue-exchange/Spot-official-api-docs#order-book
 * @see https://www.bitrue.com/api-docs#order-book
 * @see https://www.bitrue.com/api_docs_includes_file/delivery.html#order-book
 * @param {string} symbol unified symbol of the market to fetch the order book for
 * @param {int} [limit] the maximum amount of order book entries to return
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {object} A dictionary of [order book structures]{@link https://docs.ccxt.com/#/?id=order-book-structure} indexed by market symbols
 */
func (this *Bitrue) FetchOrderBook(symbol string, options ...FetchOrderBookOptions) (OrderBook, error) {

    opts := FetchOrderBookOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var limit interface{} = nil
    if opts.Limit != nil {
        limit = *opts.Limit
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.FetchOrderBook(symbol, limit, params)
    if IsError(res) {
        return OrderBook{}, CreateReturnError(res)
    }
    return NewOrderBook(res), nil
}
/**
 * @method
 * @name bitrue#fetchTicker
 * @description fetches a price ticker, a statistical calculation with the information calculated over the past 24 hours for a specific market
 * @see https://github.com/Bitrue-exchange/Spot-official-api-docs#24hr-ticker-price-change-statistics
 * @see https://www.bitrue.com/api-docs#ticker
 * @see https://www.bitrue.com/api_docs_includes_file/delivery.html#ticker
 * @param {string} symbol unified symbol of the market to fetch the ticker for
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {object} a [ticker structure]{@link https://docs.ccxt.com/#/?id=ticker-structure}
 */
func (this *Bitrue) FetchTicker(symbol string, options ...FetchTickerOptions) (Ticker, error) {

    opts := FetchTickerOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.FetchTicker(symbol, params)
    if IsError(res) {
        return Ticker{}, CreateReturnError(res)
    }
    return NewTicker(res), nil
}
/**
 * @method
 * @name bitrue#fetchOHLCV
 * @description fetches historical candlestick data containing the open, high, low, and close price, and the volume of a market
 * @see https://www.bitrue.com/api_docs_includes_file/spot/index.html#kline-data
 * @see https://www.bitrue.com/api_docs_includes_file/futures/index.html#kline-candlestick-data
 * @param {string} symbol unified symbol of the market to fetch OHLCV data for
 * @param {string} timeframe the length of time each candle represents
 * @param {int} [since] timestamp in ms of the earliest candle to fetch
 * @param {int} [limit] the maximum amount of candles to fetch
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @param {int} [params.until] the latest time in ms to fetch transfers for
 * @returns {int[][]} A list of candles ordered as timestamp, open, high, low, close, volume
 */
func (this *Bitrue) FetchOHLCV(symbol string, options ...FetchOHLCVOptions) ([]OHLCV, error) {

    opts := FetchOHLCVOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var timeframe interface{} = nil
    if opts.Timeframe != nil {
        timeframe = *opts.Timeframe
    }

    var since interface{} = nil
    if opts.Since != nil {
        since = *opts.Since
    }

    var limit interface{} = nil
    if opts.Limit != nil {
        limit = *opts.Limit
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.FetchOHLCV(symbol, timeframe, since, limit, params)
    if IsError(res) {
        return nil, CreateReturnError(res)
    }
    return NewOHLCVArray(res), nil
}
/**
 * @method
 * @name bitrue#fetchBidsAsks
 * @description fetches the bid and ask price and volume for multiple markets
 * @see https://github.com/Bitrue-exchange/Spot-official-api-docs#symbol-order-book-ticker
 * @see https://www.bitrue.com/api-docs#ticker
 * @see https://www.bitrue.com/api_docs_includes_file/delivery.html#ticker
 * @param {string[]|undefined} symbols unified symbols of the markets to fetch the bids and asks for, all markets are returned if not assigned
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {object} a dictionary of [ticker structures]{@link https://docs.ccxt.com/#/?id=ticker-structure}
 */
func (this *Bitrue) FetchBidsAsks(options ...FetchBidsAsksOptions) (Tickers, error) {

    opts := FetchBidsAsksOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var symbols interface{} = nil
    if opts.Symbols != nil {
        symbols = *opts.Symbols
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.FetchBidsAsks(symbols, params)
    if IsError(res) {
        return Tickers{}, CreateReturnError(res)
    }
    return NewTickers(res), nil
}
/**
 * @method
 * @name bitrue#fetchTickers
 * @description fetches price tickers for multiple markets, statistical information calculated over the past 24 hours for each market
 * @see https://github.com/Bitrue-exchange/Spot-official-api-docs#24hr-ticker-price-change-statistics
 * @see https://www.bitrue.com/api-docs#ticker
 * @see https://www.bitrue.com/api_docs_includes_file/delivery.html#ticker
 * @param {string[]|undefined} symbols unified symbols of the markets to fetch the ticker for, all market tickers are returned if not assigned
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {object} a dictionary of [ticker structures]{@link https://docs.ccxt.com/#/?id=ticker-structure}
 */
func (this *Bitrue) FetchTickers(options ...FetchTickersOptions) (Tickers, error) {

    opts := FetchTickersOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var symbols interface{} = nil
    if opts.Symbols != nil {
        symbols = *opts.Symbols
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.FetchTickers(symbols, params)
    if IsError(res) {
        return Tickers{}, CreateReturnError(res)
    }
    return NewTickers(res), nil
}
/**
 * @method
 * @name bitrue#fetchTrades
 * @description get the list of most recent trades for a particular symbol
 * @see https://github.com/Bitrue-exchange/Spot-official-api-docs#recent-trades-list
 * @param {string} symbol unified symbol of the market to fetch trades for
 * @param {int} [since] timestamp in ms of the earliest trade to fetch
 * @param {int} [limit] the maximum amount of trades to fetch
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {Trade[]} a list of [trade structures]{@link https://docs.ccxt.com/#/?id=public-trades}
 */
func (this *Bitrue) FetchTrades(symbol string, options ...FetchTradesOptions) ([]Trade, error) {

    opts := FetchTradesOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var since interface{} = nil
    if opts.Since != nil {
        since = *opts.Since
    }

    var limit interface{} = nil
    if opts.Limit != nil {
        limit = *opts.Limit
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.FetchTrades(symbol, since, limit, params)
    if IsError(res) {
        return nil, CreateReturnError(res)
    }
    return NewTradeArray(res), nil
}
/**
 * @method
 * @name bitrue#createMarketBuyOrderWithCost
 * @description create a market buy order by providing the symbol and cost
 * @see https://www.bitrue.com/api-docs#new-order-trade-hmac-sha256
 * @see https://www.bitrue.com/api_docs_includes_file/delivery.html#new-order-trade-hmac-sha256
 * @param {string} symbol unified symbol of the market to create an order in
 * @param {float} cost how much you want to trade in units of the quote currency
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {object} an [order structure]{@link https://docs.ccxt.com/#/?id=order-structure}
 */
func (this *Bitrue) CreateMarketBuyOrderWithCost(symbol string, cost float64, options ...CreateMarketBuyOrderWithCostOptions) (Order, error) {

    opts := CreateMarketBuyOrderWithCostOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.CreateMarketBuyOrderWithCost(symbol, cost, params)
    if IsError(res) {
        return Order{}, CreateReturnError(res)
    }
    return NewOrder(res), nil
}
/**
 * @method
 * @name bitrue#createOrder
 * @description create a trade order
 * @see https://www.bitrue.com/api_docs_includes_file/spot/index.html#new-order-trade
 * @see https://www.bitrue.com/api_docs_includes_file/futures/index.html#new-order-trade-hmac-sha256
 * @param {string} symbol unified symbol of the market to create an order in
 * @param {string} type 'market' or 'limit'
 * @param {string} side 'buy' or 'sell'
 * @param {float} amount how much of currency you want to trade in units of base currency
 * @param {float} [price] the price at which the order is to be fulfilled, in units of the quote currency, ignored in market orders
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @param {float} [params.triggerPrice] *spot only* the price at which a trigger order is triggered at
 * @param {string} [params.clientOrderId] a unique id for the order, automatically generated if not sent
 * @param {decimal} [params.leverage] in future order, the leverage value of the order should consistent with the user contract configuration, default is 1
 * @param {string} [params.timeInForce] 'fok', 'ioc' or 'po'
 * @param {bool} [params.postOnly] default false
 * @param {bool} [params.reduceOnly] default false
 * EXCHANGE SPECIFIC PARAMETERS
 * @param {decimal} [params.icebergQty]
 * @param {long} [params.recvWindow]
 * @param {float} [params.cost] *swap market buy only* the quote quantity that can be used as an alternative for the amount
 * @returns {object} an [order structure]{@link https://docs.ccxt.com/#/?id=order-structure}
 */
func (this *Bitrue) CreateOrder(symbol string, typeVar string, side string, amount float64, options ...CreateOrderOptions) (Order, error) {

    opts := CreateOrderOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var price interface{} = nil
    if opts.Price != nil {
        price = *opts.Price
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.CreateOrder(symbol, typeVar, side, amount, price, params)
    if IsError(res) {
        return Order{}, CreateReturnError(res)
    }
    return NewOrder(res), nil
}
/**
 * @method
 * @name bitrue#fetchOrder
 * @description fetches information on an order made by the user
 * @see https://www.bitrue.com/api_docs_includes_file/spot/index.html#query-order-user_data
 * @see https://www.bitrue.com/api_docs_includes_file/futures/index.html#query-order-user_data-hmac-sha256
 * @param {string} id the order id
 * @param {string} symbol unified symbol of the market the order was made in
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {object} An [order structure]{@link https://docs.ccxt.com/#/?id=order-structure}
 */
func (this *Bitrue) FetchOrder(id string, options ...FetchOrderOptions) (Order, error) {

    opts := FetchOrderOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var symbol interface{} = nil
    if opts.Symbol != nil {
        symbol = *opts.Symbol
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.FetchOrder(id, symbol, params)
    if IsError(res) {
        return Order{}, CreateReturnError(res)
    }
    return NewOrder(res), nil
}
/**
 * @method
 * @name bitrue#fetchClosedOrders
 * @description fetches information on multiple closed orders made by the user
 * @see https://www.bitrue.com/api_docs_includes_file/spot/index.html#all-orders-user_data
 * @param {string} symbol unified market symbol of the market orders were made in
 * @param {int} [since] the earliest time in ms to fetch orders for
 * @param {int} [limit] the maximum number of order structures to retrieve
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {Order[]} a list of [order structures]{@link https://docs.ccxt.com/#/?id=order-structure}
 */
func (this *Bitrue) FetchClosedOrders(options ...FetchClosedOrdersOptions) ([]Order, error) {

    opts := FetchClosedOrdersOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var symbol interface{} = nil
    if opts.Symbol != nil {
        symbol = *opts.Symbol
    }

    var since interface{} = nil
    if opts.Since != nil {
        since = *opts.Since
    }

    var limit interface{} = nil
    if opts.Limit != nil {
        limit = *opts.Limit
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.FetchClosedOrders(symbol, since, limit, params)
    if IsError(res) {
        return nil, CreateReturnError(res)
    }
    return NewOrderArray(res), nil
}
/**
 * @method
 * @name bitrue#fetchOpenOrders
 * @description fetch all unfilled currently open orders
 * @see https://www.bitrue.com/api_docs_includes_file/spot/index.html#current-open-orders-user_data
 * @see https://www.bitrue.com/api_docs_includes_file/futures/index.html#cancel-all-open-orders-trade-hmac-sha256
 * @param {string} symbol unified market symbol
 * @param {int} [since] the earliest time in ms to fetch open orders for
 * @param {int} [limit] the maximum number of open order structures to retrieve
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {Order[]} a list of [order structures]{@link https://docs.ccxt.com/#/?id=order-structure}
 */
func (this *Bitrue) FetchOpenOrders(options ...FetchOpenOrdersOptions) ([]Order, error) {

    opts := FetchOpenOrdersOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var symbol interface{} = nil
    if opts.Symbol != nil {
        symbol = *opts.Symbol
    }

    var since interface{} = nil
    if opts.Since != nil {
        since = *opts.Since
    }

    var limit interface{} = nil
    if opts.Limit != nil {
        limit = *opts.Limit
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.FetchOpenOrders(symbol, since, limit, params)
    if IsError(res) {
        return nil, CreateReturnError(res)
    }
    return NewOrderArray(res), nil
}
/**
 * @method
 * @name bitrue#cancelOrder
 * @description cancels an open order
 * @see https://github.com/Bitrue-exchange/Spot-official-api-docs#cancel-order-trade
 * @see https://www.bitrue.com/api-docs#cancel-order-trade-hmac-sha256
 * @see https://www.bitrue.com/api_docs_includes_file/delivery.html#cancel-order-trade-hmac-sha256
 * @param {string} id order id
 * @param {string} symbol unified symbol of the market the order was made in
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {object} An [order structure]{@link https://docs.ccxt.com/#/?id=order-structure}
 */
func (this *Bitrue) CancelOrder(id string, options ...CancelOrderOptions) (Order, error) {

    opts := CancelOrderOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var symbol interface{} = nil
    if opts.Symbol != nil {
        symbol = *opts.Symbol
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.CancelOrder(id, symbol, params)
    if IsError(res) {
        return Order{}, CreateReturnError(res)
    }
    return NewOrder(res), nil
}
/**
 * @method
 * @name bitrue#cancelAllOrders
 * @description cancel all open orders in a market
 * @see https://www.bitrue.com/api-docs#cancel-all-open-orders-trade-hmac-sha256
 * @see https://www.bitrue.com/api_docs_includes_file/delivery.html#cancel-all-open-orders-trade-hmac-sha256
 * @param {string} symbol unified market symbol of the market to cancel orders in
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @param {string} [params.marginMode] 'cross' or 'isolated', for spot margin trading
 * @returns {object[]} a list of [order structures]{@link https://github.com/ccxt/ccxt/wiki/Manual#order-structure}
 */
func (this *Bitrue) CancelAllOrders(options ...CancelAllOrdersOptions) ([]Order, error) {

    opts := CancelAllOrdersOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var symbol interface{} = nil
    if opts.Symbol != nil {
        symbol = *opts.Symbol
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.CancelAllOrders(symbol, params)
    if IsError(res) {
        return nil, CreateReturnError(res)
    }
    return NewOrderArray(res), nil
}
/**
 * @method
 * @name bitrue#fetchMyTrades
 * @description fetch all trades made by the user
 * @see https://www.bitrue.com/api_docs_includes_file/spot/index.html#account-trade-list-user_data
 * @see https://www.bitrue.com/api_docs_includes_file/futures/index.html#account-trade-list-user_data-hmac-sha256
 * @param {string} symbol unified market symbol
 * @param {int} [since] the earliest time in ms to fetch trades for
 * @param {int} [limit] the maximum number of trades structures to retrieve
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {Trade[]} a list of [trade structures]{@link https://docs.ccxt.com/#/?id=trade-structure}
 */
func (this *Bitrue) FetchMyTrades(options ...FetchMyTradesOptions) ([]Trade, error) {

    opts := FetchMyTradesOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var symbol interface{} = nil
    if opts.Symbol != nil {
        symbol = *opts.Symbol
    }

    var since interface{} = nil
    if opts.Since != nil {
        since = *opts.Since
    }

    var limit interface{} = nil
    if opts.Limit != nil {
        limit = *opts.Limit
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.FetchMyTrades(symbol, since, limit, params)
    if IsError(res) {
        return nil, CreateReturnError(res)
    }
    return NewTradeArray(res), nil
}
/**
 * @method
 * @name bitrue#fetchDeposits
 * @description fetch all deposits made to an account
 * @see https://github.com/Bitrue-exchange/Spot-official-api-docs#deposit-history--withdraw_data
 * @param {string} code unified currency code
 * @param {int} [since] the earliest time in ms to fetch deposits for
 * @param {int} [limit] the maximum number of deposits structures to retrieve
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {object[]} a list of [transaction structures]{@link https://docs.ccxt.com/#/?id=transaction-structure}
 */
func (this *Bitrue) FetchDeposits(options ...FetchDepositsOptions) ([]Transaction, error) {

    opts := FetchDepositsOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var code interface{} = nil
    if opts.Code != nil {
        code = *opts.Code
    }

    var since interface{} = nil
    if opts.Since != nil {
        since = *opts.Since
    }

    var limit interface{} = nil
    if opts.Limit != nil {
        limit = *opts.Limit
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.FetchDeposits(code, since, limit, params)
    if IsError(res) {
        return nil, CreateReturnError(res)
    }
    return NewTransactionArray(res), nil
}
/**
 * @method
 * @name bitrue#fetchWithdrawals
 * @description fetch all withdrawals made from an account
 * @see https://github.com/Bitrue-exchange/Spot-official-api-docs#withdraw-history--withdraw_data
 * @param {string} code unified currency code
 * @param {int} [since] the earliest time in ms to fetch withdrawals for
 * @param {int} [limit] the maximum number of withdrawals structures to retrieve
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {object[]} a list of [transaction structures]{@link https://docs.ccxt.com/#/?id=transaction-structure}
 */
func (this *Bitrue) FetchWithdrawals(options ...FetchWithdrawalsOptions) ([]Transaction, error) {

    opts := FetchWithdrawalsOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var code interface{} = nil
    if opts.Code != nil {
        code = *opts.Code
    }

    var since interface{} = nil
    if opts.Since != nil {
        since = *opts.Since
    }

    var limit interface{} = nil
    if opts.Limit != nil {
        limit = *opts.Limit
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.FetchWithdrawals(code, since, limit, params)
    if IsError(res) {
        return nil, CreateReturnError(res)
    }
    return NewTransactionArray(res), nil
}
/**
 * @method
 * @name bitrue#withdraw
 * @description make a withdrawal
 * @see https://github.com/Bitrue-exchange/Spot-official-api-docs#withdraw-commit--withdraw_data
 * @param {string} code unified currency code
 * @param {float} amount the amount to withdraw
 * @param {string} address the address to withdraw to
 * @param {string} tag
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {object} a [transaction structure]{@link https://docs.ccxt.com/#/?id=transaction-structure}
 */
func (this *Bitrue) Withdraw(code string, amount float64, address string, options ...WithdrawOptions) (Transaction, error) {

    opts := WithdrawOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var tag interface{} = nil
    if opts.Tag != nil {
        tag = *opts.Tag
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.Withdraw(code, amount, address, tag, params)
    if IsError(res) {
        return Transaction{}, CreateReturnError(res)
    }
    return NewTransaction(res), nil
}
/**
 * @method
 * @name bitrue#fetchDepositWithdrawFees
 * @description fetch deposit and withdraw fees
 * @see https://github.com/Bitrue-exchange/Spot-official-api-docs#exchangeInfo_endpoint
 * @param {string[]|undefined} codes list of unified currency codes
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {object} a list of [fee structures]{@link https://docs.ccxt.com/#/?id=fee-structure}
 */
func (this *Bitrue) FetchDepositWithdrawFees(options ...FetchDepositWithdrawFeesOptions) (map[string]interface{}, error) {

    opts := FetchDepositWithdrawFeesOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var codes interface{} = nil
    if opts.Codes != nil {
        codes = *opts.Codes
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.FetchDepositWithdrawFees(codes, params)
    if IsError(res) {
        return map[string]interface{}{}, CreateReturnError(res)
    }
    return res.(map[string]interface{}), nil
}
/**
 * @method
 * @name bitrue#fetchTransfers
 * @description fetch a history of internal transfers made on an account
 * @see https://www.bitrue.com/api-docs#get-future-account-transfer-history-list-user_data-hmac-sha256
 * @see https://www.bitrue.com/api_docs_includes_file/delivery.html#get-future-account-transfer-history-list-user_data-hmac-sha256
 * @param {string} code unified currency code of the currency transferred
 * @param {int} [since] the earliest time in ms to fetch transfers for
 * @param {int} [limit] the maximum number of transfers structures to retrieve
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @param {int} [params.until] the latest time in ms to fetch transfers for
 * @param {string} [params.type] transfer type wallet_to_contract or contract_to_wallet
 * @returns {object[]} a list of [transfer structures]{@link https://github.com/ccxt/ccxt/wiki/Manual#transfer-structure}
 */
func (this *Bitrue) FetchTransfers(options ...FetchTransfersOptions) ([]TransferEntry, error) {

    opts := FetchTransfersOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var code interface{} = nil
    if opts.Code != nil {
        code = *opts.Code
    }

    var since interface{} = nil
    if opts.Since != nil {
        since = *opts.Since
    }

    var limit interface{} = nil
    if opts.Limit != nil {
        limit = *opts.Limit
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.FetchTransfers(code, since, limit, params)
    if IsError(res) {
        return nil, CreateReturnError(res)
    }
    return NewTransferEntryArray(res), nil
}
/**
 * @method
 * @name bitrue#transfer
 * @description transfer currency internally between wallets on the same account
 * @see https://www.bitrue.com/api-docs#new-future-account-transfer-user_data-hmac-sha256
 * @see https://www.bitrue.com/api_docs_includes_file/delivery.html#user-commission-rate-user_data-hmac-sha256
 * @param {string} code unified currency code
 * @param {float} amount amount to transfer
 * @param {string} fromAccount account to transfer from
 * @param {string} toAccount account to transfer to
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {object} a [transfer structure]{@link https://github.com/ccxt/ccxt/wiki/Manual#transfer-structure}
 */
func (this *Bitrue) Transfer(code string, amount float64, fromAccount string, toAccount string, options ...TransferOptions) (TransferEntry, error) {

    opts := TransferOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.Transfer(code, amount, fromAccount, toAccount, params)
    if IsError(res) {
        return TransferEntry{}, CreateReturnError(res)
    }
    return NewTransferEntry(res), nil
}
/**
 * @method
 * @name bitrue#setLeverage
 * @description set the level of leverage for a market
 * @see https://www.bitrue.com/api-docs#change-initial-leverage-trade-hmac-sha256
 * @see https://www.bitrue.com/api_docs_includes_file/delivery.html#change-initial-leverage-trade-hmac-sha256
 * @param {float} leverage the rate of leverage
 * @param {string} symbol unified market symbol
 * @param {object} [params] extra parameters specific to the exchange API endpoint
 * @returns {object} response from the exchange
 */
func (this *Bitrue) SetLeverage(leverage int64, options ...SetLeverageOptions) (map[string]interface{}, error) {

    opts := SetLeverageOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var symbol interface{} = nil
    if opts.Symbol != nil {
        symbol = *opts.Symbol
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.SetLeverage(leverage, symbol, params)
    if IsError(res) {
        return map[string]interface{}{}, CreateReturnError(res)
    }
    return res.(map[string]interface{}), nil
}
/**
 * @method
 * @name bitrue#setMargin
 * @description Either adds or reduces margin in an isolated position in order to set the margin to a specific value
 * @see https://www.bitrue.com/api-docs#modify-isolated-position-margin-trade-hmac-sha256
 * @see https://www.bitrue.com/api_docs_includes_file/delivery.html#modify-isolated-position-margin-trade-hmac-sha256
 * @param {string} symbol unified market symbol of the market to set margin in
 * @param {float} amount the amount to set the margin to
 * @param {object} [params] parameters specific to the exchange API endpoint
 * @returns {object} A [margin structure]{@link https://github.com/ccxt/ccxt/wiki/Manual#add-margin-structure}
 */
func (this *Bitrue) SetMargin(symbol string, amount float64, options ...SetMarginOptions) (MarginModification, error) {

    opts := SetMarginOptionsStruct{}

    for _, opt := range options {
        opt(&opts)
    }

    var params interface{} = nil
    if opts.Params != nil {
        params = *opts.Params
    }
    res := <- this.Core.SetMargin(symbol, amount, params)
    if IsError(res) {
        return MarginModification{}, CreateReturnError(res)
    }
    return NewMarginModification(res), nil
}