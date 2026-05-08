// Generated from /Users/christian/export/local/works/doublegsoft.net/tebot/03.Development/tebot/gram/Tebot.g4 by ANTLR 4.9.2
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link TebotParser}.
 */
public interface TebotListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link TebotParser#tebot_selector}.
	 * @param ctx the parse tree
	 */
	void enterTebot_selector(TebotParser.Tebot_selectorContext ctx);
	/**
	 * Exit a parse tree produced by {@link TebotParser#tebot_selector}.
	 * @param ctx the parse tree
	 */
	void exitTebot_selector(TebotParser.Tebot_selectorContext ctx);
	/**
	 * Enter a parse tree produced by {@link TebotParser#tebot_value}.
	 * @param ctx the parse tree
	 */
	void enterTebot_value(TebotParser.Tebot_valueContext ctx);
	/**
	 * Exit a parse tree produced by {@link TebotParser#tebot_value}.
	 * @param ctx the parse tree
	 */
	void exitTebot_value(TebotParser.Tebot_valueContext ctx);
	/**
	 * Enter a parse tree produced by {@link TebotParser#tebot_action}.
	 * @param ctx the parse tree
	 */
	void enterTebot_action(TebotParser.Tebot_actionContext ctx);
	/**
	 * Exit a parse tree produced by {@link TebotParser#tebot_action}.
	 * @param ctx the parse tree
	 */
	void exitTebot_action(TebotParser.Tebot_actionContext ctx);
	/**
	 * Enter a parse tree produced by {@link TebotParser#tebot_assign}.
	 * @param ctx the parse tree
	 */
	void enterTebot_assign(TebotParser.Tebot_assignContext ctx);
	/**
	 * Exit a parse tree produced by {@link TebotParser#tebot_assign}.
	 * @param ctx the parse tree
	 */
	void exitTebot_assign(TebotParser.Tebot_assignContext ctx);
	/**
	 * Enter a parse tree produced by {@link TebotParser#tebot_operation}.
	 * @param ctx the parse tree
	 */
	void enterTebot_operation(TebotParser.Tebot_operationContext ctx);
	/**
	 * Exit a parse tree produced by {@link TebotParser#tebot_operation}.
	 * @param ctx the parse tree
	 */
	void exitTebot_operation(TebotParser.Tebot_operationContext ctx);
	/**
	 * Enter a parse tree produced by {@link TebotParser#tebot_operations}.
	 * @param ctx the parse tree
	 */
	void enterTebot_operations(TebotParser.Tebot_operationsContext ctx);
	/**
	 * Exit a parse tree produced by {@link TebotParser#tebot_operations}.
	 * @param ctx the parse tree
	 */
	void exitTebot_operations(TebotParser.Tebot_operationsContext ctx);
}