grammar PBXProj;

// PARSER

project: fileEncodeInfo Comment? valMap;

fileEncodeInfo: '//' '!$*' NString '*$!';

valMap: '{' (Comment? key '=' value ';' Comment?)* '}';
valArray: '(' (value ',')* value? ')';
valString: NString Comment?;

key: valString;
value: valArray
    | valMap
    | valString
    ;

// LEXER

Comment: '/*' .*? '*/';
NString: String1 | String2;

fragment String1: '"' (ESC|.)*? '"';
fragment String2: [a-zA-Z0-9.@?_/\-]+;
fragment ESC : '\\"' | '\\\\' ;

WS: [ \t\r\n\u000C]+ -> skip;
