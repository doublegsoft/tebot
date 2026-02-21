grammar Tebot;

TEBOT_WHITESPACE:     [ \t\r\n]+ -> channel(HIDDEN);
TEBOT_COMMENT:        '#' ~[\r\n]* -> channel(HIDDEN);

TEBOT_QUOTED_STRING:  '"' (~[\r\n"])* '"';

tebot_selector
  :   TEBOT_QUOTED_STRING
  ;

tebot_value
  :   TEBOT_QUOTED_STRING
  ;

tebot_action
  :   'click'
  |   'input'
  |   'select'
  |   'capture'
  |   'assert'
  |   'sleep'
  |   'goto'
  |   'move'
  |   'scroll'
  |   'save'
  |   'paste'
  ;

tebot_assign
  :   (selector=tebot_selector '=')? value=tebot_value
  ;

tebot_operation
  :   tebot_action '(' (tebot_selector | tebot_assign) ')'
  ;

tebot_operations
  :   tebot_operation*
  ;

