grammar PBXProj;

// PARSER

project: fileEncodeInfo Comment? valMap;

fileEncodeInfo: '//' '!$*' NString '*$!';

valMap: '{' ((sectionName)+ | sectionNoName) '}';
valArray: '(' (value ',')* value? ')';
valString: NString Comment?;

key: valString;
value: valArray
    | valMap
    | valString
    ;

sectionName: SectionB
        (key '=' value ';')+
        SectionE;
sectionNoName: (key '=' value ';')*;

// LEXER

SectionB: '/* Begin' .*? 'section */';
SectionE: '/* End' .*? 'section */';
Comment: '/*' .*? '*/';
NString: String1 | String2;

fragment String1: '"' (ESC|.)*? '"';
fragment String2: [a-zA-Z0-9.@?_/\-]+;
fragment ESC : '\\"' | '\\\\' ;

WS: [ \t\r\n\u000C]+ -> skip;
