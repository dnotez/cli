package main

import (
	"fmt"
	flag "github.com/dnotez/cli/mflag"
	"time"
)

func (cli *PlCli) CmdHelp(args ...string) error {
	if len(args) > 1 {
		method, exists := cli.getMethod(args[:2]...)
		if exists {
			method("--help")
			return nil
		}
	}
	if len(args) > 0 {
		method, exists := cli.getMethod(args[0])
		if !exists {
			fmt.Fprintf(cli.err, "Error: Command not found: %s\n", args[0])
		} else {
			method("--help")
			return nil
		}
	}
	flag.Usage()
	return nil
}

func (cli *PlCli) CmdSearch(args ...string) error {
	cmd := cli.SubCmd("search", "query", "Search for `query` in the saved articles")
	articleType := cmd.String([]string{"t", "-type"}, "", "Type of the article (e.g. bash, article, answer)")

	if err := cmd.Parse(args); err != nil {
		return nil
	}
	if cmd.NArg() != 1 || len(cmd.Arg(0)) < 1 {
		cmd.Usage()
		return nil
	}
	response, duration, err := suggestion.Suggest(cmd.Arg(0), *articleType)
	if err != nil {
		return err
	}
	fmt.Fprintf(cli.out, "Results (in %0.2fs)\n%s\n", duration.Seconds(), response)
	return nil
}

func (cli *PlCli) CmdSave(args ...string) error {
	saveCmd := cli.SubCmd("save", "line", "Save current line in the system")
	label := saveCmd.String([]string{"l", "-label"}, "", "set label for quick fetch")

	if err := saveCmd.Parse(args); err != nil {
		return nil
	}
	if saveCmd.NArg() < 1 || len(saveCmd.Arg(0)) < 1 {
		saveCmd.Usage()
		return nil
	}
	cmdText := saveCmd.Arg(0)

	request := cmd.SaveCmdRequest{Body: cmdText, Label: *label}
	response, duration, err := request.Submit()
	if err != nil {
		return err
	}
	if cli.verbose {
		fmt.Fprintf(cli.out, "Saved in %0.2fs\n", duration.Seconds())
	}
	fmt.Fprintf(cli.out, "%s\n", (*response).URL)
	return nil
}

func (cli *PlCli) CmdShare(args ...string) error {
	shareCmd := cli.SubCmd("share", "id", "Share a saved resource with with the given id with other team members.")
	shareWith := shareCmd.String([]string{"u", "-user"}, "", "Comma separated user names, use '*' to share with all team members.")

	if err := shareCmd.Parse(args); err != nil {
		return nil
	}
	if shareCmd.NArg() < 1 || len(shareCmd.Arg(0)) < 1 {
		shareCmd.Usage()
		return nil
	}
	fmt.Fprintf(cli.out, "Sharing resource:\"%s\", with:%s\n", shareCmd.Arg(0), *shareWith)
	return nil
}

/**
 * Remove a resource by id
 */
func (cli *PlCli) CmdRm(args ...string) error {
	deleteCmd := cli.SubCmd("rm", "id", "Remove a resource by id.")
	if err := deleteCmd.Parse(args); err != nil {
		return nil
	}
	if deleteCmd.NArg() < 1 || len(deleteCmd.Arg(0)) < 1 {
		deleteCmd.Usage()
		return nil
	}
	id := deleteCmd.Arg(0)
	_, duration, err := article.Remove(id)
	if err != nil {
		//fmt.Fprintf(cli.err, "Error in removing resource:%s\nError:%s\n", id, err);
		return err
	}
	if cli.verbose {
		fmt.Fprintf(cli.out, "Removed in %0.2fs\n", duration.Seconds())
	}
	return nil
}

func (cli *PlCli) CmdGet(args ...string) error {
	getCmd := cli.SubCmd("get", "key", "Get a resource by a key")
	key := getCmd.String([]string{"k", "-key"}, "id", "Key field for getting the article (can be id or label)")
	field := getCmd.String([]string{"f", "-field"}, "text", "Output field ( can be id, html, text, title or all)")
	size := getCmd.Int([]string{"n", "-num"}, 1, "Number of resources to return")
	if err := getCmd.Parse(args); err != nil {
		return nil
	}
	if getCmd.NArg() < 1 || len(getCmd.Arg(0)) < 1 {
		getCmd.Usage()
		return nil
	}

	articles, duration, err := article.Get(getCmd.Arg(0), *key, *size)
	if err != nil {
		return err
	}
	if cli.verbose {
		fmt.Fprintf(cli.out, "Fetched in %0.2fs\n", duration.Seconds())
	}

	for _, article := range *articles {
		switch *field {
		case "text":
			fmt.Fprintf(cli.out, "%s\n", article.Text)
		case "title":
			fmt.Fprintf(cli.out, "%s\n", article.Title)
		case "id":
			fmt.Fprintf(cli.out, "%s\n", article.Id)
		case "label":
			fmt.Fprintf(cli.out, "%s\n", article.Label)
		default:
			fmt.Fprintf(cli.out, "%s\n", article.Id)
			fmt.Fprintf(cli.out, "%s\n", article.Title)
			fmt.Fprintf(cli.out, "%s\n", article.Text)
			fmt.Fprintf(cli.out, "%s\n", article.Label)
			fmt.Fprintf(cli.out, "%s\n", time.Unix(article.SaveDate/1000, 0).Local())
			fmt.Fprint(cli.out, "\n")
		}
	}

	//fmt.Fprintf(cli.err, "Not implemented yet. (field:%s, key:%s, size:%d)value:%s \n", *field, *key, *size, getCmd.Arg(0))
	return nil
}
