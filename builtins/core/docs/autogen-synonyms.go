package docs

//Synonym is used for builtins that might have more than one internal alias
var Synonym map[string]string = map[string]string{
	`(`:               `brace-quote`,
	`echo`:            `out`,
	`!and`:            `and`,
	`!catch`:          `catch`,
	`!export`:         `export`,
	`unset`:           `export`,
	`!set`:            `set`,
	`!or`:             `or`,
	`!if`:             `if`,
	`!global`:         `global`,
	`!event`:          `event`,
	`murex-docs`:      `murex-docs`,
	`err`:             `err`,
	`ttyfd`:           `ttyfd`,
	`if`:              `if`,
	`>>`:              `>>`,
	`rx`:              `rx`,
	`read`:            `read`,
	`or`:              `or`,
	`swivel-table`:    `swivel-table`,
	`get`:             `get`,
	`out`:             `out`,
	`tout`:            `tout`,
	`try`:             `try`,
	`brace-quote`:     `brace-quote`,
	`>`:               `>`,
	`trypipe`:         `trypipe`,
	`append`:          `append`,
	`getfile`:         `getfile`,
	`f`:               `f`,
	`event`:           `event`,
	`pt`:              `pt`,
	`tread`:           `tread`,
	`and`:             `and`,
	`export`:          `export`,
	`prepend`:         `prepend`,
	`swivel-datatype`: `swivel-datatype`,
	`set`:             `set`,
	`global`:          `global`,
	`alter`:           `alter`,
	`post`:            `post`,
	`g`:               `g`,
	`catch`:           `catch`,
}
