package docs

//Synonym is used for builtins that might have more than one internal alias
var Synonym map[string]string = map[string]string{
	`!and`:            `and`,
	`!or`:             `or`,
	`!if`:             `if`,
	`!catch`:          `catch`,
	`!export`:         `export`,
	`unset`:           `export`,
	`!set`:            `set`,
	`!event`:          `event`,
	`(`:               `brace-quote`,
	`echo`:            `out`,
	`!global`:         `global`,
	`ttyfd`:           `ttyfd`,
	`tread`:           `tread`,
	`err`:             `err`,
	`out`:             `out`,
	`f`:               `f`,
	`rx`:              `rx`,
	`swivel-datatype`: `swivel-datatype`,
	`export`:          `export`,
	`trypipe`:         `trypipe`,
	`prepend`:         `prepend`,
	`getfile`:         `getfile`,
	`brace-quote`:     `brace-quote`,
	`tout`:            `tout`,
	`>`:               `>`,
	`pt`:              `pt`,
	`read`:            `read`,
	`append`:          `append`,
	`try`:             `try`,
	`or`:              `or`,
	`murex-docs`:      `murex-docs`,
	`g`:               `g`,
	`and`:             `and`,
	`set`:             `set`,
	`swivel-table`:    `swivel-table`,
	`get`:             `get`,
	`if`:              `if`,
	`global`:          `global`,
	`alter`:           `alter`,
	`catch`:           `catch`,
	`>>`:              `>>`,
	`event`:           `event`,
	`post`:            `post`,
}
