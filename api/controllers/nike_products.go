package controllers

//
//func (c *Controllers) GetProductNikeByProductID(m *discordgo.MessageCreate)  {
//	info := strings.Split(m.Content, " ")
//	if len(info) < 2 {
//		c.Logger.Error(errors.WithStack(errors.New("Missing parameter: product id")))
//		c.BadAction("Missing parameter: product id", m)
//		return
//	}
//	productID := info[1]
//	product, err := c.GsProductNike.GetProductByID(productID)
//	if err != nil {
//		c.Logger.Error(err)
//		c.BadAction(err.Error(), m)
//		return
//	}
//	styleID := discordgo.MessageEmbedField{
//		Name:   "Style ID",
//		Value:   product.StyleID,
//		Inline: false,
//	}
//	price := discordgo.MessageEmbedField{
//		Name:   "Price",
//		Value:  fmt.Sprintf("%d %s", int(product.Price), product.CurrencyCode),
//		Inline:  true,
//	}
//	storeID := discordgo.MessageEmbedField{
//		Name:   "Store ID",
//		Value:  "Nike RU",
//		Inline: true,
//	}
//	var fields []*discordgo.MessageEmbedField
//	size := 4
//	var j int
//	for i := 0; i < len(product.Sizes); i += size{
//		j += size
//		if j > len(product.Sizes) {
//			j = len(product.Sizes)
//		}
//		// do what do you want to with the sub-slice, here just printing the sub-slices
//		smallSlice := product.Sizes[i:j]
//
//		newSizesField := discordgo.MessageEmbedField{
//			Name: "Sizes",
//			Value: strings.Join(smallSlice[:],"\n"),
//			Inline: false,
//		}
//		fields = append(fields, &newSizesField)
//	}
//}