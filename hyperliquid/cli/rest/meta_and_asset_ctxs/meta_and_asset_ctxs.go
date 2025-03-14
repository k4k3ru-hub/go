//
// meta_and_asset_ctxs.go
//
package meta_and_asset_ctxs

import (
	"fmt"
	"strconv"
	"strings"

	myRest "github.com/k4k3ru-hub/go/hyperliquid/rest"
	myRestInfoMetaAndAssetCtxs "github.com/k4k3ru-hub/go/hyperliquid/rest/info/meta_and_asset_ctxs"

	"github.com/k4k3ru-hub/go/cli"
)


const (
	OptionNameToken = "token"
	OptionAliasToken = "t"
)


//
// Run.
//
func Run(options map[string]*cli.Option) {
	fmt.Printf("Started rest metaAndAssetCtxs command.\n")

	// Send API request.
	c := myRestInfoMetaAndAssetCtxs.NewMetaAndAssetCtxsClient()
    result, err := c.Send()
    if err != nil {
        fmt.Printf("%s\n", err)
        return
    }

	// Check the token option.
	var token string
	if tokenOption, ok := options[OptionNameToken]; ok {
		token = tokenOption.Value
	}

	// Optimize data.
	header := []string{
		"Name",
		"SzDecimals",
		"MaxLeverage",
		"OnlyIsolated",
		"DayNtlVlm",
		"Funding",
		"ImpactPxs",
		"MarkPx",
		"MidPx",
		"OpenInterest",
		"OpenInterest",
		"OraclePx",
		"Premium",
		"PrevDayPx",
	}
	colWidths := make([]int, len(header))
	for i, h := range header {
		colWidths[i] = len(h)+2
	}
	var data [][]string
	for _, universeEntry := range result.Universe {
		// Check the column width.
		if len(universeEntry.Name) > colWidths[0] {
			colWidths[0] = len(universeEntry.Name)
		}
		if len(strconv.Itoa(universeEntry.SzDecimals)) > colWidths[1] {
			colWidths[1] = len(strconv.Itoa(universeEntry.SzDecimals))
		}
		if len(strconv.Itoa(universeEntry.MaxLeverage)) > colWidths[2] {
            colWidths[2] = len(strconv.Itoa(universeEntry.MaxLeverage))
        }

		// Set data.
		rowData := []string{universeEntry.Name, strconv.Itoa(universeEntry.SzDecimals), strconv.Itoa(universeEntry.MaxLeverage), strconv.FormatBool(universeEntry.OnlyIsolated)}

		// Asset entries.
		for _, assetEntry := range result.Assets {
			// Check the column width.
			if len(assetEntry.DayNtlVlm)+2 > colWidths[4] {
				colWidths[4] = len(assetEntry.DayNtlVlm)+2
			}
			if len(assetEntry.Funding)+2 > colWidths[5] {
				colWidths[5] = len(assetEntry.Funding)+2
			}
			if len(strings.Join(assetEntry.ImpactPxs, " "))+2 > colWidths[6] {
				colWidths[6] = len(strings.Join(assetEntry.ImpactPxs, " "))+2
			}
			if len(assetEntry.MarkPx)+2 > colWidths[7] {
				colWidths[7] = len(assetEntry.MarkPx)+2
			}
			if len(assetEntry.MidPx)+2 > colWidths[8] {
				colWidths[8] = len(assetEntry.MidPx)+2
			}
			if len(assetEntry.OpenInterest)+2 > colWidths[9] {
				colWidths[9] = len(assetEntry.OpenInterest)+2
			}
			if len(assetEntry.OraclePx)+2 > colWidths[10] {
				colWidths[10] = len(assetEntry.OraclePx)+2
			}
			if len(assetEntry.Premium)+2 > colWidths[11] {
				colWidths[11] = len(assetEntry.Premium)+2
			}
			if len(assetEntry.PrevDayPx)+2 > colWidths[12] {
				colWidths[12] = len(assetEntry.PrevDayPx)+2
			}

			// Set data.
			rowData = append(rowData, assetEntry.DayNtlVlm, assetEntry.Funding, strings.Join(assetEntry.ImpactPxs, " "), assetEntry.MarkPx, assetEntry.MidPx, assetEntry.OpenInterest, assetEntry.OraclePx, assetEntry.Premium, assetEntry.PrevDayPx)
		}

		data = append(data, rowData)
	}

	// Output for table format.
	format := ""
	var line strings.Builder
	for _, w := range colWidths {
		format += fmt.Sprintf("%%-%ds", w)
		line.WriteString(strings.Repeat("-", w))
	}
	fmt.Printf(format+"\n", header[0], header[1], header[2], header[3], header[4], header[5], header[6], header[7], header[8], header[9], header[10], header[11], header[12], header[13])
	fmt.Printf("%s\n", line.String())
	for _, row := range data {
		if token == "" || strings.ToUpper(token) == row[0] {
			fmt.Printf(format+"\n", row[0], row[1], row[2], row[3], row[4], row[5], row[6], row[7], row[8], row[9], row[10], row[11], row[12], row[13])
		}
	}


/*
    for _, universeEntry := range result.Universe {
        fmt.Printf("Universe:\n")
        fmt.Printf("- Name: %s\n", universeEntry.Name)
        fmt.Printf("- SzDecimals: %d\n", universeEntry.SzDecimals)
        fmt.Printf("- MaxLeverage: %d\n", universeEntry.MaxLeverage)
        fmt.Printf("- OnlyIsolated: %t\n", universeEntry.OnlyIsolated)
    }
    ethAsset := result.GetAssetByName("ETH")
    fmt.Printf("ETH Asset Entry:\n")
    fmt.Printf("DayNtlVlm: %s\n", ethAsset.DayNtlVlm)
    fmt.Printf("Funding: %s\n", ethAsset.Funding)
    fmt.Printf("ImpactPxs: %v\n", ethAsset.ImpactPxs)
    fmt.Printf("MarkPx: %s\n", ethAsset.MarkPx)
    fmt.Printf("MidPx: %s\n", ethAsset.MidPx)
    fmt.Printf("OpenInterest: %s\n", ethAsset.OpenInterest)
    fmt.Printf("OraclePx: %s\n", ethAsset.OraclePx)
    fmt.Printf("Premium: %s\n", ethAsset.Premium)
    fmt.Printf("PrevDayPx: %s\n", ethAsset.PrevDayPx)
*/
}


//
// Set command.
//
func SetCommand(parentCommand *cli.Command) {
	command := cli.NewCommand(myRest.TypeMetaAndAssetCtxs)
    parentCommand.Commands = append(parentCommand.Commands, command)
    command.Options[OptionNameToken] = &cli.Option{
        Alias: OptionAliasToken,
        HasValue: true,
    }
    command.Action = Run
}

