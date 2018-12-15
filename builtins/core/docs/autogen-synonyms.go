package docs

//Synonym is used for builtins that might have more than one internal alias
var Synonym map[string]string = map[string]string{
	`!if`:             `if`,
	`!catch`:          `catch`,
	`!export`:         `export`,
	`unset`:           `export`,
	`!global`:         `global`,
	`!event`:          `event`,
	`!or`:             `or`,
	`echo`:            `out`,
	`!and`:            `and`,
	`!set`:            `set`,
	`(`:               `brace-quote`,
	`swivel-datatype`: `swivel-datatype`,
	`post`:            `post`,
	`brace-quote`:     `brace-quote`,
	`err`:             `err`,
	`out`:             `out`,
	`read`:            `read`,
	`alter`:           `alter`,
	`>`:               `>`,
	`>>`:              `>>`,
	`if`:              `if`,
	`trypipe`:         `trypipe`,
	`prepend`:         `prepend`,
	`get`:             `get`,
	`getfile`:         `getfile`,
	`try`:             `try`,
	`global`:          `global`,
	`append`:          `append`,
	`tout`:            `tout`,
	`pt`:              `pt`,
	`f`:               `f`,
	`and`:             `and`,
	`or`:              `or`,
	`catch`:           `catch`,
	`murex-docs`:      `murex-docs`,
	`ttyfd`:           `ttyfd`,
	`rx`:              `rx`,
	`set`:             `set`,
	`event`:           `event`,
	`g`:               `g`,
	`tread`:           `tread`,
	`swivel-table`:    `swivel-table`,
	`export`:          `export`,
}
