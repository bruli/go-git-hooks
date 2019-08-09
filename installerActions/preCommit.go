package installerActions

import (
	"github.com/bruli/go-git-hooks/configuration"
	"github.com/tockins/interact"
)

func PreCommitAction(conf *configuration.PreCommitConfiguration) error  {
	err := interact.Run(&interact.Interact{
		Questions: []*interact.Question{
			{
				Quest: interact.Quest{
					Msg: "Would you like enable pre-commit hook?.",
				},
				Action: func(c interact.Context) interface{} {
					val, err := c.Ans().Bool()
					if err != nil {
						return err
					}
					if true == val {
						if err := questionsTool(conf); err != nil {
							return err
						}
					}
					return nil
				},
			},
		},
		After:  nil,
		Before: nil,
	})

	return err
}

func questionsTool(conf *configuration.PreCommitConfiguration) error  {
	t := configuration.Tool{}
	err := interact.Run(&interact.Interact{
		Questions: []*interact.Question{
			{
				Quest: interact.Quest{
					Msg: "Please write a title for pre-commit tool:",
				},
				Action: func(context interact.Context) interface{} {
					title, err := context.Ans().String()
					if err != nil {
						return err
					}
					t.Title = title
					return nil

				},
			},
			{
				Quest: interact.Quest{
					Msg: "Please write command to execute:",
				},
				Action: func(context interact.Context) interface{} {
					com, err := context.Ans().String()
					if err != nil {
						return err
					}
					t.Command = com
					conf.AddTool(&t)
					return nil
				},
			},
			{
				Quest: interact.Quest{
					Msg: "Do you want add other pre-commit tool?:",
				},
				Action: func(context interact.Context) interface{} {
					ans, err := context.Ans().Bool()
					if err != nil {
						return err
					}
					if true == ans {
						if err := questionsTool(conf); err != nil {
							return err
						}
					}
					return nil
				},
			},
		},
	})
	if err != nil {
		return err
	}

	return nil
}