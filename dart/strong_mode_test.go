/*
Copyright © 2021 Glenn M. Lewis

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package dart

import (
	_ "embed"
	"testing"
)

//go:embed testfiles/strong_mode.dart.txt
var strong_mode_dart_txt string

//go:embed testfiles/strong_mode_want.txt
var strong_mode_want_txt string

func TestStrongMode_GetClasses(t *testing.T) {
	e, err := NewEditor(strong_mode_dart_txt, false)
	if err != nil {
		t.Fatalf("NewEditor: %v", err)
	}

	classes, err := e.GetClasses(false)
	if err != nil {
		t.Fatalf("GetClasses: %v", err)
	}

	if got, want := len(classes), 3; got != want {
		t.Errorf("got %v classes, want %v", len(classes), want)
	}
}

func TestStrongMode_Class1(t *testing.T) {
	source := strong_mode_dart_txt[0:91956]
	// wantSource := strong_mode_want_txt[560:769]

	opts := &Options{
		MemberOrdering: defaultMemberOrdering,
	}

	want := []EntityType{
		Unknown,                 // line #29: {
		PrivateInstanceVariable, // line #30:   TypeAssertions _assertions;
		BlankLine,               // line #31:
		PrivateInstanceVariable, // line #32:   Asserter<DartType> _isDynamic;
		PrivateInstanceVariable, // line #33:   Asserter<InterfaceType> _isFutureOfDynamic;
		PrivateInstanceVariable, // line #34:   Asserter<InterfaceType> _isFutureOfInt;
		PrivateInstanceVariable, // line #35:   Asserter<InterfaceType> _isFutureOfNull;
		PrivateInstanceVariable, // line #36:   Asserter<InterfaceType> _isFutureOrOfInt;
		PrivateInstanceVariable, // line #37:   Asserter<DartType> _isInt;
		PrivateInstanceVariable, // line #38:   Asserter<DartType> _isNull;
		PrivateInstanceVariable, // line #39:   Asserter<DartType> _isNum;
		PrivateInstanceVariable, // line #40:   Asserter<DartType> _isObject;
		PrivateInstanceVariable, // line #41:   Asserter<DartType> _isString;
		BlankLine,               // line #42:
		PrivateInstanceVariable, // line #43:   AsserterBuilder2<Asserter<DartType>, Asserter<DartType>, DartType>
		PrivateInstanceVariable, // line #44:       _isFunction2Of;
		PrivateInstanceVariable, // line #45:   AsserterBuilder<List<Asserter<DartType>>, InterfaceType> _isFutureOf;
		PrivateInstanceVariable, // line #46:   AsserterBuilder<List<Asserter<DartType>>, InterfaceType> _isFutureOrOf;
		PrivateInstanceVariable, // line #47:   AsserterBuilderBuilder<Asserter<DartType>, List<Asserter<DartType>>, DartType>
		PrivateInstanceVariable, // line #48:       _isInstantiationOf;
		PrivateInstanceVariable, // line #49:   AsserterBuilder<Asserter<DartType>, InterfaceType> _isListOf;
		PrivateInstanceVariable, // line #50:   AsserterBuilder2<Asserter<DartType>, Asserter<DartType>, InterfaceType>
		PrivateInstanceVariable, // line #51:       _isMapOf;
		PrivateInstanceVariable, // line #52:   AsserterBuilder<DartType, DartType> _isType;
		BlankLine,               // line #53:
		PrivateInstanceVariable, // line #54:   AsserterBuilder<Element, DartType> _hasElement;
		BlankLine,               // line #55:
		OtherMethod,             // line #56:   CompilationUnit get unit => result.unit;
		BlankLine,               // line #57:
		OverrideMethod,          // line #58:   @override
		OverrideMethod,          // line #59:   Future<void> resolveTestFile() async {
		OverrideMethod,          // line #60:     var result = await super.resolveTestFile();
		OverrideMethod,          // line #61:
		OverrideMethod,          // line #62:     if (_assertions == null) {
		OverrideMethod,          // line #63:       _assertions = TypeAssertions(typeProvider);
		OverrideMethod,          // line #64:       _isType = _assertions.isType;
		OverrideMethod,          // line #65:       _hasElement = _assertions.hasElement;
		OverrideMethod,          // line #66:       _isInstantiationOf = _assertions.isInstantiationOf;
		OverrideMethod,          // line #67:       _isInt = _assertions.isInt;
		OverrideMethod,          // line #68:       _isNull = _assertions.isNull;
		OverrideMethod,          // line #69:       _isNum = _assertions.isNum;
		OverrideMethod,          // line #70:       _isObject = _assertions.isObject;
		OverrideMethod,          // line #71:       _isString = _assertions.isString;
		OverrideMethod,          // line #72:       _isDynamic = _assertions.isDynamic;
		OverrideMethod,          // line #73:       _isListOf = _assertions.isListOf;
		OverrideMethod,          // line #74:       _isMapOf = _assertions.isMapOf;
		OverrideMethod,          // line #75:       _isFunction2Of = _assertions.isFunction2Of;
		OverrideMethod,          // line #76:       _isFutureOf = _isInstantiationOf(_hasElement(typeProvider.futureElement));
		OverrideMethod,          // line #77:       _isFutureOrOf =
		OverrideMethod,          // line #78:           _isInstantiationOf(_hasElement(typeProvider.futureOrElement));
		OverrideMethod,          // line #79:       _isFutureOfDynamic = _isFutureOf([_isDynamic]);
		OverrideMethod,          // line #80:       _isFutureOfInt = _isFutureOf([_isInt]);
		OverrideMethod,          // line #81:       _isFutureOfNull = _isFutureOf([_isNull]);
		OverrideMethod,          // line #82:       _isFutureOrOfInt = _isFutureOrOf([_isInt]);
		OverrideMethod,          // line #83:     }
		OverrideMethod,          // line #84:
		OverrideMethod,          // line #85:     return result;
		OverrideMethod,          // line #86:   }
		BlankLine,               // line #87:
		OtherMethod,             // line #88:   test_async_method_propagation() async {
		OtherMethod,             // line #89:     String code = r'''
		OtherMethod,             // line #90:       class A {
		OtherMethod,             // line #91:         Future f0() => new Future.value(3);
		OtherMethod,             // line #92:         Future f1() async => new Future.value(3);
		OtherMethod,             // line #93:         Future f2() async => await new Future.value(3);
		OtherMethod,             // line #94:
		OtherMethod,             // line #95:         Future<int> f3() => new Future.value(3);
		OtherMethod,             // line #96:         Future<int> f4() async => new Future.value(3);
		OtherMethod,             // line #97:         Future<int> f5() async => await new Future.value(3);
		OtherMethod,             // line #98:
		OtherMethod,             // line #99:         Future g0() { return new Future.value(3); }
		OtherMethod,             // line #100:         Future g1() async { return new Future.value(3); }
		OtherMethod,             // line #101:         Future g2() async { return await new Future.value(3); }
		OtherMethod,             // line #102:
		OtherMethod,             // line #103:         Future<int> g3() { return new Future.value(3); }
		OtherMethod,             // line #104:         Future<int> g4() async { return new Future.value(3); }
		OtherMethod,             // line #105:         Future<int> g5() async { return await new Future.value(3); }
		OtherMethod,             // line #106:       }
		OtherMethod,             // line #107:    ''';
		OtherMethod,             // line #108:     await resolveTestCode(code);
		OtherMethod,             // line #109:
		OtherMethod,             // line #110:     void check(String name, Asserter<InterfaceType> typeTest) {
		OtherMethod,             // line #111:       MethodDeclaration test = AstFinder.getMethodInClass(unit, "A", name);
		OtherMethod,             // line #112:       FunctionBody body = test.body;
		OtherMethod,             // line #113:       Expression returnExp;
		OtherMethod,             // line #114:       if (body is ExpressionFunctionBody) {
		OtherMethod,             // line #115:         returnExp = body.expression;
		OtherMethod,             // line #116:       } else {
		OtherMethod,             // line #117:         ReturnStatement stmt = (body as BlockFunctionBody).block.statements[0];
		OtherMethod,             // line #118:         returnExp = stmt.expression;
		OtherMethod,             // line #119:       }
		OtherMethod,             // line #120:       DartType type = returnExp.staticType;
		OtherMethod,             // line #121:       if (returnExp is AwaitExpression) {
		OtherMethod,             // line #122:         type = returnExp.expression.staticType;
		OtherMethod,             // line #123:       }
		OtherMethod,             // line #124:       typeTest(type);
		OtherMethod,             // line #125:     }
		OtherMethod,             // line #126:
		OtherMethod,             // line #127:     check("f0", _isFutureOfDynamic);
		OtherMethod,             // line #128:     check("f1", _isFutureOfDynamic);
		OtherMethod,             // line #129:     check("f2", _isFutureOfDynamic);
		OtherMethod,             // line #130:
		OtherMethod,             // line #131:     check("f3", _isFutureOfInt);
		OtherMethod,             // line #132:     check("f4", _isFutureOfInt);
		OtherMethod,             // line #133:     check("f5", _isFutureOfInt);
		OtherMethod,             // line #134:
		OtherMethod,             // line #135:     check("g0", _isFutureOfDynamic);
		OtherMethod,             // line #136:     check("g1", _isFutureOfDynamic);
		OtherMethod,             // line #137:     check("g2", _isFutureOfDynamic);
		OtherMethod,             // line #138:
		OtherMethod,             // line #139:     check("g3", _isFutureOfInt);
		OtherMethod,             // line #140:     check("g4", _isFutureOfInt);
		OtherMethod,             // line #141:     check("g5", _isFutureOfInt);
		OtherMethod,             // line #142:   }
		BlankLine,               // line #143:
		OtherMethod,             // line #144:   test_async_propagation() async {
		OtherMethod,             // line #145:     String code = r'''
		OtherMethod,             // line #146:       Future f0() => new Future.value(3);
		OtherMethod,             // line #147:       Future f1() async => new Future.value(3);
		OtherMethod,             // line #148:       Future f2() async => await new Future.value(3);
		OtherMethod,             // line #149:
		OtherMethod,             // line #150:       Future<int> f3() => new Future.value(3);
		OtherMethod,             // line #151:       Future<int> f4() async => new Future.value(3);
		OtherMethod,             // line #152:       Future<int> f5() async => await new Future.value(3);
		OtherMethod,             // line #153:
		OtherMethod,             // line #154:       Future g0() { return new Future.value(3); }
		OtherMethod,             // line #155:       Future g1() async { return new Future.value(3); }
		OtherMethod,             // line #156:       Future g2() async { return await new Future.value(3); }
		OtherMethod,             // line #157:
		OtherMethod,             // line #158:       Future<int> g3() { return new Future.value(3); }
		OtherMethod,             // line #159:       Future<int> g4() async { return new Future.value(3); }
		OtherMethod,             // line #160:       Future<int> g5() async { return await new Future.value(3); }
		OtherMethod,             // line #161:    ''';
		OtherMethod,             // line #162:     await resolveTestCode(code);
		OtherMethod,             // line #163:
		OtherMethod,             // line #164:     void check(String name, Asserter<InterfaceType> typeTest) {
		OtherMethod,             // line #165:       FunctionDeclaration test = AstFinder.getTopLevelFunction(unit, name);
		OtherMethod,             // line #166:       FunctionBody body = test.functionExpression.body;
		OtherMethod,             // line #167:       Expression returnExp;
		OtherMethod,             // line #168:       if (body is ExpressionFunctionBody) {
		OtherMethod,             // line #169:         returnExp = body.expression;
		OtherMethod,             // line #170:       } else {
		OtherMethod,             // line #171:         ReturnStatement stmt = (body as BlockFunctionBody).block.statements[0];
		OtherMethod,             // line #172:         returnExp = stmt.expression;
		OtherMethod,             // line #173:       }
		OtherMethod,             // line #174:       DartType type = returnExp.staticType;
		OtherMethod,             // line #175:       if (returnExp is AwaitExpression) {
		OtherMethod,             // line #176:         type = returnExp.expression.staticType;
		OtherMethod,             // line #177:       }
		OtherMethod,             // line #178:       typeTest(type);
		OtherMethod,             // line #179:     }
		OtherMethod,             // line #180:
		OtherMethod,             // line #181:     check("f0", _isFutureOfDynamic);
		OtherMethod,             // line #182:     check("f1", _isFutureOfDynamic);
		OtherMethod,             // line #183:     check("f2", _isFutureOfDynamic);
		OtherMethod,             // line #184:
		OtherMethod,             // line #185:     check("f3", _isFutureOfInt);
		OtherMethod,             // line #186:     check("f4", _isFutureOfInt);
		OtherMethod,             // line #187:     check("f5", _isFutureOfInt);
		OtherMethod,             // line #188:
		OtherMethod,             // line #189:     check("g0", _isFutureOfDynamic);
		OtherMethod,             // line #190:     check("g1", _isFutureOfDynamic);
		OtherMethod,             // line #191:     check("g2", _isFutureOfDynamic);
		OtherMethod,             // line #192:
		OtherMethod,             // line #193:     check("g3", _isFutureOfInt);
		OtherMethod,             // line #194:     check("g4", _isFutureOfInt);
		OtherMethod,             // line #195:     check("g5", _isFutureOfInt);
		OtherMethod,             // line #196:   }
		BlankLine,               // line #197:
		OtherMethod,             // line #198:   test_cascadeExpression() async {
		OtherMethod,             // line #199:     String code = r'''
		OtherMethod,             // line #200:       class A<T> {
		OtherMethod,             // line #201:         List<T> map(T a, List<T> mapper(T x)) => mapper(a);
		OtherMethod,             // line #202:       }
		OtherMethod,             // line #203:
		OtherMethod,             // line #204:       void main () {
		OtherMethod,             // line #205:         A<int> a = new A()..map(0, (x) => [x]);
		OtherMethod,             // line #206:      }
		OtherMethod,             // line #207:    ''';
		OtherMethod,             // line #208:     await resolveTestCode(code);
		OtherMethod,             // line #209:     List<Statement> statements =
		OtherMethod,             // line #210:         AstFinder.getStatementsInTopLevelFunction(unit, "main");
		OtherMethod,             // line #211:     CascadeExpression fetch(int i) {
		OtherMethod,             // line #212:       VariableDeclarationStatement stmt = statements[i];
		OtherMethod,             // line #213:       VariableDeclaration decl = stmt.variables.variables[0];
		OtherMethod,             // line #214:       CascadeExpression exp = decl.initializer;
		OtherMethod,             // line #215:       return exp;
		OtherMethod,             // line #216:     }
		OtherMethod,             // line #217:
		OtherMethod,             // line #218:     Element elementA = AstFinder.getClass(unit, "A").declaredElement;
		OtherMethod,             // line #219:
		OtherMethod,             // line #220:     CascadeExpression cascade = fetch(0);
		OtherMethod,             // line #221:     _isInstantiationOf(_hasElement(elementA))([_isInt])(cascade.staticType);
		OtherMethod,             // line #222:     MethodInvocation invoke = cascade.cascadeSections[0];
		OtherMethod,             // line #223:     FunctionExpression function = invoke.argumentList.arguments[1];
		OtherMethod,             // line #224:     ExecutableElement f0 = function.declaredElement;
		OtherMethod,             // line #225:     _isListOf(_isInt)(f0.type.returnType);
		OtherMethod,             // line #226:     expect(f0.type.normalParameterTypes[0], typeProvider.intType);
		OtherMethod,             // line #227:   }
		BlankLine,               // line #228:
		OtherMethod,             // line #229:   test_constrainedByBounds1() async {
		OtherMethod,             // line #230:     // Test that upwards inference with two type variables correctly
		OtherMethod,             // line #231:     // propogates from the constrained variable to the unconstrained
		OtherMethod,             // line #232:     // variable if they are ordered left to right.
		OtherMethod,             // line #233:     String code = r'''
		OtherMethod,             // line #234:     T f<S, T extends S>(S x) => null;
		OtherMethod,             // line #235:     void test() { var x = f(3); }
		OtherMethod,             // line #236:    ''';
		OtherMethod,             // line #237:     await assertErrorsInCode(code, [
		OtherMethod,             // line #238:       error(HintCode.UNUSED_LOCAL_VARIABLE, 60, 1),
		OtherMethod,             // line #239:     ]);
		OtherMethod,             // line #240:
		OtherMethod,             // line #241:     List<Statement> statements =
		OtherMethod,             // line #242:         AstFinder.getStatementsInTopLevelFunction(unit, "test");
		OtherMethod,             // line #243:     VariableDeclarationStatement stmt = statements[0];
		OtherMethod,             // line #244:     VariableDeclaration decl = stmt.variables.variables[0];
		OtherMethod,             // line #245:     Expression call = decl.initializer;
		OtherMethod,             // line #246:     _isInt(call.staticType);
		OtherMethod,             // line #247:   }
		BlankLine,               // line #248:
		OtherMethod,             // line #249:   test_constrainedByBounds2() async {
		OtherMethod,             // line #250:     // Test that upwards inference with two type variables does
		OtherMethod,             // line #251:     // propogate from the constrained variable to the unconstrained
		OtherMethod,             // line #252:     // variable if they are ordered right to left.
		OtherMethod,             // line #253:     String code = r'''
		OtherMethod,             // line #254:     T f<T extends S, S>(S x) => null;
		OtherMethod,             // line #255:     void test() { var x = f(3); }
		OtherMethod,             // line #256:    ''';
		OtherMethod,             // line #257:     await assertErrorsInCode(code, [
		OtherMethod,             // line #258:       error(HintCode.UNUSED_LOCAL_VARIABLE, 60, 1),
		OtherMethod,             // line #259:     ]);
		OtherMethod,             // line #260:
		OtherMethod,             // line #261:     List<Statement> statements =
		OtherMethod,             // line #262:         AstFinder.getStatementsInTopLevelFunction(unit, "test");
		OtherMethod,             // line #263:     VariableDeclarationStatement stmt = statements[0];
		OtherMethod,             // line #264:     VariableDeclaration decl = stmt.variables.variables[0];
		OtherMethod,             // line #265:     Expression call = decl.initializer;
		OtherMethod,             // line #266:     _isInt(call.staticType);
		OtherMethod,             // line #267:   }
		BlankLine,               // line #268:
		OtherMethod,             // line #269:   test_constrainedByBounds3() async {
		OtherMethod,             // line #270:     var code = r'''
		OtherMethod,             // line #271:       T f<T extends S, S extends int>(S x) => null;
		OtherMethod,             // line #272:       void test() { var x = f(3); }
		OtherMethod,             // line #273:    ''';
		OtherMethod,             // line #274:     await assertErrorsInCode(code, [
		OtherMethod,             // line #275:       error(HintCode.UNUSED_LOCAL_VARIABLE, 76, 1),
		OtherMethod,             // line #276:     ]);
		OtherMethod,             // line #277:
		OtherMethod,             // line #278:     List<Statement> statements =
		OtherMethod,             // line #279:         AstFinder.getStatementsInTopLevelFunction(unit, "test");
		OtherMethod,             // line #280:     VariableDeclarationStatement stmt = statements[0];
		OtherMethod,             // line #281:     VariableDeclaration decl = stmt.variables.variables[0];
		OtherMethod,             // line #282:     Expression call = decl.initializer;
		OtherMethod,             // line #283:     _isInt(call.staticType);
		OtherMethod,             // line #284:   }
		BlankLine,               // line #285:
		OtherMethod,             // line #286:   test_constrainedByBounds4() async {
		OtherMethod,             // line #287:     // Test that upwards inference with two type variables correctly
		OtherMethod,             // line #288:     // propogates from the constrained variable to the unconstrained
		OtherMethod,             // line #289:     // variable if they are ordered left to right, when the variable
		OtherMethod,             // line #290:     // appears co and contra variantly
		OtherMethod,             // line #291:     String code = r'''
		OtherMethod,             // line #292:     typedef To Func1<From, To>(From x);
		OtherMethod,             // line #293:     T f<S, T extends Func1<S, S>>(S x) => null;
		OtherMethod,             // line #294:     void test() { var x = f(3)(4); }
		OtherMethod,             // line #295:    ''';
		OtherMethod,             // line #296:     await assertErrorsInCode(code, [
		OtherMethod,             // line #297:       error(HintCode.UNUSED_LOCAL_VARIABLE, 110, 1),
		OtherMethod,             // line #298:     ]);
		OtherMethod,             // line #299:
		OtherMethod,             // line #300:     List<Statement> statements =
		OtherMethod,             // line #301:         AstFinder.getStatementsInTopLevelFunction(unit, "test");
		OtherMethod,             // line #302:     VariableDeclarationStatement stmt = statements[0];
		OtherMethod,             // line #303:     VariableDeclaration decl = stmt.variables.variables[0];
		OtherMethod,             // line #304:     Expression call = decl.initializer;
		OtherMethod,             // line #305:     _isInt(call.staticType);
		OtherMethod,             // line #306:   }
		BlankLine,               // line #307:
		OtherMethod,             // line #308:   test_constrainedByBounds5() async {
		OtherMethod,             // line #309:     // Test that upwards inference with two type variables does not
		OtherMethod,             // line #310:     // propagate from the constrained variable to the unconstrained
		OtherMethod,             // line #311:     // variable if they are ordered right to left, when the variable
		OtherMethod,             // line #312:     // appears co- and contra-variantly, and that an error is issued
		OtherMethod,             // line #313:     // for the non-matching bound.
		OtherMethod,             // line #314:     String code = r'''
		OtherMethod,             // line #315:     typedef To Func1<From, To>(From x);
		OtherMethod,             // line #316:     T f<T extends Func1<S, S>, S>(S x) => null;
		OtherMethod,             // line #317:     void test() { var x = f(3)(null); }
		OtherMethod,             // line #318:    ''';
		OtherMethod,             // line #319:     await assertErrorsInCode(code, [
		OtherMethod,             // line #320:       error(HintCode.UNUSED_LOCAL_VARIABLE, 110, 1),
		OtherMethod,             // line #321:       error(CompileTimeErrorCode.COULD_NOT_INFER, 114, 1),
		OtherMethod,             // line #322:     ]);
		OtherMethod,             // line #323:
		OtherMethod,             // line #324:     List<Statement> statements =
		OtherMethod,             // line #325:         AstFinder.getStatementsInTopLevelFunction(unit, "test");
		OtherMethod,             // line #326:     VariableDeclarationStatement stmt = statements[0];
		OtherMethod,             // line #327:     VariableDeclaration decl = stmt.variables.variables[0];
		OtherMethod,             // line #328:     Expression call = decl.initializer;
		OtherMethod,             // line #329:     _isDynamic(call.staticType);
		OtherMethod,             // line #330:   }
		BlankLine,               // line #331:
		OtherMethod,             // line #332:   test_constructorInitializer_propagation() async {
		OtherMethod,             // line #333:     String code = r'''
		OtherMethod,             // line #334:       class A {
		OtherMethod,             // line #335:         List<String> x;
		OtherMethod,             // line #336:         A() : this.x = [];
		OtherMethod,             // line #337:       }
		OtherMethod,             // line #338:    ''';
		OtherMethod,             // line #339:     await assertNoErrorsInCode(code);
		OtherMethod,             // line #340:     ConstructorDeclaration constructor =
		OtherMethod,             // line #341:         AstFinder.getConstructorInClass(unit, "A", null);
		OtherMethod,             // line #342:     ConstructorFieldInitializer assignment = constructor.initializers[0];
		OtherMethod,             // line #343:     Expression exp = assignment.expression;
		OtherMethod,             // line #344:     _isListOf(_isString)(exp.staticType);
		OtherMethod,             // line #345:   }
		BlankLine,               // line #346:
		OtherMethod,             // line #347:   test_factoryConstructor_propagation() async {
		OtherMethod,             // line #348:     String code = r'''
		OtherMethod,             // line #349:       class A<T> {
		OtherMethod,             // line #350:         factory A() { return new B(); }
		OtherMethod,             // line #351:       }
		OtherMethod,             // line #352:       class B<S> extends A<S> {}
		OtherMethod,             // line #353:    ''';
		OtherMethod,             // line #354:     await assertErrorsInCode(code, [
		OtherMethod,             // line #355:       error(
		OtherMethod,             // line #356:           CompileTimeErrorCode.NO_GENERATIVE_CONSTRUCTORS_IN_SUPERCLASS, 92, 4),
		OtherMethod,             // line #357:     ]);
		OtherMethod,             // line #358:
		OtherMethod,             // line #359:     ConstructorDeclaration constructor =
		OtherMethod,             // line #360:         AstFinder.getConstructorInClass(unit, "A", null);
		OtherMethod,             // line #361:     BlockFunctionBody body = constructor.body;
		OtherMethod,             // line #362:     ReturnStatement stmt = body.block.statements[0];
		OtherMethod,             // line #363:     InstanceCreationExpression exp = stmt.expression;
		OtherMethod,             // line #364:     ClassElement elementB = AstFinder.getClass(unit, "B").declaredElement;
		OtherMethod,             // line #365:     ClassElement elementA = AstFinder.getClass(unit, "A").declaredElement;
		OtherMethod,             // line #366:     expect(exp.constructorName.type.type.element, elementB);
		OtherMethod,             // line #367:     _isInstantiationOf(_hasElement(elementB))([
		OtherMethod,             // line #368:       _isType(elementA.typeParameters[0]
		OtherMethod,             // line #369:           .instantiate(nullabilitySuffix: NullabilitySuffix.star))
		OtherMethod,             // line #370:     ])(exp.staticType);
		OtherMethod,             // line #371:   }
		BlankLine,               // line #372:
		OtherMethod,             // line #373:   test_fieldDeclaration_propagation() async {
		OtherMethod,             // line #374:     String code = r'''
		OtherMethod,             // line #375:       class A {
		OtherMethod,             // line #376:         List<String> f0 = ["hello"];
		OtherMethod,             // line #377:       }
		OtherMethod,             // line #378:    ''';
		OtherMethod,             // line #379:     await assertNoErrorsInCode(code);
		OtherMethod,             // line #380:
		OtherMethod,             // line #381:     VariableDeclaration field = AstFinder.getFieldInClass(unit, "A", "f0");
		OtherMethod,             // line #382:
		OtherMethod,             // line #383:     _isListOf(_isString)(field.initializer.staticType);
		OtherMethod,             // line #384:   }
		BlankLine,               // line #385:
		OtherMethod,             // line #386:   test_functionDeclaration_body_propagation() async {
		OtherMethod,             // line #387:     String code = r'''
		OtherMethod,             // line #388:       typedef T Function2<S, T>(S x);
		OtherMethod,             // line #389:
		OtherMethod,             // line #390:       List<int> test1() => [];
		OtherMethod,             // line #391:
		OtherMethod,             // line #392:       Function2<int, int> test2 (int x) {
		OtherMethod,             // line #393:         Function2<String, int> inner() {
		OtherMethod,             // line #394:           return (x) => x.length;
		OtherMethod,             // line #395:         }
		OtherMethod,             // line #396:         return (x) => x;
		OtherMethod,             // line #397:      }
		OtherMethod,             // line #398:    ''';
		OtherMethod,             // line #399:     await assertErrorsInCode(code, [
		OtherMethod,             // line #400:       error(HintCode.UNUSED_ELEMENT, 144, 5),
		OtherMethod,             // line #401:     ]);
		OtherMethod,             // line #402:
		OtherMethod,             // line #403:     Asserter<InterfaceType> assertListOfInt = _isListOf(_isInt);
		OtherMethod,             // line #404:
		OtherMethod,             // line #405:     FunctionDeclaration test1 = AstFinder.getTopLevelFunction(unit, "test1");
		OtherMethod,             // line #406:     ExpressionFunctionBody body = test1.functionExpression.body;
		OtherMethod,             // line #407:     assertListOfInt(body.expression.staticType);
		OtherMethod,             // line #408:
		OtherMethod,             // line #409:     List<Statement> statements =
		OtherMethod,             // line #410:         AstFinder.getStatementsInTopLevelFunction(unit, "test2");
		OtherMethod,             // line #411:
		OtherMethod,             // line #412:     FunctionDeclaration inner =
		OtherMethod,             // line #413:         (statements[0] as FunctionDeclarationStatement).functionDeclaration;
		OtherMethod,             // line #414:     BlockFunctionBody body0 = inner.functionExpression.body;
		OtherMethod,             // line #415:     ReturnStatement return0 = body0.block.statements[0];
		OtherMethod,             // line #416:     Expression anon0 = return0.expression;
		OtherMethod,             // line #417:     FunctionType type0 = anon0.staticType;
		OtherMethod,             // line #418:     expect(type0.returnType, typeProvider.intType);
		OtherMethod,             // line #419:     expect(type0.normalParameterTypes[0], typeProvider.stringType);
		OtherMethod,             // line #420:
		OtherMethod,             // line #421:     FunctionExpression anon1 = (statements[1] as ReturnStatement).expression;
		OtherMethod,             // line #422:     FunctionType type1 = anon1.declaredElement.type;
		OtherMethod,             // line #423:     expect(type1.returnType, typeProvider.intType);
		OtherMethod,             // line #424:     expect(type1.normalParameterTypes[0], typeProvider.intType);
		OtherMethod,             // line #425:   }
		BlankLine,               // line #426:
		OtherMethod,             // line #427:   test_functionLiteral_assignment_typedArguments() async {
		OtherMethod,             // line #428:     String code = r'''
		OtherMethod,             // line #429:       typedef T Function2<S, T>(S x);
		OtherMethod,             // line #430:
		OtherMethod,             // line #431:       void main () {
		OtherMethod,             // line #432:         Function2<int, String> l0 = (int x) => null;
		OtherMethod,             // line #433:         Function2<int, String> l1 = (int x) => "hello";
		OtherMethod,             // line #434:         Function2<int, String> l2 = (String x) => "hello";
		OtherMethod,             // line #435:         Function2<int, String> l3 = (int x) => 3;
		OtherMethod,             // line #436:         Function2<int, String> l4 = (int x) {return 3;};
		OtherMethod,             // line #437:      }
		OtherMethod,             // line #438:    ''';
		OtherMethod,             // line #439:     await assertErrorsInCode(code, [
		OtherMethod,             // line #440:       error(HintCode.UNUSED_LOCAL_VARIABLE, 91, 2),
		OtherMethod,             // line #441:       error(HintCode.UNUSED_LOCAL_VARIABLE, 144, 2),
		OtherMethod,             // line #442:       error(HintCode.UNUSED_LOCAL_VARIABLE, 200, 2),
		OtherMethod,             // line #443:       error(CompileTimeErrorCode.INVALID_ASSIGNMENT, 205, 21),
		OtherMethod,             // line #444:       error(HintCode.UNUSED_LOCAL_VARIABLE, 259, 2),
		OtherMethod,             // line #445:       error(CompileTimeErrorCode.RETURN_OF_INVALID_TYPE_FROM_CLOSURE, 275, 1),
		OtherMethod,             // line #446:       error(HintCode.UNUSED_LOCAL_VARIABLE, 309, 2),
		OtherMethod,             // line #447:       error(CompileTimeErrorCode.RETURN_OF_INVALID_TYPE_FROM_CLOSURE, 330, 1),
		OtherMethod,             // line #448:     ]);
		OtherMethod,             // line #449:
		OtherMethod,             // line #450:     List<Statement> statements =
		OtherMethod,             // line #451:         AstFinder.getStatementsInTopLevelFunction(unit, "main");
		OtherMethod,             // line #452:     DartType literal(int i) {
		OtherMethod,             // line #453:       VariableDeclarationStatement stmt = statements[i];
		OtherMethod,             // line #454:       VariableDeclaration decl = stmt.variables.variables[0];
		OtherMethod,             // line #455:       FunctionExpression exp = decl.initializer;
		OtherMethod,             // line #456:       return exp.declaredElement.type;
		OtherMethod,             // line #457:     }
		OtherMethod,             // line #458:
		OtherMethod,             // line #459:     _isFunction2Of(_isInt, _isNull)(literal(0));
		OtherMethod,             // line #460:     _isFunction2Of(_isInt, _isString)(literal(1));
		OtherMethod,             // line #461:     _isFunction2Of(_isString, _isString)(literal(2));
		OtherMethod,             // line #462:     _isFunction2Of(_isInt, _isString)(literal(3));
		OtherMethod,             // line #463:     _isFunction2Of(_isInt, _isString)(literal(4));
		OtherMethod,             // line #464:   }
		BlankLine,               // line #465:
		OtherMethod,             // line #466:   test_functionLiteral_assignment_unTypedArguments() async {
		OtherMethod,             // line #467:     String code = r'''
		OtherMethod,             // line #468:       typedef T Function2<S, T>(S x);
		OtherMethod,             // line #469:
		OtherMethod,             // line #470:       void main () {
		OtherMethod,             // line #471:         Function2<int, String> l0 = (x) => null;
		OtherMethod,             // line #472:         Function2<int, String> l1 = (x) => "hello";
		OtherMethod,             // line #473:         Function2<int, String> l2 = (x) => "hello";
		OtherMethod,             // line #474:         Function2<int, String> l3 = (x) => 3;
		OtherMethod,             // line #475:         Function2<int, String> l4 = (x) {return 3;};
		OtherMethod,             // line #476:      }
		OtherMethod,             // line #477:    ''';
		OtherMethod,             // line #478:     await assertErrorsInCode(code, [
		OtherMethod,             // line #479:       error(HintCode.UNUSED_LOCAL_VARIABLE, 91, 2),
		OtherMethod,             // line #480:       error(HintCode.UNUSED_LOCAL_VARIABLE, 140, 2),
		OtherMethod,             // line #481:       error(HintCode.UNUSED_LOCAL_VARIABLE, 192, 2),
		OtherMethod,             // line #482:       error(HintCode.UNUSED_LOCAL_VARIABLE, 244, 2),
		OtherMethod,             // line #483:       error(CompileTimeErrorCode.RETURN_OF_INVALID_TYPE_FROM_CLOSURE, 256, 1),
		OtherMethod,             // line #484:       error(HintCode.UNUSED_LOCAL_VARIABLE, 290, 2),
		OtherMethod,             // line #485:       error(CompileTimeErrorCode.RETURN_OF_INVALID_TYPE_FROM_CLOSURE, 307, 1),
		OtherMethod,             // line #486:     ]);
		OtherMethod,             // line #487:
		OtherMethod,             // line #488:     List<Statement> statements =
		OtherMethod,             // line #489:         AstFinder.getStatementsInTopLevelFunction(unit, "main");
		OtherMethod,             // line #490:     DartType literal(int i) {
		OtherMethod,             // line #491:       VariableDeclarationStatement stmt = statements[i];
		OtherMethod,             // line #492:       VariableDeclaration decl = stmt.variables.variables[0];
		OtherMethod,             // line #493:       FunctionExpression exp = decl.initializer;
		OtherMethod,             // line #494:       return exp.declaredElement.type;
		OtherMethod,             // line #495:     }
		OtherMethod,             // line #496:
		OtherMethod,             // line #497:     _isFunction2Of(_isInt, _isNull)(literal(0));
		OtherMethod,             // line #498:     _isFunction2Of(_isInt, _isString)(literal(1));
		OtherMethod,             // line #499:     _isFunction2Of(_isInt, _isString)(literal(2));
		OtherMethod,             // line #500:     _isFunction2Of(_isInt, _isString)(literal(3));
		OtherMethod,             // line #501:     _isFunction2Of(_isInt, _isString)(literal(4));
		OtherMethod,             // line #502:   }
		BlankLine,               // line #503:
		OtherMethod,             // line #504:   test_functionLiteral_body_propagation() async {
		OtherMethod,             // line #505:     String code = r'''
		OtherMethod,             // line #506:       typedef T Function2<S, T>(S x);
		OtherMethod,             // line #507:
		OtherMethod,             // line #508:       void main () {
		OtherMethod,             // line #509:         Function2<int, List<String>> l0 = (int x) => ["hello"];
		OtherMethod,             // line #510:         Function2<int, List<String>> l1 = (String x) => ["hello"];
		OtherMethod,             // line #511:         Function2<int, List<String>> l2 = (int x) => [3];
		OtherMethod,             // line #512:         Function2<int, List<String>> l3 = (int x) {return [3];};
		OtherMethod,             // line #513:      }
		OtherMethod,             // line #514:    ''';
		OtherMethod,             // line #515:     await assertErrorsInCode(code, [
		OtherMethod,             // line #516:       error(HintCode.UNUSED_LOCAL_VARIABLE, 97, 2),
		OtherMethod,             // line #517:       error(HintCode.UNUSED_LOCAL_VARIABLE, 161, 2),
		OtherMethod,             // line #518:       error(CompileTimeErrorCode.INVALID_ASSIGNMENT, 166, 23),
		OtherMethod,             // line #519:       error(HintCode.UNUSED_LOCAL_VARIABLE, 228, 2),
		OtherMethod,             // line #520:       error(CompileTimeErrorCode.LIST_ELEMENT_TYPE_NOT_ASSIGNABLE, 245, 1),
		OtherMethod,             // line #521:       error(HintCode.UNUSED_LOCAL_VARIABLE, 286, 2),
		OtherMethod,             // line #522:       error(CompileTimeErrorCode.LIST_ELEMENT_TYPE_NOT_ASSIGNABLE, 308, 1),
		OtherMethod,             // line #523:     ]);
		OtherMethod,             // line #524:
		OtherMethod,             // line #525:     List<Statement> statements =
		OtherMethod,             // line #526:         AstFinder.getStatementsInTopLevelFunction(unit, "main");
		OtherMethod,             // line #527:     Expression functionReturnValue(int i) {
		OtherMethod,             // line #528:       VariableDeclarationStatement stmt = statements[i];
		OtherMethod,             // line #529:       VariableDeclaration decl = stmt.variables.variables[0];
		OtherMethod,             // line #530:       FunctionExpression exp = decl.initializer;
		OtherMethod,             // line #531:       FunctionBody body = exp.body;
		OtherMethod,             // line #532:       if (body is ExpressionFunctionBody) {
		OtherMethod,             // line #533:         return body.expression;
		OtherMethod,             // line #534:       } else {
		OtherMethod,             // line #535:         Statement stmt = (body as BlockFunctionBody).block.statements[0];
		OtherMethod,             // line #536:         return (stmt as ReturnStatement).expression;
		OtherMethod,             // line #537:       }
		OtherMethod,             // line #538:     }
		OtherMethod,             // line #539:
		OtherMethod,             // line #540:     Asserter<InterfaceType> assertListOfString = _isListOf(_isString);
		OtherMethod,             // line #541:     assertListOfString(functionReturnValue(0).staticType);
		OtherMethod,             // line #542:     assertListOfString(functionReturnValue(1).staticType);
		OtherMethod,             // line #543:     assertListOfString(functionReturnValue(2).staticType);
		OtherMethod,             // line #544:     assertListOfString(functionReturnValue(3).staticType);
		OtherMethod,             // line #545:   }
		BlankLine,               // line #546:
		OtherMethod,             // line #547:   test_functionLiteral_functionExpressionInvocation_typedArguments() async {
		OtherMethod,             // line #548:     String code = r'''
		OtherMethod,             // line #549:       class Mapper<F, T> {
		OtherMethod,             // line #550:         T map(T mapper(F x)) => mapper(null);
		OtherMethod,             // line #551:       }
		OtherMethod,             // line #552:
		OtherMethod,             // line #553:       void main () {
		OtherMethod,             // line #554:         (new Mapper<int, String>().map)((int x) => null);
		OtherMethod,             // line #555:         (new Mapper<int, String>().map)((int x) => "hello");
		OtherMethod,             // line #556:         (new Mapper<int, String>().map)((String x) => "hello");
		OtherMethod,             // line #557:         (new Mapper<int, String>().map)((int x) => 3);
		OtherMethod,             // line #558:         (new Mapper<int, String>().map)((int x) {return 3;});
		OtherMethod,             // line #559:      }
		OtherMethod,             // line #560:    ''';
		OtherMethod,             // line #561:     await assertErrorsInCode(code, [
		OtherMethod,             // line #562:       error(CompileTimeErrorCode.ARGUMENT_TYPE_NOT_ASSIGNABLE, 262, 21),
		OtherMethod,             // line #563:       error(CompileTimeErrorCode.RETURN_OF_INVALID_TYPE_FROM_CLOSURE, 337, 1),
		OtherMethod,             // line #564:       error(CompileTimeErrorCode.RETURN_OF_INVALID_TYPE_FROM_CLOSURE, 397, 1),
		OtherMethod,             // line #565:     ]);
		OtherMethod,             // line #566:
		OtherMethod,             // line #567:     List<Statement> statements =
		OtherMethod,             // line #568:         AstFinder.getStatementsInTopLevelFunction(unit, "main");
		OtherMethod,             // line #569:     DartType literal(int i) {
		OtherMethod,             // line #570:       ExpressionStatement stmt = statements[i];
		OtherMethod,             // line #571:       FunctionExpressionInvocation invk = stmt.expression;
		OtherMethod,             // line #572:       FunctionExpression exp = invk.argumentList.arguments[0];
		OtherMethod,             // line #573:       return exp.declaredElement.type;
		OtherMethod,             // line #574:     }
		OtherMethod,             // line #575:
		OtherMethod,             // line #576:     _isFunction2Of(_isInt, _isNull)(literal(0));
		OtherMethod,             // line #577:     _isFunction2Of(_isInt, _isString)(literal(1));
		OtherMethod,             // line #578:     _isFunction2Of(_isString, _isString)(literal(2));
		OtherMethod,             // line #579:     _isFunction2Of(_isInt, _isString)(literal(3));
		OtherMethod,             // line #580:     _isFunction2Of(_isInt, _isString)(literal(4));
		OtherMethod,             // line #581:   }
		BlankLine,               // line #582:
		OtherMethod,             // line #583:   test_functionLiteral_functionExpressionInvocation_unTypedArguments() async {
		OtherMethod,             // line #584:     String code = r'''
		OtherMethod,             // line #585:       class Mapper<F, T> {
		OtherMethod,             // line #586:         T map(T mapper(F x)) => mapper(null);
		OtherMethod,             // line #587:       }
		OtherMethod,             // line #588:
		OtherMethod,             // line #589:       void main () {
		OtherMethod,             // line #590:         (new Mapper<int, String>().map)((x) => null);
		OtherMethod,             // line #591:         (new Mapper<int, String>().map)((x) => "hello");
		OtherMethod,             // line #592:         (new Mapper<int, String>().map)((x) => "hello");
		OtherMethod,             // line #593:         (new Mapper<int, String>().map)((x) => 3);
		OtherMethod,             // line #594:         (new Mapper<int, String>().map)((x) {return 3;});
		OtherMethod,             // line #595:      }
		OtherMethod,             // line #596:    ''';
		OtherMethod,             // line #597:     await assertErrorsInCode(code, [
		OtherMethod,             // line #598:       error(CompileTimeErrorCode.RETURN_OF_INVALID_TYPE_FROM_CLOSURE, 318, 1),
		OtherMethod,             // line #599:       error(CompileTimeErrorCode.RETURN_OF_INVALID_TYPE_FROM_CLOSURE, 374, 1),
		OtherMethod,             // line #600:     ]);
		OtherMethod,             // line #601:
		OtherMethod,             // line #602:     List<Statement> statements =
		OtherMethod,             // line #603:         AstFinder.getStatementsInTopLevelFunction(unit, "main");
		OtherMethod,             // line #604:     DartType literal(int i) {
		OtherMethod,             // line #605:       ExpressionStatement stmt = statements[i];
		OtherMethod,             // line #606:       FunctionExpressionInvocation invk = stmt.expression;
		OtherMethod,             // line #607:       FunctionExpression exp = invk.argumentList.arguments[0];
		OtherMethod,             // line #608:       return exp.declaredElement.type;
		OtherMethod,             // line #609:     }
		OtherMethod,             // line #610:
		OtherMethod,             // line #611:     _isFunction2Of(_isInt, _isNull)(literal(0));
		OtherMethod,             // line #612:     _isFunction2Of(_isInt, _isString)(literal(1));
		OtherMethod,             // line #613:     _isFunction2Of(_isInt, _isString)(literal(2));
		OtherMethod,             // line #614:     _isFunction2Of(_isInt, _isString)(literal(3));
		OtherMethod,             // line #615:     _isFunction2Of(_isInt, _isString)(literal(4));
		OtherMethod,             // line #616:   }
		BlankLine,               // line #617:
		OtherMethod,             // line #618:   test_functionLiteral_functionInvocation_typedArguments() async {
		OtherMethod,             // line #619:     String code = r'''
		OtherMethod,             // line #620:       String map(String mapper(int x)) => mapper(null);
		OtherMethod,             // line #621:
		OtherMethod,             // line #622:       void main () {
		OtherMethod,             // line #623:         map((int x) => null);
		OtherMethod,             // line #624:         map((int x) => "hello");
		OtherMethod,             // line #625:         map((String x) => "hello");
		OtherMethod,             // line #626:         map((int x) => 3);
		OtherMethod,             // line #627:         map((int x) {return 3;});
		OtherMethod,             // line #628:      }
		OtherMethod,             // line #629:    ''';
		OtherMethod,             // line #630:     await assertErrorsInCode(code, [
		OtherMethod,             // line #631:       error(CompileTimeErrorCode.ARGUMENT_TYPE_NOT_ASSIGNABLE, 153, 21),
		OtherMethod,             // line #632:       error(CompileTimeErrorCode.RETURN_OF_INVALID_TYPE_FROM_CLOSURE, 200, 1),
		OtherMethod,             // line #633:       error(CompileTimeErrorCode.RETURN_OF_INVALID_TYPE_FROM_CLOSURE, 232, 1),
		OtherMethod,             // line #634:     ]);
		OtherMethod,             // line #635:
		OtherMethod,             // line #636:     List<Statement> statements =
		OtherMethod,             // line #637:         AstFinder.getStatementsInTopLevelFunction(unit, "main");
		OtherMethod,             // line #638:     DartType literal(int i) {
		OtherMethod,             // line #639:       ExpressionStatement stmt = statements[i];
		OtherMethod,             // line #640:       MethodInvocation invk = stmt.expression;
		OtherMethod,             // line #641:       FunctionExpression exp = invk.argumentList.arguments[0];
		OtherMethod,             // line #642:       return exp.declaredElement.type;
		OtherMethod,             // line #643:     }
		OtherMethod,             // line #644:
		OtherMethod,             // line #645:     _isFunction2Of(_isInt, _isNull)(literal(0));
		OtherMethod,             // line #646:     _isFunction2Of(_isInt, _isString)(literal(1));
		OtherMethod,             // line #647:     _isFunction2Of(_isString, _isString)(literal(2));
		OtherMethod,             // line #648:     _isFunction2Of(_isInt, _isString)(literal(3));
		OtherMethod,             // line #649:     _isFunction2Of(_isInt, _isString)(literal(4));
		OtherMethod,             // line #650:   }
		BlankLine,               // line #651:
		OtherMethod,             // line #652:   test_functionLiteral_functionInvocation_unTypedArguments() async {
		OtherMethod,             // line #653:     String code = r'''
		OtherMethod,             // line #654:       String map(String mapper(int x)) => mapper(null);
		OtherMethod,             // line #655:
		OtherMethod,             // line #656:       void main () {
		OtherMethod,             // line #657:         map((x) => null);
		OtherMethod,             // line #658:         map((x) => "hello");
		OtherMethod,             // line #659:         map((x) => "hello");
		OtherMethod,             // line #660:         map((x) => 3);
		OtherMethod,             // line #661:         map((x) {return 3;});
		OtherMethod,             // line #662:      }
		OtherMethod,             // line #663:    ''';
		OtherMethod,             // line #664:     await assertErrorsInCode(code, [
		OtherMethod,             // line #665:       error(CompileTimeErrorCode.RETURN_OF_INVALID_TYPE_FROM_CLOSURE, 181, 1),
		OtherMethod,             // line #666:       error(CompileTimeErrorCode.RETURN_OF_INVALID_TYPE_FROM_CLOSURE, 209, 1),
		OtherMethod,             // line #667:     ]);
		OtherMethod,             // line #668:
		OtherMethod,             // line #669:     List<Statement> statements =
		OtherMethod,             // line #670:         AstFinder.getStatementsInTopLevelFunction(unit, "main");
		OtherMethod,             // line #671:     DartType literal(int i) {
		OtherMethod,             // line #672:       ExpressionStatement stmt = statements[i];
		OtherMethod,             // line #673:       MethodInvocation invk = stmt.expression;
		OtherMethod,             // line #674:       FunctionExpression exp = invk.argumentList.arguments[0];
		OtherMethod,             // line #675:       return exp.declaredElement.type;
		OtherMethod,             // line #676:     }
		OtherMethod,             // line #677:
		OtherMethod,             // line #678:     _isFunction2Of(_isInt, _isNull)(literal(0));
		OtherMethod,             // line #679:     _isFunction2Of(_isInt, _isString)(literal(1));
		OtherMethod,             // line #680:     _isFunction2Of(_isInt, _isString)(literal(2));
		OtherMethod,             // line #681:     _isFunction2Of(_isInt, _isString)(literal(3));
		OtherMethod,             // line #682:     _isFunction2Of(_isInt, _isString)(literal(4));
		OtherMethod,             // line #683:   }
		BlankLine,               // line #684:
		OtherMethod,             // line #685:   test_functionLiteral_methodInvocation_typedArguments() async {
		OtherMethod,             // line #686:     String code = r'''
		OtherMethod,             // line #687:       class Mapper<F, T> {
		OtherMethod,             // line #688:         T map(T mapper(F x)) => mapper(null);
		OtherMethod,             // line #689:       }
		OtherMethod,             // line #690:
		OtherMethod,             // line #691:       void main () {
		OtherMethod,             // line #692:         new Mapper<int, String>().map((int x) => null);
		OtherMethod,             // line #693:         new Mapper<int, String>().map((int x) => "hello");
		OtherMethod,             // line #694:         new Mapper<int, String>().map((String x) => "hello");
		OtherMethod,             // line #695:         new Mapper<int, String>().map((int x) => 3);
		OtherMethod,             // line #696:         new Mapper<int, String>().map((int x) {return 3;});
		OtherMethod,             // line #697:      }
		OtherMethod,             // line #698:    ''';
		OtherMethod,             // line #699:     await assertErrorsInCode(code, [
		OtherMethod,             // line #700:       error(CompileTimeErrorCode.ARGUMENT_TYPE_NOT_ASSIGNABLE, 256, 21),
		OtherMethod,             // line #701:       error(CompileTimeErrorCode.RETURN_OF_INVALID_TYPE_FROM_CLOSURE, 329, 1),
		OtherMethod,             // line #702:       error(CompileTimeErrorCode.RETURN_OF_INVALID_TYPE_FROM_CLOSURE, 387, 1),
		OtherMethod,             // line #703:     ]);
		OtherMethod,             // line #704:
		OtherMethod,             // line #705:     List<Statement> statements =
		OtherMethod,             // line #706:         AstFinder.getStatementsInTopLevelFunction(unit, "main");
		OtherMethod,             // line #707:     DartType literal(int i) {
		OtherMethod,             // line #708:       ExpressionStatement stmt = statements[i];
		OtherMethod,             // line #709:       MethodInvocation invk = stmt.expression;
		OtherMethod,             // line #710:       FunctionExpression exp = invk.argumentList.arguments[0];
		OtherMethod,             // line #711:       return exp.declaredElement.type;
		OtherMethod,             // line #712:     }
		OtherMethod,             // line #713:
		OtherMethod,             // line #714:     _isFunction2Of(_isInt, _isNull)(literal(0));
		OtherMethod,             // line #715:     _isFunction2Of(_isInt, _isString)(literal(1));
		OtherMethod,             // line #716:     _isFunction2Of(_isString, _isString)(literal(2));
		OtherMethod,             // line #717:     _isFunction2Of(_isInt, _isString)(literal(3));
		OtherMethod,             // line #718:     _isFunction2Of(_isInt, _isString)(literal(4));
		OtherMethod,             // line #719:   }
		BlankLine,               // line #720:
		OtherMethod,             // line #721:   test_functionLiteral_methodInvocation_unTypedArguments() async {
		OtherMethod,             // line #722:     String code = r'''
		OtherMethod,             // line #723:       class Mapper<F, T> {
		OtherMethod,             // line #724:         T map(T mapper(F x)) => mapper(null);
		OtherMethod,             // line #725:       }
		OtherMethod,             // line #726:
		OtherMethod,             // line #727:       void main () {
		OtherMethod,             // line #728:         new Mapper<int, String>().map((x) => null);
		OtherMethod,             // line #729:         new Mapper<int, String>().map((x) => "hello");
		OtherMethod,             // line #730:         new Mapper<int, String>().map((x) => "hello");
		OtherMethod,             // line #731:         new Mapper<int, String>().map((x) => 3);
		OtherMethod,             // line #732:         new Mapper<int, String>().map((x) {return 3;});
		OtherMethod,             // line #733:      }
		OtherMethod,             // line #734:    ''';
		OtherMethod,             // line #735:     await assertErrorsInCode(code, [
		OtherMethod,             // line #736:       error(CompileTimeErrorCode.RETURN_OF_INVALID_TYPE_FROM_CLOSURE, 310, 1),
		OtherMethod,             // line #737:       error(CompileTimeErrorCode.RETURN_OF_INVALID_TYPE_FROM_CLOSURE, 364, 1),
		OtherMethod,             // line #738:     ]);
		OtherMethod,             // line #739:
		OtherMethod,             // line #740:     List<Statement> statements =
		OtherMethod,             // line #741:         AstFinder.getStatementsInTopLevelFunction(unit, "main");
		OtherMethod,             // line #742:     DartType literal(int i) {
		OtherMethod,             // line #743:       ExpressionStatement stmt = statements[i];
		OtherMethod,             // line #744:       MethodInvocation invk = stmt.expression;
		OtherMethod,             // line #745:       FunctionExpression exp = invk.argumentList.arguments[0];
		OtherMethod,             // line #746:       return exp.declaredElement.type;
		OtherMethod,             // line #747:     }
		OtherMethod,             // line #748:
		OtherMethod,             // line #749:     _isFunction2Of(_isInt, _isNull)(literal(0));
		OtherMethod,             // line #750:     _isFunction2Of(_isInt, _isString)(literal(1));
		OtherMethod,             // line #751:     _isFunction2Of(_isInt, _isString)(literal(2));
		OtherMethod,             // line #752:     _isFunction2Of(_isInt, _isString)(literal(3));
		OtherMethod,             // line #753:     _isFunction2Of(_isInt, _isString)(literal(4));
		OtherMethod,             // line #754:   }
		BlankLine,               // line #755:
		OtherMethod,             // line #756:   test_functionLiteral_unTypedArgument_propagation() async {
		OtherMethod,             // line #757:     String code = r'''
		OtherMethod,             // line #758:       typedef T Function2<S, T>(S x);
		OtherMethod,             // line #759:
		OtherMethod,             // line #760:       void main () {
		OtherMethod,             // line #761:         Function2<int, int> l0 = (x) => x;
		OtherMethod,             // line #762:         Function2<int, int> l1 = (x) => x+1;
		OtherMethod,             // line #763:         Function2<int, String> l2 = (x) => x;
		OtherMethod,             // line #764:         Function2<int, String> l3 = (x) => x.toLowerCase();
		OtherMethod,             // line #765:         Function2<String, String> l4 = (x) => x.toLowerCase();
		OtherMethod,             // line #766:      }
		OtherMethod,             // line #767:    ''';
		OtherMethod,             // line #768:     await assertErrorsInCode(code, [
		OtherMethod,             // line #769:       error(HintCode.UNUSED_LOCAL_VARIABLE, 88, 2),
		OtherMethod,             // line #770:       error(HintCode.UNUSED_LOCAL_VARIABLE, 131, 2),
		OtherMethod,             // line #771:       error(HintCode.UNUSED_LOCAL_VARIABLE, 179, 2),
		OtherMethod,             // line #772:       error(CompileTimeErrorCode.RETURN_OF_INVALID_TYPE_FROM_CLOSURE, 191, 1),
		OtherMethod,             // line #773:       error(HintCode.UNUSED_LOCAL_VARIABLE, 225, 2),
		OtherMethod,             // line #774:       error(CompileTimeErrorCode.UNDEFINED_METHOD, 239, 11),
		OtherMethod,             // line #775:       error(HintCode.UNUSED_LOCAL_VARIABLE, 288, 2),
		OtherMethod,             // line #776:     ]);
		OtherMethod,             // line #777:
		OtherMethod,             // line #778:     List<Statement> statements =
		OtherMethod,             // line #779:         AstFinder.getStatementsInTopLevelFunction(unit, "main");
		OtherMethod,             // line #780:     Expression functionReturnValue(int i) {
		OtherMethod,             // line #781:       VariableDeclarationStatement stmt = statements[i];
		OtherMethod,             // line #782:       VariableDeclaration decl = stmt.variables.variables[0];
		OtherMethod,             // line #783:       FunctionExpression exp = decl.initializer;
		OtherMethod,             // line #784:       FunctionBody body = exp.body;
		OtherMethod,             // line #785:       if (body is ExpressionFunctionBody) {
		OtherMethod,             // line #786:         return body.expression;
		OtherMethod,             // line #787:       } else {
		OtherMethod,             // line #788:         Statement stmt = (body as BlockFunctionBody).block.statements[0];
		OtherMethod,             // line #789:         return (stmt as ReturnStatement).expression;
		OtherMethod,             // line #790:       }
		OtherMethod,             // line #791:     }
		OtherMethod,             // line #792:
		OtherMethod,             // line #793:     expect(functionReturnValue(0).staticType, typeProvider.intType);
		OtherMethod,             // line #794:     expect(functionReturnValue(1).staticType, typeProvider.intType);
		OtherMethod,             // line #795:     expect(functionReturnValue(2).staticType, typeProvider.intType);
		OtherMethod,             // line #796:     expect(functionReturnValue(3).staticType, typeProvider.dynamicType);
		OtherMethod,             // line #797:     expect(functionReturnValue(4).staticType, typeProvider.stringType);
		OtherMethod,             // line #798:   }
		BlankLine,               // line #799:
		OtherMethod,             // line #800:   test_futureOr_assignFromFuture() async {
		OtherMethod,             // line #801:     // Test a Future<T> can be assigned to FutureOr<T>.
		OtherMethod,             // line #802:     MethodInvocation invoke = await _testFutureOr(r'''
		OtherMethod,             // line #803:     FutureOr<T> mk<T>(Future<T> x) => x;
		OtherMethod,             // line #804:     test() => mk(new Future<int>.value(42));
		OtherMethod,             // line #805:     ''');
		OtherMethod,             // line #806:     _isFutureOrOfInt(invoke.staticType);
		OtherMethod,             // line #807:   }
		BlankLine,               // line #808:
		OtherMethod,             // line #809:   test_futureOr_assignFromValue() async {
		OtherMethod,             // line #810:     // Test a T can be assigned to FutureOr<T>.
		OtherMethod,             // line #811:     MethodInvocation invoke = await _testFutureOr(r'''
		OtherMethod,             // line #812:     FutureOr<T> mk<T>(T x) => x;
		OtherMethod,             // line #813:     test() => mk(42);
		OtherMethod,             // line #814:     ''');
		OtherMethod,             // line #815:     _isFutureOrOfInt(invoke.staticType);
		OtherMethod,             // line #816:   }
		BlankLine,               // line #817:
		OtherMethod,             // line #818:   test_futureOr_asyncExpressionBody() async {
		OtherMethod,             // line #819:     // A FutureOr<T> can be used as the expression body for an async function
		OtherMethod,             // line #820:     MethodInvocation invoke = await _testFutureOr(r'''
		OtherMethod,             // line #821:     Future<T> mk<T>(FutureOr<T> x) async => x;
		OtherMethod,             // line #822:     test() => mk(42);
		OtherMethod,             // line #823:     ''');
		OtherMethod,             // line #824:     _isFutureOfInt(invoke.staticType);
		OtherMethod,             // line #825:   }
		BlankLine,               // line #826:
		OtherMethod,             // line #827:   test_futureOr_asyncReturn() async {
		OtherMethod,             // line #828:     // A FutureOr<T> can be used as the return value for an async function
		OtherMethod,             // line #829:     MethodInvocation invoke = await _testFutureOr(r'''
		OtherMethod,             // line #830:     Future<T> mk<T>(FutureOr<T> x) async { return x; }
		OtherMethod,             // line #831:     test() => mk(42);
		OtherMethod,             // line #832:     ''');
		OtherMethod,             // line #833:     _isFutureOfInt(invoke.staticType);
		OtherMethod,             // line #834:   }
		BlankLine,               // line #835:
		OtherMethod,             // line #836:   test_futureOr_await() async {
		OtherMethod,             // line #837:     // Test a FutureOr<T> can be awaited.
		OtherMethod,             // line #838:     MethodInvocation invoke = await _testFutureOr(r'''
		OtherMethod,             // line #839:     Future<T> mk<T>(FutureOr<T> x) async => await x;
		OtherMethod,             // line #840:     test() => mk(42);
		OtherMethod,             // line #841:     ''');
		OtherMethod,             // line #842:     _isFutureOfInt(invoke.staticType);
		OtherMethod,             // line #843:   }
		BlankLine,               // line #844:
		OtherMethod,             // line #845:   test_futureOr_downwards1() async {
		OtherMethod,             // line #846:     // Test that downwards inference interacts correctly with FutureOr
		OtherMethod,             // line #847:     // parameters.
		OtherMethod,             // line #848:     MethodInvocation invoke = await _testFutureOr(r'''
		OtherMethod,             // line #849:     Future<T> mk<T>(FutureOr<T> x) => null;
		OtherMethod,             // line #850:     Future<int> test() => mk(new Future<int>.value(42));
		OtherMethod,             // line #851:     ''');
		OtherMethod,             // line #852:     _isFutureOfInt(invoke.staticType);
		OtherMethod,             // line #853:   }
		BlankLine,               // line #854:
		OtherMethod,             // line #855:   test_futureOr_downwards2() async {
		OtherMethod,             // line #856:     // Test that downwards inference interacts correctly with FutureOr
		OtherMethod,             // line #857:     // parameters when the downwards context is FutureOr
		OtherMethod,             // line #858:     MethodInvocation invoke = await _testFutureOr(r'''
		OtherMethod,             // line #859:     Future<T> mk<T>(FutureOr<T> x) => null;
		OtherMethod,             // line #860:     FutureOr<int> test() => mk(new Future<int>.value(42));
		OtherMethod,             // line #861:     ''');
		OtherMethod,             // line #862:     _isFutureOfInt(invoke.staticType);
		OtherMethod,             // line #863:   }
		BlankLine,               // line #864:
		OtherMethod,             // line #865:   test_futureOr_downwards3() async {
		OtherMethod,             // line #866:     // Test that downwards inference correctly propogates into
		OtherMethod,             // line #867:     // arguments.
		OtherMethod,             // line #868:     MethodInvocation invoke = await _testFutureOr(r'''
		OtherMethod,             // line #869:     Future<T> mk<T>(FutureOr<T> x) => null;
		OtherMethod,             // line #870:     Future<int> test() => mk(new Future.value(42));
		OtherMethod,             // line #871:     ''');
		OtherMethod,             // line #872:     _isFutureOfInt(invoke.staticType);
		OtherMethod,             // line #873:     _isFutureOfInt(invoke.argumentList.arguments[0].staticType);
		OtherMethod,             // line #874:   }
		BlankLine,               // line #875:
		OtherMethod,             // line #876:   test_futureOr_downwards4() async {
		OtherMethod,             // line #877:     // Test that downwards inference interacts correctly with FutureOr
		OtherMethod,             // line #878:     // parameters when the downwards context is FutureOr
		OtherMethod,             // line #879:     MethodInvocation invoke = await _testFutureOr(r'''
		OtherMethod,             // line #880:     Future<T> mk<T>(FutureOr<T> x) => null;
		OtherMethod,             // line #881:     FutureOr<int> test() => mk(new Future.value(42));
		OtherMethod,             // line #882:     ''');
		OtherMethod,             // line #883:     _isFutureOfInt(invoke.staticType);
		OtherMethod,             // line #884:     _isFutureOfInt(invoke.argumentList.arguments[0].staticType);
		OtherMethod,             // line #885:   }
		BlankLine,               // line #886:
		OtherMethod,             // line #887:   test_futureOr_downwards5() async {
		OtherMethod,             // line #888:     // Test that downwards inference correctly pins the type when it
		OtherMethod,             // line #889:     // comes from a FutureOr
		OtherMethod,             // line #890:     MethodInvocation invoke = await _testFutureOr(r'''
		OtherMethod,             // line #891:     Future<T> mk<T>(FutureOr<T> x) => null;
		OtherMethod,             // line #892:     FutureOr<num> test() => mk(new Future.value(42));
		OtherMethod,             // line #893:     ''');
		OtherMethod,             // line #894:     _isFutureOf([_isNum])(invoke.staticType);
		OtherMethod,             // line #895:     _isFutureOf([_isNum])(invoke.argumentList.arguments[0].staticType);
		OtherMethod,             // line #896:   }
		BlankLine,               // line #897:
		OtherMethod,             // line #898:   test_futureOr_downwards6() async {
		OtherMethod,             // line #899:     // Test that downwards inference doesn't decompose FutureOr
		OtherMethod,             // line #900:     // when instantiating type variables.
		OtherMethod,             // line #901:     MethodInvocation invoke = await _testFutureOr(r'''
		OtherMethod,             // line #902:     T mk<T>(T x) => null;
		OtherMethod,             // line #903:     FutureOr<int> test() => mk(new Future.value(42));
		OtherMethod,             // line #904:     ''');
		OtherMethod,             // line #905:     _isFutureOrOfInt(invoke.staticType);
		OtherMethod,             // line #906:     _isFutureOfInt(invoke.argumentList.arguments[0].staticType);
		OtherMethod,             // line #907:   }
		BlankLine,               // line #908:
		OtherMethod,             // line #909:   test_futureOr_downwards7() async {
		OtherMethod,             // line #910:     // Test that downwards inference incorporates bounds correctly
		OtherMethod,             // line #911:     // when instantiating type variables.
		OtherMethod,             // line #912:     MethodInvocation invoke = await _testFutureOr(r'''
		OtherMethod,             // line #913:       T mk<T extends Future<int>>(T x) => null;
		OtherMethod,             // line #914:       FutureOr<int> test() => mk(new Future.value(42));
		OtherMethod,             // line #915:     ''');
		OtherMethod,             // line #916:     _isFutureOfInt(invoke.staticType);
		OtherMethod,             // line #917:     _isFutureOfInt(invoke.argumentList.arguments[0].staticType);
		OtherMethod,             // line #918:   }
		BlankLine,               // line #919:
		OtherMethod,             // line #920:   test_futureOr_downwards8() async {
		OtherMethod,             // line #921:     // Test that downwards inference incorporates bounds correctly
		OtherMethod,             // line #922:     // when instantiating type variables.
		OtherMethod,             // line #923:     // TODO(leafp): I think this should pass once the inference changes
		OtherMethod,             // line #924:     // that jmesserly is adding are landed.
		OtherMethod,             // line #925:     MethodInvocation invoke = await _testFutureOr(r'''
		OtherMethod,             // line #926:     T mk<T extends Future<Object>>(T x) => null;
		OtherMethod,             // line #927:     FutureOr<int> test() => mk(new Future.value(42));
		OtherMethod,             // line #928:     ''');
		OtherMethod,             // line #929:     _isFutureOfInt(invoke.staticType);
		OtherMethod,             // line #930:     _isFutureOfInt(invoke.argumentList.arguments[0].staticType);
		OtherMethod,             // line #931:   }
		BlankLine,               // line #932:
		OtherMethod,             // line #933:   test_futureOr_downwards9() async {
		OtherMethod,             // line #934:     // Test that downwards inference decomposes correctly with
		OtherMethod,             // line #935:     // other composite types
		OtherMethod,             // line #936:     MethodInvocation invoke = await _testFutureOr(r'''
		OtherMethod,             // line #937:     List<T> mk<T>(T x) => null;
		OtherMethod,             // line #938:     FutureOr<List<int>> test() => mk(3);
		OtherMethod,             // line #939:     ''');
		OtherMethod,             // line #940:     _isListOf(_isInt)(invoke.staticType);
		OtherMethod,             // line #941:     _isInt(invoke.argumentList.arguments[0].staticType);
		OtherMethod,             // line #942:   }
		BlankLine,               // line #943:
		OtherMethod,             // line #944:   test_futureOr_methods1() async {
		OtherMethod,             // line #945:     // Test that FutureOr has the Object methods
		OtherMethod,             // line #946:     MethodInvocation invoke = await _testFutureOr(r'''
		OtherMethod,             // line #947:     dynamic test(FutureOr<int> x) => x.toString();
		OtherMethod,             // line #948:     ''');
		OtherMethod,             // line #949:     _isString(invoke.staticType);
		OtherMethod,             // line #950:   }
		BlankLine,               // line #951:
		OtherMethod,             // line #952:   test_futureOr_methods2() async {
		OtherMethod,             // line #953:     // Test that FutureOr does not have the constituent type methods
		OtherMethod,             // line #954:     MethodInvocation invoke = await _testFutureOr(r'''
		OtherMethod,             // line #955:     dynamic test(FutureOr<int> x) => x.abs();
		OtherMethod,             // line #956:     ''', expectedErrors: [
		OtherMethod,             // line #957:       error(CompileTimeErrorCode.UNDEFINED_METHOD, 61, 3),
		OtherMethod,             // line #958:     ]);
		OtherMethod,             // line #959:     _isDynamic(invoke.staticType);
		OtherMethod,             // line #960:   }
		BlankLine,               // line #961:
		OtherMethod,             // line #962:   test_futureOr_methods3() async {
		OtherMethod,             // line #963:     // Test that FutureOr does not have the Future type methods
		OtherMethod,             // line #964:     MethodInvocation invoke = await _testFutureOr(r'''
		OtherMethod,             // line #965:     dynamic test(FutureOr<int> x) => x.then((x) => x);
		OtherMethod,             // line #966:     ''', expectedErrors: [
		OtherMethod,             // line #967:       error(CompileTimeErrorCode.UNDEFINED_METHOD, 61, 4),
		OtherMethod,             // line #968:     ]);
		OtherMethod,             // line #969:     _isDynamic(invoke.staticType);
		OtherMethod,             // line #970:   }
		BlankLine,               // line #971:
		OtherMethod,             // line #972:   test_futureOr_methods4() async {
		OtherMethod,             // line #973:     // Test that FutureOr<dynamic> does not have all methods
		OtherMethod,             // line #974:     MethodInvocation invoke = await _testFutureOr(r'''
		OtherMethod,             // line #975:     dynamic test(FutureOr<dynamic> x) => x.abs();
		OtherMethod,             // line #976:     ''', expectedErrors: [
		OtherMethod,             // line #977:       error(CompileTimeErrorCode.UNDEFINED_METHOD, 65, 3),
		OtherMethod,             // line #978:     ]);
		OtherMethod,             // line #979:     _isDynamic(invoke.staticType);
		OtherMethod,             // line #980:   }
		BlankLine,               // line #981:
		OtherMethod,             // line #982:   test_futureOr_no_return() async {
		OtherMethod,             // line #983:     MethodInvocation invoke = await _testFutureOr(r'''
		OtherMethod,             // line #984:     FutureOr<T> mk<T>(Future<T> x) => x;
		OtherMethod,             // line #985:     Future<int> f;
		OtherMethod,             // line #986:     test() => f.then((int x) {});
		OtherMethod,             // line #987:     ''');
		OtherMethod,             // line #988:     _isFunction2Of(_isInt, _isNull)(
		OtherMethod,             // line #989:         invoke.argumentList.arguments[0].staticType);
		OtherMethod,             // line #990:     _isFutureOfNull(invoke.staticType);
		OtherMethod,             // line #991:   }
		BlankLine,               // line #992:
		OtherMethod,             // line #993:   test_futureOr_no_return_value() async {
		OtherMethod,             // line #994:     MethodInvocation invoke = await _testFutureOr(r'''
		OtherMethod,             // line #995:     FutureOr<T> mk<T>(Future<T> x) => x;
		OtherMethod,             // line #996:     Future<int> f;
		OtherMethod,             // line #997:     test() => f.then((int x) {return;});
		OtherMethod,             // line #998:     ''');
		OtherMethod,             // line #999:     _isFunction2Of(_isInt, _isNull)(
		OtherMethod,             // line #1000:         invoke.argumentList.arguments[0].staticType);
		OtherMethod,             // line #1001:     _isFutureOfNull(invoke.staticType);
		OtherMethod,             // line #1002:   }
		BlankLine,               // line #1003:
		OtherMethod,             // line #1004:   test_futureOr_return_null() async {
		OtherMethod,             // line #1005:     MethodInvocation invoke = await _testFutureOr(r'''
		OtherMethod,             // line #1006:     FutureOr<T> mk<T>(Future<T> x) => x;
		OtherMethod,             // line #1007:     Future<int> f;
		OtherMethod,             // line #1008:     test() => f.then((int x) {return null;});
		OtherMethod,             // line #1009:     ''');
		OtherMethod,             // line #1010:     _isFunction2Of(_isInt, _isNull)(
		OtherMethod,             // line #1011:         invoke.argumentList.arguments[0].staticType);
		OtherMethod,             // line #1012:     _isFutureOfNull(invoke.staticType);
		OtherMethod,             // line #1013:   }
		BlankLine,               // line #1014:
		OtherMethod,             // line #1015:   test_futureOr_upwards1() async {
		OtherMethod,             // line #1016:     // Test that upwards inference correctly prefers to instantiate type
		OtherMethod,             // line #1017:     // variables with the "smaller" solution when both are possible.
		OtherMethod,             // line #1018:     MethodInvocation invoke = await _testFutureOr(r'''
		OtherMethod,             // line #1019:     Future<T> mk<T>(FutureOr<T> x) => null;
		OtherMethod,             // line #1020:     dynamic test() => mk(new Future<int>.value(42));
		OtherMethod,             // line #1021:     ''');
		OtherMethod,             // line #1022:     _isFutureOfInt(invoke.staticType);
		OtherMethod,             // line #1023:   }
		BlankLine,               // line #1024:
		OtherMethod,             // line #1025:   test_futureOr_upwards2() async {
		OtherMethod,             // line #1026:     // Test that upwards inference fails when the solution doesn't
		OtherMethod,             // line #1027:     // match the bound.
		OtherMethod,             // line #1028:     MethodInvocation invoke = await _testFutureOr(r'''
		OtherMethod,             // line #1029:     Future<T> mk<T extends Future<Object>>(FutureOr<T> x) => null;
		OtherMethod,             // line #1030:     dynamic test() => mk(new Future<int>.value(42));
		OtherMethod,             // line #1031:     ''', expectedErrors: [
		OtherMethod,             // line #1032:       error(CompileTimeErrorCode.COULD_NOT_INFER, 111, 2),
		OtherMethod,             // line #1033:     ]);
		OtherMethod,             // line #1034:     _isFutureOfInt(invoke.staticType);
		OtherMethod,             // line #1035:   }
		BlankLine,               // line #1036:
		OtherMethod,             // line #1037:   test_futureOrNull_no_return() async {
		OtherMethod,             // line #1038:     MethodInvocation invoke = await _testFutureOr(r'''
		OtherMethod,             // line #1039:     FutureOr<T> mk<T>(Future<T> x) => x;
		OtherMethod,             // line #1040:     Future<int> f;
		OtherMethod,             // line #1041:     test() => f.then<Null>((int x) {});
		OtherMethod,             // line #1042:     ''');
		OtherMethod,             // line #1043:     _isFunction2Of(_isInt, _isNull)(
		OtherMethod,             // line #1044:         invoke.argumentList.arguments[0].staticType);
		OtherMethod,             // line #1045:     _isFutureOfNull(invoke.staticType);
		OtherMethod,             // line #1046:   }
		BlankLine,               // line #1047:
		OtherMethod,             // line #1048:   test_futureOrNull_no_return_value() async {
		OtherMethod,             // line #1049:     MethodInvocation invoke = await _testFutureOr(r'''
		OtherMethod,             // line #1050:     FutureOr<T> mk<T>(Future<T> x) => x;
		OtherMethod,             // line #1051:     Future<int> f;
		OtherMethod,             // line #1052:     test() => f.then<Null>((int x) {return;});
		OtherMethod,             // line #1053:     ''');
		OtherMethod,             // line #1054:     _isFunction2Of(_isInt, _isNull)(
		OtherMethod,             // line #1055:         invoke.argumentList.arguments[0].staticType);
		OtherMethod,             // line #1056:     _isFutureOfNull(invoke.staticType);
		OtherMethod,             // line #1057:   }
		BlankLine,               // line #1058:
		OtherMethod,             // line #1059:   test_futureOrNull_return_null() async {
		OtherMethod,             // line #1060:     MethodInvocation invoke = await _testFutureOr(r'''
		OtherMethod,             // line #1061:     FutureOr<T> mk<T>(Future<T> x) => x;
		OtherMethod,             // line #1062:     Future<int> f;
		OtherMethod,             // line #1063:     test() => f.then<Null>((int x) { return null;});
		OtherMethod,             // line #1064:     ''');
		OtherMethod,             // line #1065:     _isFunction2Of(_isInt, _isNull)(
		OtherMethod,             // line #1066:         invoke.argumentList.arguments[0].staticType);
		OtherMethod,             // line #1067:     _isFutureOfNull(invoke.staticType);
		OtherMethod,             // line #1068:   }
		BlankLine,               // line #1069:
		OtherMethod,             // line #1070:   test_generic_partial() async {
		OtherMethod,             // line #1071:     // Test that upward and downward type inference handles partial
		OtherMethod,             // line #1072:     // type schemas correctly.  Downwards inference in a partial context
		OtherMethod,             // line #1073:     // (e.g. Map<String, ?>) should still allow upwards inference to fill
		OtherMethod,             // line #1074:     // in the missing information.
		OtherMethod,             // line #1075:     String code = r'''
		OtherMethod,             // line #1076: class A<T> {
		OtherMethod,             // line #1077:   A(T x);
		OtherMethod,             // line #1078:   A.fromA(A<T> a) {}
		OtherMethod,             // line #1079:   A.fromMap(Map<String, T> m) {}
		OtherMethod,             // line #1080:   A.fromList(List<T> m) {}
		OtherMethod,             // line #1081:   A.fromT(T t) {}
		OtherMethod,             // line #1082:   A.fromB(B<T, String> a) {}
		OtherMethod,             // line #1083: }
		OtherMethod,             // line #1084:
		OtherMethod,             // line #1085: class B<S, T> {
		OtherMethod,             // line #1086:   B(S s);
		OtherMethod,             // line #1087: }
		OtherMethod,             // line #1088:
		OtherMethod,             // line #1089: void test() {
		OtherMethod,             // line #1090:     var a0 = new A.fromA(new A(3));
		OtherMethod,             // line #1091:     var a1 = new A.fromMap({'hello' : 3});
		OtherMethod,             // line #1092:     var a2 = new A.fromList([3]);
		OtherMethod,             // line #1093:     var a3 = new A.fromT(3);
		OtherMethod,             // line #1094:     var a4 = new A.fromB(new B(3));
		OtherMethod,             // line #1095: }
		OtherMethod,             // line #1096:    ''';
		OtherMethod,             // line #1097:     await assertErrorsInCode(code, [
		OtherMethod,             // line #1098:       error(HintCode.UNUSED_LOCAL_VARIABLE, 205, 2),
		OtherMethod,             // line #1099:       error(HintCode.UNUSED_LOCAL_VARIABLE, 241, 2),
		OtherMethod,             // line #1100:       error(HintCode.UNUSED_LOCAL_VARIABLE, 284, 2),
		OtherMethod,             // line #1101:       error(HintCode.UNUSED_LOCAL_VARIABLE, 318, 2),
		OtherMethod,             // line #1102:       error(HintCode.UNUSED_LOCAL_VARIABLE, 347, 2),
		OtherMethod,             // line #1103:     ]);
		OtherMethod,             // line #1104:
		OtherMethod,             // line #1105:     Element elementA = AstFinder.getClass(unit, "A").declaredElement;
		OtherMethod,             // line #1106:     List<Statement> statements =
		OtherMethod,             // line #1107:         AstFinder.getStatementsInTopLevelFunction(unit, "test");
		OtherMethod,             // line #1108:     void check(int i) {
		OtherMethod,             // line #1109:       VariableDeclarationStatement stmt = statements[i];
		OtherMethod,             // line #1110:       VariableDeclaration decl = stmt.variables.variables[0];
		OtherMethod,             // line #1111:       Expression init = decl.initializer;
		OtherMethod,             // line #1112:       _isInstantiationOf(_hasElement(elementA))([_isInt])(init.staticType);
		OtherMethod,             // line #1113:     }
		OtherMethod,             // line #1114:
		OtherMethod,             // line #1115:     for (var i = 0; i < 5; i++) {
		OtherMethod,             // line #1116:       check(i);
		OtherMethod,             // line #1117:     }
		OtherMethod,             // line #1118:   }
		BlankLine,               // line #1119:
		OtherMethod,             // line #1120:   test_inferConstructor_unknownTypeLowerBound() async {
		OtherMethod,             // line #1121:     var code = r'''
		OtherMethod,             // line #1122:         class C<T> {
		OtherMethod,             // line #1123:           C(void callback(List<T> a));
		OtherMethod,             // line #1124:         }
		OtherMethod,             // line #1125:         test() {
		OtherMethod,             // line #1126:           // downwards inference pushes List<?> and in parameter position this
		OtherMethod,             // line #1127:           // becomes inferred as List<Null>.
		OtherMethod,             // line #1128:           var c = new C((items) {});
		OtherMethod,             // line #1129:         }
		OtherMethod,             // line #1130:         ''';
		OtherMethod,             // line #1131:     await assertErrorsInCode(code, [
		OtherMethod,             // line #1132:       error(HintCode.UNUSED_LOCAL_VARIABLE, 225, 1),
		OtherMethod,             // line #1133:     ]);
		OtherMethod,             // line #1134:
		OtherMethod,             // line #1135:     DartType cType = findLocalVariable(unit, 'c').type;
		OtherMethod,             // line #1136:     Element elementC = AstFinder.getClass(unit, "C").declaredElement;
		OtherMethod,             // line #1137:
		OtherMethod,             // line #1138:     _isInstantiationOf(_hasElement(elementC))([_isDynamic])(cType);
		OtherMethod,             // line #1139:   }
		BlankLine,               // line #1140:
		OtherMethod,             // line #1141:   test_inference_error_arguments() async {
		OtherMethod,             // line #1142:     var code = r'''
		OtherMethod,             // line #1143: typedef R F<T, R>(T t);
		OtherMethod,             // line #1144:
		OtherMethod,             // line #1145: F<T, T> g<T>(F<T, T> f) => (x) => f(f(x));
		OtherMethod,             // line #1146:
		OtherMethod,             // line #1147: test() {
		OtherMethod,             // line #1148:   var h = g((int x) => 42.0);
		OtherMethod,             // line #1149: }
		OtherMethod,             // line #1150:  ''';
		OtherMethod,             // line #1151:     await assertErrorsInCode(code, [
		OtherMethod,             // line #1152:       error(HintCode.UNUSED_LOCAL_VARIABLE, 84, 1),
		OtherMethod,             // line #1153:       error(CompileTimeErrorCode.COULD_NOT_INFER, 88, 1),
		OtherMethod,             // line #1154:       error(CompileTimeErrorCode.ARGUMENT_TYPE_NOT_ASSIGNABLE, 90, 15),
		OtherMethod,             // line #1155:     ]);
		OtherMethod,             // line #1156:     _expectInferenceError(r'''
		OtherMethod,             // line #1157: Couldn't infer type parameter 'T'.
		OtherMethod,             // line #1158:
		OtherMethod,             // line #1159: Tried to infer 'double' for 'T' which doesn't work:
		OtherMethod,             // line #1160:   Parameter 'f' declared as     'T Function(T)'
		OtherMethod,             // line #1161:                 but argument is 'double Function(int)'.
		OtherMethod,             // line #1162:
		OtherMethod,             // line #1163: Consider passing explicit type argument(s) to the generic.
		OtherMethod,             // line #1164:
		OtherMethod,             // line #1165: ''');
		OtherMethod,             // line #1166:   }
		BlankLine,               // line #1167:
		OtherMethod,             // line #1168:   test_inference_error_arguments2() async {
		OtherMethod,             // line #1169:     var code = r'''
		OtherMethod,             // line #1170: typedef R F<T, R>(T t);
		OtherMethod,             // line #1171:
		OtherMethod,             // line #1172: F<T, T> g<T>(F<T, T> a, F<T, T> b) => (x) => a(b(x));
		OtherMethod,             // line #1173:
		OtherMethod,             // line #1174: test() {
		OtherMethod,             // line #1175:   var h = g((int x) => 42.0, (double x) => 42);
		OtherMethod,             // line #1176: }
		OtherMethod,             // line #1177:  ''';
		OtherMethod,             // line #1178:     await assertErrorsInCode(code, [
		OtherMethod,             // line #1179:       error(HintCode.UNUSED_LOCAL_VARIABLE, 95, 1),
		OtherMethod,             // line #1180:       error(CompileTimeErrorCode.COULD_NOT_INFER, 99, 1),
		OtherMethod,             // line #1181:       error(CompileTimeErrorCode.ARGUMENT_TYPE_NOT_ASSIGNABLE, 101, 15),
		OtherMethod,             // line #1182:       error(CompileTimeErrorCode.ARGUMENT_TYPE_NOT_ASSIGNABLE, 118, 16),
		OtherMethod,             // line #1183:     ]);
		OtherMethod,             // line #1184:     _expectInferenceError(r'''
		OtherMethod,             // line #1185: Couldn't infer type parameter 'T'.
		OtherMethod,             // line #1186:
		OtherMethod,             // line #1187: Tried to infer 'num' for 'T' which doesn't work:
		OtherMethod,             // line #1188:   Parameter 'a' declared as     'T Function(T)'
		OtherMethod,             // line #1189:                 but argument is 'double Function(int)'.
		OtherMethod,             // line #1190:   Parameter 'b' declared as     'T Function(T)'
		OtherMethod,             // line #1191:                 but argument is 'int Function(double)'.
		OtherMethod,             // line #1192:
		OtherMethod,             // line #1193: Consider passing explicit type argument(s) to the generic.
		OtherMethod,             // line #1194:
		OtherMethod,             // line #1195: ''');
		OtherMethod,             // line #1196:   }
		BlankLine,               // line #1197:
		OtherMethod,             // line #1198:   test_inference_error_extendsFromReturn() async {
		OtherMethod,             // line #1199:     // This is not an inference error because we successfully infer Null.
		OtherMethod,             // line #1200:     var code = r'''
		OtherMethod,             // line #1201: T max<T extends num>(T x, T y) => x;
		OtherMethod,             // line #1202:
		OtherMethod,             // line #1203: test() {
		OtherMethod,             // line #1204:   String hello = max(1, 2);
		OtherMethod,             // line #1205: }
		OtherMethod,             // line #1206:  ''';
		OtherMethod,             // line #1207:     await assertErrorsInCode(code, [
		OtherMethod,             // line #1208:       error(HintCode.UNUSED_LOCAL_VARIABLE, 56, 5),
		OtherMethod,             // line #1209:       error(CompileTimeErrorCode.INVALID_CAST_LITERAL, 68, 1),
		OtherMethod,             // line #1210:       error(CompileTimeErrorCode.INVALID_CAST_LITERAL, 71, 1),
		OtherMethod,             // line #1211:     ]);
		OtherMethod,             // line #1212:
		OtherMethod,             // line #1213:     var h = (AstFinder.getStatementsInTopLevelFunction(unit, "test")[0]
		OtherMethod,             // line #1214:             as VariableDeclarationStatement)
		OtherMethod,             // line #1215:         .variables
		OtherMethod,             // line #1216:         .variables[0];
		OtherMethod,             // line #1217:     var call = h.initializer as MethodInvocation;
		OtherMethod,             // line #1218:     assertInvokeType(call, 'Null Function(Null, Null)');
		OtherMethod,             // line #1219:   }
		BlankLine,               // line #1220:
		OtherMethod,             // line #1221:   test_inference_error_extendsFromReturn2() async {
		OtherMethod,             // line #1222:     var code = r'''
		OtherMethod,             // line #1223: typedef R F<T, R>(T t);
		OtherMethod,             // line #1224: F<T, T> g<T extends num>() => (y) => y;
		OtherMethod,             // line #1225:
		OtherMethod,             // line #1226: test() {
		OtherMethod,             // line #1227:   F<String, String> hello = g();
		OtherMethod,             // line #1228: }
		OtherMethod,             // line #1229:  ''';
		OtherMethod,             // line #1230:     await assertErrorsInCode(code, [
		OtherMethod,             // line #1231:       error(HintCode.UNUSED_LOCAL_VARIABLE, 94, 5),
		OtherMethod,             // line #1232:       error(CompileTimeErrorCode.COULD_NOT_INFER, 102, 1),
		OtherMethod,             // line #1233:     ]);
		OtherMethod,             // line #1234:
		OtherMethod,             // line #1235:     _expectInferenceError(r'''
		OtherMethod,             // line #1236: Couldn't infer type parameter 'T'.
		OtherMethod,             // line #1237:
		OtherMethod,             // line #1238: Tried to infer 'String' for 'T' which doesn't work:
		OtherMethod,             // line #1239:   Type parameter 'T' declared to extend 'num'.
		OtherMethod,             // line #1240: The type 'String' was inferred from:
		OtherMethod,             // line #1241:   Return type declared as 'T Function(T)'
		OtherMethod,             // line #1242:               used where  'String Function(String)' is required.
		OtherMethod,             // line #1243:
		OtherMethod,             // line #1244: Consider passing explicit type argument(s) to the generic.
		OtherMethod,             // line #1245:
		OtherMethod,             // line #1246: ''');
		OtherMethod,             // line #1247:   }
		BlankLine,               // line #1248:
		OtherMethod,             // line #1249:   test_inference_error_genericFunction() async {
		OtherMethod,             // line #1250:     var code = r'''
		OtherMethod,             // line #1251: T max<T extends num>(T x, T y) => x < y ? y : x;
		OtherMethod,             // line #1252: abstract class Iterable<T> {
		OtherMethod,             // line #1253:   T get first;
		OtherMethod,             // line #1254:   S fold<S>(S s, S f(S s, T t));
		OtherMethod,             // line #1255: }
		OtherMethod,             // line #1256: test(Iterable values) {
		OtherMethod,             // line #1257:   num n = values.fold(values.first as num, max);
		OtherMethod,             // line #1258: }
		OtherMethod,             // line #1259:  ''';
		OtherMethod,             // line #1260:     await assertErrorsInCode(code, [
		OtherMethod,             // line #1261:       error(HintCode.UNUSED_LOCAL_VARIABLE, 158, 1),
		OtherMethod,             // line #1262:       error(CompileTimeErrorCode.ARGUMENT_TYPE_NOT_ASSIGNABLE, 195, 3),
		OtherMethod,             // line #1263:       error(CompileTimeErrorCode.COULD_NOT_INFER, 195, 3),
		OtherMethod,             // line #1264:     ]);
		OtherMethod,             // line #1265:     _expectInferenceError(r'''
		OtherMethod,             // line #1266: Couldn't infer type parameter 'T'.
		OtherMethod,             // line #1267:
		OtherMethod,             // line #1268: Tried to infer 'dynamic' for 'T' which doesn't work:
		OtherMethod,             // line #1269:   Function type declared as 'T Function<T extends num>(T, T)'
		OtherMethod,             // line #1270:                 used where  'num Function(num, dynamic)' is required.
		OtherMethod,             // line #1271:
		OtherMethod,             // line #1272: Consider passing explicit type argument(s) to the generic.
		OtherMethod,             // line #1273:
		OtherMethod,             // line #1274: ''');
		OtherMethod,             // line #1275:   }
		BlankLine,               // line #1276:
		OtherMethod,             // line #1277:   test_inference_error_returnContext() async {
		OtherMethod,             // line #1278:     var code = r'''
		OtherMethod,             // line #1279: typedef R F<T, R>(T t);
		OtherMethod,             // line #1280:
		OtherMethod,             // line #1281: F<T, T> g<T>(T t) => (x) => t;
		OtherMethod,             // line #1282:
		OtherMethod,             // line #1283: test() {
		OtherMethod,             // line #1284:   F<num, int> h = g(42);
		OtherMethod,             // line #1285: }
		OtherMethod,             // line #1286:  ''';
		OtherMethod,             // line #1287:     await assertErrorsInCode(code, [
		OtherMethod,             // line #1288:       error(HintCode.UNUSED_LOCAL_VARIABLE, 80, 1),
		OtherMethod,             // line #1289:       error(CompileTimeErrorCode.COULD_NOT_INFER, 84, 1),
		OtherMethod,             // line #1290:     ]);
		OtherMethod,             // line #1291:     _expectInferenceError(r'''
		OtherMethod,             // line #1292: Couldn't infer type parameter 'T'.
		OtherMethod,             // line #1293:
		OtherMethod,             // line #1294: Tried to infer 'num' for 'T' which doesn't work:
		OtherMethod,             // line #1295:   Return type declared as 'T Function(T)'
		OtherMethod,             // line #1296:               used where  'int Function(num)' is required.
		OtherMethod,             // line #1297:
		OtherMethod,             // line #1298: Consider passing explicit type argument(s) to the generic.
		OtherMethod,             // line #1299:
		OtherMethod,             // line #1300: ''');
		OtherMethod,             // line #1301:   }
		BlankLine,               // line #1302:
		OtherMethod,             // line #1303:   test_inference_hints() async {
		OtherMethod,             // line #1304:     var code = r'''
		OtherMethod,             // line #1305:       void main () {
		OtherMethod,             // line #1306:         var x = 3;
		OtherMethod,             // line #1307:         List<int> l0 = [];
		OtherMethod,             // line #1308:      }
		OtherMethod,             // line #1309:    ''';
		OtherMethod,             // line #1310:     await assertErrorsInCode(code, [
		OtherMethod,             // line #1311:       error(HintCode.UNUSED_LOCAL_VARIABLE, 33, 1),
		OtherMethod,             // line #1312:       error(HintCode.UNUSED_LOCAL_VARIABLE, 58, 2),
		OtherMethod,             // line #1313:     ]);
		OtherMethod,             // line #1314:   }
		BlankLine,               // line #1315:
		OtherMethod,             // line #1316:   test_inference_simplePolymorphicRecursion_function() async {
		OtherMethod,             // line #1317:     // Regression test for https://github.com/dart-lang/sdk/issues/30980
		OtherMethod,             // line #1318:     // Check that inference works properly when inferring the type argument
		OtherMethod,             // line #1319:     // for a self-recursive call with a function type
		OtherMethod,             // line #1320:     var code = r'''
		OtherMethod,             // line #1321: void _mergeSort<T>(T Function(T) list, int compare(T a, T b), T Function(T) target) {
		OtherMethod,             // line #1322:   _mergeSort(list, compare, target);
		OtherMethod,             // line #1323:   _mergeSort(list, compare, list);
		OtherMethod,             // line #1324:   _mergeSort(target, compare, target);
		OtherMethod,             // line #1325:   _mergeSort(target, compare, list);
		OtherMethod,             // line #1326: }
		OtherMethod,             // line #1327:     ''';
		OtherMethod,             // line #1328:     await assertErrorsInCode(code, [
		OtherMethod,             // line #1329:       error(HintCode.UNUSED_ELEMENT, 5, 10),
		OtherMethod,             // line #1330:     ]);
		OtherMethod,             // line #1331:
		OtherMethod,             // line #1332:     var body = AstFinder.getTopLevelFunction(unit, '_mergeSort')
		OtherMethod,             // line #1333:         .functionExpression
		OtherMethod,             // line #1334:         .body as BlockFunctionBody;
		OtherMethod,             // line #1335:     var stmts = body.block.statements;
		OtherMethod,             // line #1336:     for (ExpressionStatement stmt in stmts) {
		OtherMethod,             // line #1337:       MethodInvocation invoke = stmt.expression;
		OtherMethod,             // line #1338:       assertInvokeType(invoke,
		OtherMethod,             // line #1339:           'void Function(T Function(T), int Function(T, T), T Function(T))');
		OtherMethod,             // line #1340:     }
		OtherMethod,             // line #1341:   }
		BlankLine,               // line #1342:
		OtherMethod,             // line #1343:   test_inference_simplePolymorphicRecursion_interface() async {
		OtherMethod,             // line #1344:     // Regression test for https://github.com/dart-lang/sdk/issues/30980
		OtherMethod,             // line #1345:     // Check that inference works properly when inferring the type argument
		OtherMethod,             // line #1346:     // for a self-recursive call with an interface type
		OtherMethod,             // line #1347:     var code = r'''
		OtherMethod,             // line #1348: void _mergeSort<T>(List<T> list, int compare(T a, T b), List<T> target) {
		OtherMethod,             // line #1349:   _mergeSort(list, compare, target);
		OtherMethod,             // line #1350:   _mergeSort(list, compare, list);
		OtherMethod,             // line #1351:   _mergeSort(target, compare, target);
		OtherMethod,             // line #1352:   _mergeSort(target, compare, list);
		OtherMethod,             // line #1353: }
		OtherMethod,             // line #1354:     ''';
		OtherMethod,             // line #1355:     await assertErrorsInCode(code, [
		OtherMethod,             // line #1356:       error(HintCode.UNUSED_ELEMENT, 5, 10),
		OtherMethod,             // line #1357:     ]);
		OtherMethod,             // line #1358:
		OtherMethod,             // line #1359:     var body = AstFinder.getTopLevelFunction(unit, '_mergeSort')
		OtherMethod,             // line #1360:         .functionExpression
		OtherMethod,             // line #1361:         .body as BlockFunctionBody;
		OtherMethod,             // line #1362:     var stmts = body.block.statements;
		OtherMethod,             // line #1363:     for (ExpressionStatement stmt in stmts) {
		OtherMethod,             // line #1364:       MethodInvocation invoke = stmt.expression;
		OtherMethod,             // line #1365:       assertInvokeType(
		OtherMethod,             // line #1366:           invoke, 'void Function(List<T>, int Function(T, T), List<T>)');
		OtherMethod,             // line #1367:     }
		OtherMethod,             // line #1368:   }
		BlankLine,               // line #1369:
		OtherMethod,             // line #1370:   test_inference_simplePolymorphicRecursion_simple() async {
		OtherMethod,             // line #1371:     // Regression test for https://github.com/dart-lang/sdk/issues/30980
		OtherMethod,             // line #1372:     // Check that inference works properly when inferring the type argument
		OtherMethod,             // line #1373:     // for a self-recursive call with a simple type parameter
		OtherMethod,             // line #1374:     var code = r'''
		OtherMethod,             // line #1375: void _mergeSort<T>(T list, int compare(T a, T b), T target) {
		OtherMethod,             // line #1376:   _mergeSort(list, compare, target);
		OtherMethod,             // line #1377:   _mergeSort(list, compare, list);
		OtherMethod,             // line #1378:   _mergeSort(target, compare, target);
		OtherMethod,             // line #1379:   _mergeSort(target, compare, list);
		OtherMethod,             // line #1380: }
		OtherMethod,             // line #1381:     ''';
		OtherMethod,             // line #1382:     await assertErrorsInCode(code, [
		OtherMethod,             // line #1383:       error(HintCode.UNUSED_ELEMENT, 5, 10),
		OtherMethod,             // line #1384:     ]);
		OtherMethod,             // line #1385:
		OtherMethod,             // line #1386:     var body = AstFinder.getTopLevelFunction(unit, '_mergeSort')
		OtherMethod,             // line #1387:         .functionExpression
		OtherMethod,             // line #1388:         .body as BlockFunctionBody;
		OtherMethod,             // line #1389:     var stmts = body.block.statements;
		OtherMethod,             // line #1390:     for (ExpressionStatement stmt in stmts) {
		OtherMethod,             // line #1391:       MethodInvocation invoke = stmt.expression;
		OtherMethod,             // line #1392:       assertInvokeType(invoke, 'void Function(T, int Function(T, T), T)');
		OtherMethod,             // line #1393:     }
		OtherMethod,             // line #1394:   }
		BlankLine,               // line #1395:
		OtherMethod,             // line #1396:   test_inferGenericInstantiation() async {
		OtherMethod,             // line #1397:     // Verify that we don't infer '?` when we instantiate a generic function.
		OtherMethod,             // line #1398:     var code = r'''
		OtherMethod,             // line #1399: T f<T>(T x(T t)) => x(null);
		OtherMethod,             // line #1400: S g<S>(S s) => s;
		OtherMethod,             // line #1401: test() {
		OtherMethod,             // line #1402:  var h = f(g);
		OtherMethod,             // line #1403: }
		OtherMethod,             // line #1404:     ''';
		OtherMethod,             // line #1405:     await assertErrorsInCode(code, [
		OtherMethod,             // line #1406:       error(HintCode.UNUSED_LOCAL_VARIABLE, 61, 1),
		OtherMethod,             // line #1407:     ]);
		OtherMethod,             // line #1408:
		OtherMethod,             // line #1409:     var h = (AstFinder.getStatementsInTopLevelFunction(unit, "test")[0]
		OtherMethod,             // line #1410:             as VariableDeclarationStatement)
		OtherMethod,             // line #1411:         .variables
		OtherMethod,             // line #1412:         .variables[0];
		OtherMethod,             // line #1413:     _isDynamic(h.declaredElement.type);
		OtherMethod,             // line #1414:     var fCall = h.initializer as MethodInvocation;
		OtherMethod,             // line #1415:     assertInvokeType(fCall, 'dynamic Function(dynamic Function(dynamic))');
		OtherMethod,             // line #1416:     var g = fCall.argumentList.arguments[0];
		OtherMethod,             // line #1417:     assertType(g.staticType, 'dynamic Function(dynamic)');
		OtherMethod,             // line #1418:   }
		BlankLine,               // line #1419:
		OtherMethod,             // line #1420:   test_inferGenericInstantiation2() async {
		OtherMethod,             // line #1421:     // Verify the behavior when we cannot infer an instantiation due to invalid
		OtherMethod,             // line #1422:     // constraints from an outer generic method.
		OtherMethod,             // line #1423:     var code = r'''
		OtherMethod,             // line #1424: T max<T extends num>(T x, T y) => x < y ? y : x;
		OtherMethod,             // line #1425: abstract class Iterable<T> {
		OtherMethod,             // line #1426:   T get first;
		OtherMethod,             // line #1427:   S fold<S>(S s, S f(S s, T t));
		OtherMethod,             // line #1428: }
		OtherMethod,             // line #1429: num test(Iterable values) => values.fold(values.first as num, max);
		OtherMethod,             // line #1430:     ''';
		OtherMethod,             // line #1431:     await assertErrorsInCode(code, [
		OtherMethod,             // line #1432:       error(CompileTimeErrorCode.COULD_NOT_INFER, 190, 3),
		OtherMethod,             // line #1433:       error(CompileTimeErrorCode.ARGUMENT_TYPE_NOT_ASSIGNABLE, 190, 3),
		OtherMethod,             // line #1434:     ]);
		OtherMethod,             // line #1435:
		OtherMethod,             // line #1436:     var fold = (AstFinder.getTopLevelFunction(unit, 'test')
		OtherMethod,             // line #1437:             .functionExpression
		OtherMethod,             // line #1438:             .body as ExpressionFunctionBody)
		OtherMethod,             // line #1439:         .expression as MethodInvocation;
		OtherMethod,             // line #1440:     assertInvokeType(fold, 'num Function(num, num Function(num, dynamic))');
		OtherMethod,             // line #1441:     var max = fold.argumentList.arguments[1];
		OtherMethod,             // line #1442:     // TODO(jmesserly): arguably num Function(num, num) is better here.
		OtherMethod,             // line #1443:     assertType(max.staticType, 'dynamic Function(dynamic, dynamic)');
		OtherMethod,             // line #1444:   }
		BlankLine,               // line #1445:
		OtherMethod,             // line #1446:   test_inferredFieldDeclaration_propagation() async {
		OtherMethod,             // line #1447:     // Regression test for https://github.com/dart-lang/sdk/issues/25546
		OtherMethod,             // line #1448:     String code = r'''
		OtherMethod,             // line #1449:       abstract class A {
		OtherMethod,             // line #1450:         Map<int, List<int>> get map;
		OtherMethod,             // line #1451:       }
		OtherMethod,             // line #1452:       class B extends A {
		OtherMethod,             // line #1453:         var map = { 42: [] };
		OtherMethod,             // line #1454:       }
		OtherMethod,             // line #1455:       class C extends A {
		OtherMethod,             // line #1456:         get map => { 43: [] };
		OtherMethod,             // line #1457:       }
		OtherMethod,             // line #1458:    ''';
		OtherMethod,             // line #1459:     await assertNoErrorsInCode(code);
		OtherMethod,             // line #1460:
		OtherMethod,             // line #1461:     Asserter<InterfaceType> assertListOfInt = _isListOf(_isInt);
		OtherMethod,             // line #1462:     Asserter<InterfaceType> assertMapOfIntToListOfInt =
		OtherMethod,             // line #1463:         _isMapOf(_isInt, (DartType type) => assertListOfInt(type));
		OtherMethod,             // line #1464:
		OtherMethod,             // line #1465:     VariableDeclaration mapB = AstFinder.getFieldInClass(unit, "B", "map");
		OtherMethod,             // line #1466:     MethodDeclaration mapC = AstFinder.getMethodInClass(unit, "C", "map");
		OtherMethod,             // line #1467:     assertMapOfIntToListOfInt(mapB.declaredElement.type);
		OtherMethod,             // line #1468:     assertMapOfIntToListOfInt(mapC.declaredElement.returnType);
		OtherMethod,             // line #1469:
		OtherMethod,             // line #1470:     SetOrMapLiteral mapLiteralB = mapB.initializer;
		OtherMethod,             // line #1471:     SetOrMapLiteral mapLiteralC =
		OtherMethod,             // line #1472:         (mapC.body as ExpressionFunctionBody).expression;
		OtherMethod,             // line #1473:     assertMapOfIntToListOfInt(mapLiteralB.staticType);
		OtherMethod,             // line #1474:     assertMapOfIntToListOfInt(mapLiteralC.staticType);
		OtherMethod,             // line #1475:
		OtherMethod,             // line #1476:     ListLiteral listLiteralB =
		OtherMethod,             // line #1477:         (mapLiteralB.elements[0] as MapLiteralEntry).value;
		OtherMethod,             // line #1478:     ListLiteral listLiteralC =
		OtherMethod,             // line #1479:         (mapLiteralC.elements[0] as MapLiteralEntry).value;
		OtherMethod,             // line #1480:     assertListOfInt(listLiteralB.staticType);
		OtherMethod,             // line #1481:     assertListOfInt(listLiteralC.staticType);
		OtherMethod,             // line #1482:   }
		BlankLine,               // line #1483:
		OtherMethod,             // line #1484:   test_instanceCreation() async {
		OtherMethod,             // line #1485:     String code = r'''
		OtherMethod,             // line #1486:       class A<S, T> {
		OtherMethod,             // line #1487:         S x;
		OtherMethod,             // line #1488:         T y;
		OtherMethod,             // line #1489:         A(this.x, this.y);
		OtherMethod,             // line #1490:         A.named(this.x, this.y);
		OtherMethod,             // line #1491:       }
		OtherMethod,             // line #1492:
		OtherMethod,             // line #1493:       class B<S, T> extends A<T, S> {
		OtherMethod,             // line #1494:         B(S y, T x) : super(x, y);
		OtherMethod,             // line #1495:         B.named(S y, T x) : super.named(x, y);
		OtherMethod,             // line #1496:       }
		OtherMethod,             // line #1497:
		OtherMethod,             // line #1498:       class C<S> extends B<S, S> {
		OtherMethod,             // line #1499:         C(S a) : super(a, a);
		OtherMethod,             // line #1500:         C.named(S a) : super.named(a, a);
		OtherMethod,             // line #1501:       }
		OtherMethod,             // line #1502:
		OtherMethod,             // line #1503:       class D<S, T> extends B<T, int> {
		OtherMethod,             // line #1504:         D(T a) : super(a, 3);
		OtherMethod,             // line #1505:         D.named(T a) : super.named(a, 3);
		OtherMethod,             // line #1506:       }
		OtherMethod,             // line #1507:
		OtherMethod,             // line #1508:       class E<S, T> extends A<C<S>, T> {
		OtherMethod,             // line #1509:         E(T a) : super(null, a);
		OtherMethod,             // line #1510:       }
		OtherMethod,             // line #1511:
		OtherMethod,             // line #1512:       class F<S, T> extends A<S, T> {
		OtherMethod,             // line #1513:         F(S x, T y, {List<S> a, List<T> b}) : super(x, y);
		OtherMethod,             // line #1514:         F.named(S x, T y, [S a, T b]) : super(a, b);
		OtherMethod,             // line #1515:       }
		OtherMethod,             // line #1516:
		OtherMethod,             // line #1517:       void test0() {
		OtherMethod,             // line #1518:         A<int, String> a0 = new A(3, "hello");
		OtherMethod,             // line #1519:         A<int, String> a1 = new A.named(3, "hello");
		OtherMethod,             // line #1520:         A<int, String> a2 = new A<int, String>(3, "hello");
		OtherMethod,             // line #1521:         A<int, String> a3 = new A<int, String>.named(3, "hello");
		OtherMethod,             // line #1522:         A<int, String> a4 = new A<int, dynamic>(3, "hello");
		OtherMethod,             // line #1523:         A<int, String> a5 = new A<dynamic, dynamic>.named(3, "hello");
		OtherMethod,             // line #1524:       }
		OtherMethod,             // line #1525:       void test1()  {
		OtherMethod,             // line #1526:         A<int, String> a0 = new A("hello", 3);
		OtherMethod,             // line #1527:         A<int, String> a1 = new A.named("hello", 3);
		OtherMethod,             // line #1528:       }
		OtherMethod,             // line #1529:       void test2() {
		OtherMethod,             // line #1530:         A<int, String> a0 = new B("hello", 3);
		OtherMethod,             // line #1531:         A<int, String> a1 = new B.named("hello", 3);
		OtherMethod,             // line #1532:         A<int, String> a2 = new B<String, int>("hello", 3);
		OtherMethod,             // line #1533:         A<int, String> a3 = new B<String, int>.named("hello", 3);
		OtherMethod,             // line #1534:         A<int, String> a4 = new B<String, dynamic>("hello", 3);
		OtherMethod,             // line #1535:         A<int, String> a5 = new B<dynamic, dynamic>.named("hello", 3);
		OtherMethod,             // line #1536:       }
		OtherMethod,             // line #1537:       void test3() {
		OtherMethod,             // line #1538:         A<int, String> a0 = new B(3, "hello");
		OtherMethod,             // line #1539:         A<int, String> a1 = new B.named(3, "hello");
		OtherMethod,             // line #1540:       }
		OtherMethod,             // line #1541:       void test4() {
		OtherMethod,             // line #1542:         A<int, int> a0 = new C(3);
		OtherMethod,             // line #1543:         A<int, int> a1 = new C.named(3);
		OtherMethod,             // line #1544:         A<int, int> a2 = new C<int>(3);
		OtherMethod,             // line #1545:         A<int, int> a3 = new C<int>.named(3);
		OtherMethod,             // line #1546:         A<int, int> a4 = new C<dynamic>(3);
		OtherMethod,             // line #1547:         A<int, int> a5 = new C<dynamic>.named(3);
		OtherMethod,             // line #1548:       }
		OtherMethod,             // line #1549:       void test5() {
		OtherMethod,             // line #1550:         A<int, int> a0 = new C("hello");
		OtherMethod,             // line #1551:         A<int, int> a1 = new C.named("hello");
		OtherMethod,             // line #1552:       }
		OtherMethod,             // line #1553:       void test6()  {
		OtherMethod,             // line #1554:         A<int, String> a0 = new D("hello");
		OtherMethod,             // line #1555:         A<int, String> a1 = new D.named("hello");
		OtherMethod,             // line #1556:         A<int, String> a2 = new D<int, String>("hello");
		OtherMethod,             // line #1557:         A<int, String> a3 = new D<String, String>.named("hello");
		OtherMethod,             // line #1558:         A<int, String> a4 = new D<num, dynamic>("hello");
		OtherMethod,             // line #1559:         A<int, String> a5 = new D<dynamic, dynamic>.named("hello");
		OtherMethod,             // line #1560:       }
		OtherMethod,             // line #1561:       void test7() {
		OtherMethod,             // line #1562:         A<int, String> a0 = new D(3);
		OtherMethod,             // line #1563:         A<int, String> a1 = new D.named(3);
		OtherMethod,             // line #1564:       }
		OtherMethod,             // line #1565:       void test8() {
		OtherMethod,             // line #1566:         A<C<int>, String> a0 = new E("hello");
		OtherMethod,             // line #1567:       }
		OtherMethod,             // line #1568:       void test9() { // Check named and optional arguments
		OtherMethod,             // line #1569:         A<int, String> a0 = new F(3, "hello", a: [3], b: ["hello"]);
		OtherMethod,             // line #1570:         A<int, String> a1 = new F(3, "hello", a: ["hello"], b:[3]);
		OtherMethod,             // line #1571:         A<int, String> a2 = new F.named(3, "hello", 3, "hello");
		OtherMethod,             // line #1572:         A<int, String> a3 = new F.named(3, "hello");
		OtherMethod,             // line #1573:         A<int, String> a4 = new F.named(3, "hello", "hello", 3);
		OtherMethod,             // line #1574:         A<int, String> a5 = new F.named(3, "hello", "hello");
		OtherMethod,             // line #1575:       }''';
		OtherMethod,             // line #1576:     await assertErrorsInCode(code, [
		OtherMethod,             // line #1577:       error(HintCode.UNUSED_LOCAL_VARIABLE, 769, 2),
		OtherMethod,             // line #1578:       error(HintCode.UNUSED_LOCAL_VARIABLE, 816, 2),
		OtherMethod,             // line #1579:       error(HintCode.UNUSED_LOCAL_VARIABLE, 869, 2),
		OtherMethod,             // line #1580:       error(HintCode.UNUSED_LOCAL_VARIABLE, 929, 2),
		OtherMethod,             // line #1581:       error(HintCode.UNUSED_LOCAL_VARIABLE, 995, 2),
		OtherMethod,             // line #1582:       error(CompileTimeErrorCode.INVALID_CAST_NEW_EXPR, 1000, 31),
		OtherMethod,             // line #1583:       error(HintCode.UNUSED_LOCAL_VARIABLE, 1056, 2),
		OtherMethod,             // line #1584:       error(CompileTimeErrorCode.INVALID_CAST_NEW_EXPR, 1061, 41),
		OtherMethod,             // line #1585:       error(HintCode.UNUSED_LOCAL_VARIABLE, 1157, 2),
		OtherMethod,             // line #1586:       error(CompileTimeErrorCode.ARGUMENT_TYPE_NOT_ASSIGNABLE, 1168, 7),
		OtherMethod,             // line #1587:       error(CompileTimeErrorCode.ARGUMENT_TYPE_NOT_ASSIGNABLE, 1177, 1),
		OtherMethod,             // line #1588:       error(HintCode.UNUSED_LOCAL_VARIABLE, 1204, 2),
		OtherMethod,             // line #1589:       error(CompileTimeErrorCode.ARGUMENT_TYPE_NOT_ASSIGNABLE, 1221, 7),
		OtherMethod,             // line #1590:       error(CompileTimeErrorCode.ARGUMENT_TYPE_NOT_ASSIGNABLE, 1230, 1),
		OtherMethod,             // line #1591:       error(HintCode.UNUSED_LOCAL_VARIABLE, 1286, 2),
		OtherMethod,             // line #1592:       error(HintCode.UNUSED_LOCAL_VARIABLE, 1333, 2),
		OtherMethod,             // line #1593:       error(HintCode.UNUSED_LOCAL_VARIABLE, 1386, 2),
		OtherMethod,             // line #1594:       error(HintCode.UNUSED_LOCAL_VARIABLE, 1446, 2),
		OtherMethod,             // line #1595:       error(HintCode.UNUSED_LOCAL_VARIABLE, 1512, 2),
		OtherMethod,             // line #1596:       error(CompileTimeErrorCode.INVALID_ASSIGNMENT, 1517, 34),
		OtherMethod,             // line #1597:       error(HintCode.UNUSED_LOCAL_VARIABLE, 1576, 2),
		OtherMethod,             // line #1598:       error(CompileTimeErrorCode.INVALID_ASSIGNMENT, 1581, 41),
		OtherMethod,             // line #1599:       error(HintCode.UNUSED_LOCAL_VARIABLE, 1676, 2),
		OtherMethod,             // line #1600:       error(CompileTimeErrorCode.ARGUMENT_TYPE_NOT_ASSIGNABLE, 1687, 1),
		OtherMethod,             // line #1601:       error(CompileTimeErrorCode.ARGUMENT_TYPE_NOT_ASSIGNABLE, 1690, 7),
		OtherMethod,             // line #1602:       error(HintCode.UNUSED_LOCAL_VARIABLE, 1723, 2),
		OtherMethod,             // line #1603:       error(CompileTimeErrorCode.ARGUMENT_TYPE_NOT_ASSIGNABLE, 1740, 1),
		OtherMethod,             // line #1604:       error(CompileTimeErrorCode.ARGUMENT_TYPE_NOT_ASSIGNABLE, 1743, 7),
		OtherMethod,             // line #1605:       error(HintCode.UNUSED_LOCAL_VARIABLE, 1802, 2),
		OtherMethod,             // line #1606:       error(HintCode.UNUSED_LOCAL_VARIABLE, 1837, 2),
		OtherMethod,             // line #1607:       error(HintCode.UNUSED_LOCAL_VARIABLE, 1878, 2),
		OtherMethod,             // line #1608:       error(HintCode.UNUSED_LOCAL_VARIABLE, 1918, 2),
		OtherMethod,             // line #1609:       error(HintCode.UNUSED_LOCAL_VARIABLE, 1964, 2),
		OtherMethod,             // line #1610:       error(CompileTimeErrorCode.INVALID_ASSIGNMENT, 1969, 17),
		OtherMethod,             // line #1611:       error(HintCode.UNUSED_LOCAL_VARIABLE, 2008, 2),
		OtherMethod,             // line #1612:       error(CompileTimeErrorCode.INVALID_ASSIGNMENT, 2013, 23),
		OtherMethod,             // line #1613:       error(HintCode.UNUSED_LOCAL_VARIABLE, 2087, 2),
		OtherMethod,             // line #1614:       error(CompileTimeErrorCode.ARGUMENT_TYPE_NOT_ASSIGNABLE, 2098, 7),
		OtherMethod,             // line #1615:       error(HintCode.UNUSED_LOCAL_VARIABLE, 2128, 2),
		OtherMethod,             // line #1616:       error(CompileTimeErrorCode.ARGUMENT_TYPE_NOT_ASSIGNABLE, 2145, 7),
		OtherMethod,             // line #1617:       error(HintCode.UNUSED_LOCAL_VARIABLE, 2208, 2),
		OtherMethod,             // line #1618:       error(HintCode.UNUSED_LOCAL_VARIABLE, 2252, 2),
		OtherMethod,             // line #1619:       error(HintCode.UNUSED_LOCAL_VARIABLE, 2302, 2),
		OtherMethod,             // line #1620:       error(HintCode.UNUSED_LOCAL_VARIABLE, 2359, 2),
		OtherMethod,             // line #1621:       error(HintCode.UNUSED_LOCAL_VARIABLE, 2425, 2),
		OtherMethod,             // line #1622:       error(CompileTimeErrorCode.INVALID_ASSIGNMENT, 2430, 28),
		OtherMethod,             // line #1623:       error(HintCode.UNUSED_LOCAL_VARIABLE, 2483, 2),
		OtherMethod,             // line #1624:       error(CompileTimeErrorCode.INVALID_ASSIGNMENT, 2488, 38),
		OtherMethod,             // line #1625:       error(HintCode.UNUSED_LOCAL_VARIABLE, 2580, 2),
		OtherMethod,             // line #1626:       error(CompileTimeErrorCode.ARGUMENT_TYPE_NOT_ASSIGNABLE, 2591, 1),
		OtherMethod,             // line #1627:       error(HintCode.UNUSED_LOCAL_VARIABLE, 2618, 2),
		OtherMethod,             // line #1628:       error(CompileTimeErrorCode.ARGUMENT_TYPE_NOT_ASSIGNABLE, 2635, 1),
		OtherMethod,             // line #1629:       error(HintCode.UNUSED_LOCAL_VARIABLE, 2694, 2),
		OtherMethod,             // line #1630:       error(HintCode.UNUSED_LOCAL_VARIABLE, 2805, 2),
		OtherMethod,             // line #1631:       error(HintCode.UNUSED_LOCAL_VARIABLE, 2874, 2),
		OtherMethod,             // line #1632:       error(CompileTimeErrorCode.LIST_ELEMENT_TYPE_NOT_ASSIGNABLE, 2901, 7),
		OtherMethod,             // line #1633:       error(CompileTimeErrorCode.LIST_ELEMENT_TYPE_NOT_ASSIGNABLE, 2914, 1),
		OtherMethod,             // line #1634:       error(HintCode.UNUSED_LOCAL_VARIABLE, 2942, 2),
		OtherMethod,             // line #1635:       error(HintCode.UNUSED_LOCAL_VARIABLE, 3007, 2),
		OtherMethod,             // line #1636:       error(HintCode.UNUSED_LOCAL_VARIABLE, 3060, 2),
		OtherMethod,             // line #1637:       error(CompileTimeErrorCode.ARGUMENT_TYPE_NOT_ASSIGNABLE, 3089, 7),
		OtherMethod,             // line #1638:       error(CompileTimeErrorCode.ARGUMENT_TYPE_NOT_ASSIGNABLE, 3098, 1),
		OtherMethod,             // line #1639:       error(HintCode.UNUSED_LOCAL_VARIABLE, 3125, 2),
		OtherMethod,             // line #1640:       error(CompileTimeErrorCode.ARGUMENT_TYPE_NOT_ASSIGNABLE, 3154, 7),
		OtherMethod,             // line #1641:     ]);
		OtherMethod,             // line #1642:
		OtherMethod,             // line #1643:     Expression rhs(VariableDeclarationStatement stmt) {
		OtherMethod,             // line #1644:       VariableDeclaration decl = stmt.variables.variables[0];
		OtherMethod,             // line #1645:       Expression exp = decl.initializer;
		OtherMethod,             // line #1646:       return exp;
		OtherMethod,             // line #1647:     }
		OtherMethod,             // line #1648:
		OtherMethod,             // line #1649:     void hasType(Asserter<DartType> assertion, Expression exp) =>
		OtherMethod,             // line #1650:         assertion(exp.staticType);
		OtherMethod,             // line #1651:
		OtherMethod,             // line #1652:     Element elementA = AstFinder.getClass(unit, "A").declaredElement;
		OtherMethod,             // line #1653:     Element elementB = AstFinder.getClass(unit, "B").declaredElement;
		OtherMethod,             // line #1654:     Element elementC = AstFinder.getClass(unit, "C").declaredElement;
		OtherMethod,             // line #1655:     Element elementD = AstFinder.getClass(unit, "D").declaredElement;
		OtherMethod,             // line #1656:     Element elementE = AstFinder.getClass(unit, "E").declaredElement;
		OtherMethod,             // line #1657:     Element elementF = AstFinder.getClass(unit, "F").declaredElement;
		OtherMethod,             // line #1658:
		OtherMethod,             // line #1659:     AsserterBuilder<List<Asserter<DartType>>, DartType> assertAOf =
		OtherMethod,             // line #1660:         _isInstantiationOf(_hasElement(elementA));
		OtherMethod,             // line #1661:     AsserterBuilder<List<Asserter<DartType>>, DartType> assertBOf =
		OtherMethod,             // line #1662:         _isInstantiationOf(_hasElement(elementB));
		OtherMethod,             // line #1663:     AsserterBuilder<List<Asserter<DartType>>, DartType> assertCOf =
		OtherMethod,             // line #1664:         _isInstantiationOf(_hasElement(elementC));
		OtherMethod,             // line #1665:     AsserterBuilder<List<Asserter<DartType>>, DartType> assertDOf =
		OtherMethod,             // line #1666:         _isInstantiationOf(_hasElement(elementD));
		OtherMethod,             // line #1667:     AsserterBuilder<List<Asserter<DartType>>, DartType> assertEOf =
		OtherMethod,             // line #1668:         _isInstantiationOf(_hasElement(elementE));
		OtherMethod,             // line #1669:     AsserterBuilder<List<Asserter<DartType>>, DartType> assertFOf =
		OtherMethod,             // line #1670:         _isInstantiationOf(_hasElement(elementF));
		OtherMethod,             // line #1671:
		OtherMethod,             // line #1672:     {
		OtherMethod,             // line #1673:       List<Statement> statements =
		OtherMethod,             // line #1674:           AstFinder.getStatementsInTopLevelFunction(unit, "test0");
		OtherMethod,             // line #1675:
		OtherMethod,             // line #1676:       hasType(assertAOf([_isInt, _isString]), rhs(statements[0]));
		OtherMethod,             // line #1677:       hasType(assertAOf([_isInt, _isString]), rhs(statements[0]));
		OtherMethod,             // line #1678:       hasType(assertAOf([_isInt, _isString]), rhs(statements[1]));
		OtherMethod,             // line #1679:       hasType(assertAOf([_isInt, _isString]), rhs(statements[2]));
		OtherMethod,             // line #1680:       hasType(assertAOf([_isInt, _isString]), rhs(statements[3]));
		OtherMethod,             // line #1681:       hasType(assertAOf([_isInt, _isDynamic]), rhs(statements[4]));
		OtherMethod,             // line #1682:       hasType(assertAOf([_isDynamic, _isDynamic]), rhs(statements[5]));
		OtherMethod,             // line #1683:     }
		OtherMethod,             // line #1684:
		OtherMethod,             // line #1685:     {
		OtherMethod,             // line #1686:       List<Statement> statements =
		OtherMethod,             // line #1687:           AstFinder.getStatementsInTopLevelFunction(unit, "test1");
		OtherMethod,             // line #1688:       hasType(assertAOf([_isInt, _isString]), rhs(statements[0]));
		OtherMethod,             // line #1689:       hasType(assertAOf([_isInt, _isString]), rhs(statements[1]));
		OtherMethod,             // line #1690:     }
		OtherMethod,             // line #1691:
		OtherMethod,             // line #1692:     {
		OtherMethod,             // line #1693:       List<Statement> statements =
		OtherMethod,             // line #1694:           AstFinder.getStatementsInTopLevelFunction(unit, "test2");
		OtherMethod,             // line #1695:       hasType(assertBOf([_isString, _isInt]), rhs(statements[0]));
		OtherMethod,             // line #1696:       hasType(assertBOf([_isString, _isInt]), rhs(statements[1]));
		OtherMethod,             // line #1697:       hasType(assertBOf([_isString, _isInt]), rhs(statements[2]));
		OtherMethod,             // line #1698:       hasType(assertBOf([_isString, _isInt]), rhs(statements[3]));
		OtherMethod,             // line #1699:       hasType(assertBOf([_isString, _isDynamic]), rhs(statements[4]));
		OtherMethod,             // line #1700:       hasType(assertBOf([_isDynamic, _isDynamic]), rhs(statements[5]));
		OtherMethod,             // line #1701:     }
		OtherMethod,             // line #1702:
		OtherMethod,             // line #1703:     {
		OtherMethod,             // line #1704:       List<Statement> statements =
		OtherMethod,             // line #1705:           AstFinder.getStatementsInTopLevelFunction(unit, "test3");
		OtherMethod,             // line #1706:       hasType(assertBOf([_isString, _isInt]), rhs(statements[0]));
		OtherMethod,             // line #1707:       hasType(assertBOf([_isString, _isInt]), rhs(statements[1]));
		OtherMethod,             // line #1708:     }
		OtherMethod,             // line #1709:
		OtherMethod,             // line #1710:     {
		OtherMethod,             // line #1711:       List<Statement> statements =
		OtherMethod,             // line #1712:           AstFinder.getStatementsInTopLevelFunction(unit, "test4");
		OtherMethod,             // line #1713:       hasType(assertCOf([_isInt]), rhs(statements[0]));
		OtherMethod,             // line #1714:       hasType(assertCOf([_isInt]), rhs(statements[1]));
		OtherMethod,             // line #1715:       hasType(assertCOf([_isInt]), rhs(statements[2]));
		OtherMethod,             // line #1716:       hasType(assertCOf([_isInt]), rhs(statements[3]));
		OtherMethod,             // line #1717:       hasType(assertCOf([_isDynamic]), rhs(statements[4]));
		OtherMethod,             // line #1718:       hasType(assertCOf([_isDynamic]), rhs(statements[5]));
		OtherMethod,             // line #1719:     }
		OtherMethod,             // line #1720:
		OtherMethod,             // line #1721:     {
		OtherMethod,             // line #1722:       List<Statement> statements =
		OtherMethod,             // line #1723:           AstFinder.getStatementsInTopLevelFunction(unit, "test5");
		OtherMethod,             // line #1724:       hasType(assertCOf([_isInt]), rhs(statements[0]));
		OtherMethod,             // line #1725:       hasType(assertCOf([_isInt]), rhs(statements[1]));
		OtherMethod,             // line #1726:     }
		OtherMethod,             // line #1727:
		OtherMethod,             // line #1728:     {
		OtherMethod,             // line #1729:       // The first type parameter is not constrained by the
		OtherMethod,             // line #1730:       // context.  We could choose a tighter type, but currently
		OtherMethod,             // line #1731:       // we just use dynamic.
		OtherMethod,             // line #1732:       List<Statement> statements =
		OtherMethod,             // line #1733:           AstFinder.getStatementsInTopLevelFunction(unit, "test6");
		OtherMethod,             // line #1734:       hasType(assertDOf([_isDynamic, _isString]), rhs(statements[0]));
		OtherMethod,             // line #1735:       hasType(assertDOf([_isDynamic, _isString]), rhs(statements[1]));
		OtherMethod,             // line #1736:       hasType(assertDOf([_isInt, _isString]), rhs(statements[2]));
		OtherMethod,             // line #1737:       hasType(assertDOf([_isString, _isString]), rhs(statements[3]));
		OtherMethod,             // line #1738:       hasType(assertDOf([_isNum, _isDynamic]), rhs(statements[4]));
		OtherMethod,             // line #1739:       hasType(assertDOf([_isDynamic, _isDynamic]), rhs(statements[5]));
		OtherMethod,             // line #1740:     }
		OtherMethod,             // line #1741:
		OtherMethod,             // line #1742:     {
		OtherMethod,             // line #1743:       List<Statement> statements =
		OtherMethod,             // line #1744:           AstFinder.getStatementsInTopLevelFunction(unit, "test7");
		OtherMethod,             // line #1745:       hasType(assertDOf([_isDynamic, _isString]), rhs(statements[0]));
		OtherMethod,             // line #1746:       hasType(assertDOf([_isDynamic, _isString]), rhs(statements[1]));
		OtherMethod,             // line #1747:     }
		OtherMethod,             // line #1748:
		OtherMethod,             // line #1749:     {
		OtherMethod,             // line #1750:       List<Statement> statements =
		OtherMethod,             // line #1751:           AstFinder.getStatementsInTopLevelFunction(unit, "test8");
		OtherMethod,             // line #1752:       hasType(assertEOf([_isInt, _isString]), rhs(statements[0]));
		OtherMethod,             // line #1753:     }
		OtherMethod,             // line #1754:
		OtherMethod,             // line #1755:     {
		OtherMethod,             // line #1756:       List<Statement> statements =
		OtherMethod,             // line #1757:           AstFinder.getStatementsInTopLevelFunction(unit, "test9");
		OtherMethod,             // line #1758:       hasType(assertFOf([_isInt, _isString]), rhs(statements[0]));
		OtherMethod,             // line #1759:       hasType(assertFOf([_isInt, _isString]), rhs(statements[1]));
		OtherMethod,             // line #1760:       hasType(assertFOf([_isInt, _isString]), rhs(statements[2]));
		OtherMethod,             // line #1761:       hasType(assertFOf([_isInt, _isString]), rhs(statements[3]));
		OtherMethod,             // line #1762:       hasType(assertFOf([_isInt, _isString]), rhs(statements[4]));
		OtherMethod,             // line #1763:       hasType(assertFOf([_isInt, _isString]), rhs(statements[5]));
		OtherMethod,             // line #1764:     }
		OtherMethod,             // line #1765:   }
		BlankLine,               // line #1766:
		OtherMethod,             // line #1767:   test_listLiteral_nested() async {
		OtherMethod,             // line #1768:     String code = r'''
		OtherMethod,             // line #1769:       void main () {
		OtherMethod,             // line #1770:         List<List<int>> l0 = [[]];
		OtherMethod,             // line #1771:         Iterable<List<int>> l1 = [[3]];
		OtherMethod,             // line #1772:         Iterable<List<int>> l2 = [[3], [4]];
		OtherMethod,             // line #1773:         List<List<int>> l3 = [["hello", 3], []];
		OtherMethod,             // line #1774:      }
		OtherMethod,             // line #1775:    ''';
		OtherMethod,             // line #1776:     await assertErrorsInCode(code, [
		OtherMethod,             // line #1777:       error(HintCode.UNUSED_LOCAL_VARIABLE, 45, 2),
		OtherMethod,             // line #1778:       error(HintCode.UNUSED_LOCAL_VARIABLE, 84, 2),
		OtherMethod,             // line #1779:       error(HintCode.UNUSED_LOCAL_VARIABLE, 124, 2),
		OtherMethod,             // line #1780:       error(HintCode.UNUSED_LOCAL_VARIABLE, 165, 2),
		OtherMethod,             // line #1781:       error(CompileTimeErrorCode.LIST_ELEMENT_TYPE_NOT_ASSIGNABLE, 172, 7),
		OtherMethod,             // line #1782:     ]);
		OtherMethod,             // line #1783:
		OtherMethod,             // line #1784:     List<Statement> statements =
		OtherMethod,             // line #1785:         AstFinder.getStatementsInTopLevelFunction(unit, "main");
		OtherMethod,             // line #1786:     ListLiteral literal(int i) {
		OtherMethod,             // line #1787:       VariableDeclarationStatement stmt = statements[i];
		OtherMethod,             // line #1788:       VariableDeclaration decl = stmt.variables.variables[0];
		OtherMethod,             // line #1789:       ListLiteral exp = decl.initializer;
		OtherMethod,             // line #1790:       return exp;
		OtherMethod,             // line #1791:     }
		OtherMethod,             // line #1792:
		OtherMethod,             // line #1793:     Asserter<InterfaceType> assertListOfInt = _isListOf(_isInt);
		OtherMethod,             // line #1794:     Asserter<InterfaceType> assertListOfListOfInt =
		OtherMethod,             // line #1795:         _isListOf((DartType type) => assertListOfInt(type));
		OtherMethod,             // line #1796:
		OtherMethod,             // line #1797:     assertListOfListOfInt(literal(0).staticType);
		OtherMethod,             // line #1798:     assertListOfListOfInt(literal(1).staticType);
		OtherMethod,             // line #1799:     assertListOfListOfInt(literal(2).staticType);
		OtherMethod,             // line #1800:     assertListOfListOfInt(literal(3).staticType);
		OtherMethod,             // line #1801:
		OtherMethod,             // line #1802:     assertListOfInt((literal(1).elements[0] as Expression).staticType);
		OtherMethod,             // line #1803:     assertListOfInt((literal(2).elements[0] as Expression).staticType);
		OtherMethod,             // line #1804:     assertListOfInt((literal(3).elements[0] as Expression).staticType);
		OtherMethod,             // line #1805:   }
		BlankLine,               // line #1806:
		OtherMethod,             // line #1807:   test_listLiteral_simple() async {
		OtherMethod,             // line #1808:     String code = r'''
		OtherMethod,             // line #1809:       void main () {
		OtherMethod,             // line #1810:         List<int> l0 = [];
		OtherMethod,             // line #1811:         List<int> l1 = [3];
		OtherMethod,             // line #1812:         List<int> l2 = ["hello"];
		OtherMethod,             // line #1813:         List<int> l3 = ["hello", 3];
		OtherMethod,             // line #1814:      }
		OtherMethod,             // line #1815:    ''';
		OtherMethod,             // line #1816:     await assertErrorsInCode(code, [
		OtherMethod,             // line #1817:       error(HintCode.UNUSED_LOCAL_VARIABLE, 39, 2),
		OtherMethod,             // line #1818:       error(HintCode.UNUSED_LOCAL_VARIABLE, 66, 2),
		OtherMethod,             // line #1819:       error(HintCode.UNUSED_LOCAL_VARIABLE, 94, 2),
		OtherMethod,             // line #1820:       error(CompileTimeErrorCode.LIST_ELEMENT_TYPE_NOT_ASSIGNABLE, 100, 7),
		OtherMethod,             // line #1821:       error(HintCode.UNUSED_LOCAL_VARIABLE, 128, 2),
		OtherMethod,             // line #1822:       error(CompileTimeErrorCode.LIST_ELEMENT_TYPE_NOT_ASSIGNABLE, 134, 7),
		OtherMethod,             // line #1823:     ]);
		OtherMethod,             // line #1824:
		OtherMethod,             // line #1825:     List<Statement> statements =
		OtherMethod,             // line #1826:         AstFinder.getStatementsInTopLevelFunction(unit, "main");
		OtherMethod,             // line #1827:     DartType literal(int i) {
		OtherMethod,             // line #1828:       VariableDeclarationStatement stmt = statements[i];
		OtherMethod,             // line #1829:       VariableDeclaration decl = stmt.variables.variables[0];
		OtherMethod,             // line #1830:       ListLiteral exp = decl.initializer;
		OtherMethod,             // line #1831:       return exp.staticType;
		OtherMethod,             // line #1832:     }
		OtherMethod,             // line #1833:
		OtherMethod,             // line #1834:     Asserter<InterfaceType> assertListOfInt = _isListOf(_isInt);
		OtherMethod,             // line #1835:
		OtherMethod,             // line #1836:     assertListOfInt(literal(0));
		OtherMethod,             // line #1837:     assertListOfInt(literal(1));
		OtherMethod,             // line #1838:     assertListOfInt(literal(2));
		OtherMethod,             // line #1839:     assertListOfInt(literal(3));
		OtherMethod,             // line #1840:   }
		BlankLine,               // line #1841:
		OtherMethod,             // line #1842:   test_listLiteral_simple_const() async {
		OtherMethod,             // line #1843:     String code = r'''
		OtherMethod,             // line #1844:       void main () {
		OtherMethod,             // line #1845:         const List<int> c0 = const [];
		OtherMethod,             // line #1846:         const List<int> c1 = const [3];
		OtherMethod,             // line #1847:         const List<int> c2 = const ["hello"];
		OtherMethod,             // line #1848:         const List<int> c3 = const ["hello", 3];
		OtherMethod,             // line #1849:      }
		OtherMethod,             // line #1850:    ''';
		OtherMethod,             // line #1851:     await assertErrorsInCode(code, [
		OtherMethod,             // line #1852:       error(HintCode.UNUSED_LOCAL_VARIABLE, 45, 2),
		OtherMethod,             // line #1853:       error(HintCode.UNUSED_LOCAL_VARIABLE, 84, 2),
		OtherMethod,             // line #1854:       error(HintCode.UNUSED_LOCAL_VARIABLE, 124, 2),
		OtherMethod,             // line #1855:       error(CompileTimeErrorCode.LIST_ELEMENT_TYPE_NOT_ASSIGNABLE, 136, 7),
		OtherMethod,             // line #1856:       error(HintCode.UNUSED_LOCAL_VARIABLE, 170, 2),
		OtherMethod,             // line #1857:       error(CompileTimeErrorCode.LIST_ELEMENT_TYPE_NOT_ASSIGNABLE, 182, 7),
		OtherMethod,             // line #1858:     ]);
		OtherMethod,             // line #1859:
		OtherMethod,             // line #1860:     List<Statement> statements =
		OtherMethod,             // line #1861:         AstFinder.getStatementsInTopLevelFunction(unit, "main");
		OtherMethod,             // line #1862:     DartType literal(int i) {
		OtherMethod,             // line #1863:       VariableDeclarationStatement stmt = statements[i];
		OtherMethod,             // line #1864:       VariableDeclaration decl = stmt.variables.variables[0];
		OtherMethod,             // line #1865:       ListLiteral exp = decl.initializer;
		OtherMethod,             // line #1866:       return exp.staticType;
		OtherMethod,             // line #1867:     }
		OtherMethod,             // line #1868:
		OtherMethod,             // line #1869:     Asserter<InterfaceType> assertListOfInt = _isListOf(_isInt);
		OtherMethod,             // line #1870:
		OtherMethod,             // line #1871:     assertListOfInt(literal(0));
		OtherMethod,             // line #1872:     assertListOfInt(literal(1));
		OtherMethod,             // line #1873:     assertListOfInt(literal(2));
		OtherMethod,             // line #1874:     assertListOfInt(literal(3));
		OtherMethod,             // line #1875:   }
		BlankLine,               // line #1876:
		OtherMethod,             // line #1877:   test_listLiteral_simple_disabled() async {
		OtherMethod,             // line #1878:     String code = r'''
		OtherMethod,             // line #1879:       void main () {
		OtherMethod,             // line #1880:         List<int> l0 = <num>[];
		OtherMethod,             // line #1881:         List<int> l1 = <num>[3];
		OtherMethod,             // line #1882:         List<int> l2 = <String>["hello"];
		OtherMethod,             // line #1883:         List<int> l3 = <dynamic>["hello", 3];
		OtherMethod,             // line #1884:      }
		OtherMethod,             // line #1885:    ''';
		OtherMethod,             // line #1886:     await assertErrorsInCode(code, [
		OtherMethod,             // line #1887:       error(HintCode.UNUSED_LOCAL_VARIABLE, 39, 2),
		OtherMethod,             // line #1888:       error(CompileTimeErrorCode.INVALID_CAST_LITERAL_LIST, 44, 7),
		OtherMethod,             // line #1889:       error(HintCode.UNUSED_LOCAL_VARIABLE, 71, 2),
		OtherMethod,             // line #1890:       error(CompileTimeErrorCode.INVALID_CAST_LITERAL_LIST, 76, 8),
		OtherMethod,             // line #1891:       error(HintCode.UNUSED_LOCAL_VARIABLE, 104, 2),
		OtherMethod,             // line #1892:       error(CompileTimeErrorCode.INVALID_ASSIGNMENT, 109, 17),
		OtherMethod,             // line #1893:       error(HintCode.UNUSED_LOCAL_VARIABLE, 146, 2),
		OtherMethod,             // line #1894:       error(CompileTimeErrorCode.INVALID_CAST_LITERAL_LIST, 151, 21),
		OtherMethod,             // line #1895:     ]);
		OtherMethod,             // line #1896:
		OtherMethod,             // line #1897:     List<Statement> statements =
		OtherMethod,             // line #1898:         AstFinder.getStatementsInTopLevelFunction(unit, "main");
		OtherMethod,             // line #1899:     DartType literal(int i) {
		OtherMethod,             // line #1900:       VariableDeclarationStatement stmt = statements[i];
		OtherMethod,             // line #1901:       VariableDeclaration decl = stmt.variables.variables[0];
		OtherMethod,             // line #1902:       ListLiteral exp = decl.initializer;
		OtherMethod,             // line #1903:       return exp.staticType;
		OtherMethod,             // line #1904:     }
		OtherMethod,             // line #1905:
		OtherMethod,             // line #1906:     _isListOf(_isNum)(literal(0));
		OtherMethod,             // line #1907:     _isListOf(_isNum)(literal(1));
		OtherMethod,             // line #1908:     _isListOf(_isString)(literal(2));
		OtherMethod,             // line #1909:     _isListOf(_isDynamic)(literal(3));
		OtherMethod,             // line #1910:   }
		BlankLine,               // line #1911:
		OtherMethod,             // line #1912:   test_listLiteral_simple_subtype() async {
		OtherMethod,             // line #1913:     String code = r'''
		OtherMethod,             // line #1914:       void main () {
		OtherMethod,             // line #1915:         Iterable<int> l0 = [];
		OtherMethod,             // line #1916:         Iterable<int> l1 = [3];
		OtherMethod,             // line #1917:         Iterable<int> l2 = ["hello"];
		OtherMethod,             // line #1918:         Iterable<int> l3 = ["hello", 3];
		OtherMethod,             // line #1919:      }
		OtherMethod,             // line #1920:    ''';
		OtherMethod,             // line #1921:     await assertErrorsInCode(code, [
		OtherMethod,             // line #1922:       error(HintCode.UNUSED_LOCAL_VARIABLE, 43, 2),
		OtherMethod,             // line #1923:       error(HintCode.UNUSED_LOCAL_VARIABLE, 74, 2),
		OtherMethod,             // line #1924:       error(HintCode.UNUSED_LOCAL_VARIABLE, 106, 2),
		OtherMethod,             // line #1925:       error(CompileTimeErrorCode.LIST_ELEMENT_TYPE_NOT_ASSIGNABLE, 112, 7),
		OtherMethod,             // line #1926:       error(HintCode.UNUSED_LOCAL_VARIABLE, 144, 2),
		OtherMethod,             // line #1927:       error(CompileTimeErrorCode.LIST_ELEMENT_TYPE_NOT_ASSIGNABLE, 150, 7),
		OtherMethod,             // line #1928:     ]);
		OtherMethod,             // line #1929:
		OtherMethod,             // line #1930:     List<Statement> statements =
		OtherMethod,             // line #1931:         AstFinder.getStatementsInTopLevelFunction(unit, "main");
		OtherMethod,             // line #1932:     DartType literal(int i) {
		OtherMethod,             // line #1933:       VariableDeclarationStatement stmt = statements[i];
		OtherMethod,             // line #1934:       VariableDeclaration decl = stmt.variables.variables[0];
		OtherMethod,             // line #1935:       ListLiteral exp = decl.initializer;
		OtherMethod,             // line #1936:       return exp.staticType;
		OtherMethod,             // line #1937:     }
		OtherMethod,             // line #1938:
		OtherMethod,             // line #1939:     Asserter<InterfaceType> assertListOfInt = _isListOf(_isInt);
		OtherMethod,             // line #1940:
		OtherMethod,             // line #1941:     assertListOfInt(literal(0));
		OtherMethod,             // line #1942:     assertListOfInt(literal(1));
		OtherMethod,             // line #1943:     assertListOfInt(literal(2));
		OtherMethod,             // line #1944:     assertListOfInt(literal(3));
		OtherMethod,             // line #1945:   }
		BlankLine,               // line #1946:
		OtherMethod,             // line #1947:   test_mapLiteral_nested() async {
		OtherMethod,             // line #1948:     String code = r'''
		OtherMethod,             // line #1949:       void main () {
		OtherMethod,             // line #1950:         Map<int, List<String>> l0 = {};
		OtherMethod,             // line #1951:         Map<int, List<String>> l1 = {3: ["hello"]};
		OtherMethod,             // line #1952:         Map<int, List<String>> l2 = {"hello": ["hello"]};
		OtherMethod,             // line #1953:         Map<int, List<String>> l3 = {3: [3]};
		OtherMethod,             // line #1954:         Map<int, List<String>> l4 = {3:["hello"], "hello": [3]};
		OtherMethod,             // line #1955:      }
		OtherMethod,             // line #1956:    ''';
		OtherMethod,             // line #1957:     await assertErrorsInCode(code, [
		OtherMethod,             // line #1958:       error(HintCode.UNUSED_LOCAL_VARIABLE, 52, 2),
		OtherMethod,             // line #1959:       error(HintCode.UNUSED_LOCAL_VARIABLE, 92, 2),
		OtherMethod,             // line #1960:       error(HintCode.UNUSED_LOCAL_VARIABLE, 144, 2),
		OtherMethod,             // line #1961:       error(CompileTimeErrorCode.MAP_KEY_TYPE_NOT_ASSIGNABLE, 150, 7),
		OtherMethod,             // line #1962:       error(HintCode.UNUSED_LOCAL_VARIABLE, 202, 2),
		OtherMethod,             // line #1963:       error(CompileTimeErrorCode.LIST_ELEMENT_TYPE_NOT_ASSIGNABLE, 212, 1),
		OtherMethod,             // line #1964:       error(HintCode.UNUSED_LOCAL_VARIABLE, 248, 2),
		OtherMethod,             // line #1965:       error(CompileTimeErrorCode.MAP_KEY_TYPE_NOT_ASSIGNABLE, 267, 7),
		OtherMethod,             // line #1966:       error(CompileTimeErrorCode.LIST_ELEMENT_TYPE_NOT_ASSIGNABLE, 277, 1),
		OtherMethod,             // line #1967:     ]);
		OtherMethod,             // line #1968:
		OtherMethod,             // line #1969:     List<Statement> statements =
		OtherMethod,             // line #1970:         AstFinder.getStatementsInTopLevelFunction(unit, "main");
		OtherMethod,             // line #1971:     SetOrMapLiteral literal(int i) {
		OtherMethod,             // line #1972:       VariableDeclarationStatement stmt = statements[i];
		OtherMethod,             // line #1973:       VariableDeclaration decl = stmt.variables.variables[0];
		OtherMethod,             // line #1974:       SetOrMapLiteral exp = decl.initializer;
		OtherMethod,             // line #1975:       return exp;
		OtherMethod,             // line #1976:     }
		OtherMethod,             // line #1977:
		OtherMethod,             // line #1978:     Asserter<InterfaceType> assertListOfString = _isListOf(_isString);
		OtherMethod,             // line #1979:     Asserter<InterfaceType> assertMapOfIntToListOfString =
		OtherMethod,             // line #1980:         _isMapOf(_isInt, (DartType type) => assertListOfString(type));
		OtherMethod,             // line #1981:
		OtherMethod,             // line #1982:     assertMapOfIntToListOfString(literal(0).staticType);
		OtherMethod,             // line #1983:     assertMapOfIntToListOfString(literal(1).staticType);
		OtherMethod,             // line #1984:     assertMapOfIntToListOfString(literal(2).staticType);
		OtherMethod,             // line #1985:     assertMapOfIntToListOfString(literal(3).staticType);
		OtherMethod,             // line #1986:     assertMapOfIntToListOfString(literal(4).staticType);
		OtherMethod,             // line #1987:
		OtherMethod,             // line #1988:     assertListOfString(
		OtherMethod,             // line #1989:         (literal(1).elements[0] as MapLiteralEntry).value.staticType);
		OtherMethod,             // line #1990:     assertListOfString(
		OtherMethod,             // line #1991:         (literal(2).elements[0] as MapLiteralEntry).value.staticType);
		OtherMethod,             // line #1992:     assertListOfString(
		OtherMethod,             // line #1993:         (literal(3).elements[0] as MapLiteralEntry).value.staticType);
		OtherMethod,             // line #1994:     assertListOfString(
		OtherMethod,             // line #1995:         (literal(4).elements[0] as MapLiteralEntry).value.staticType);
		OtherMethod,             // line #1996:   }
		BlankLine,               // line #1997:
		OtherMethod,             // line #1998:   test_mapLiteral_simple() async {
		OtherMethod,             // line #1999:     String code = r'''
		OtherMethod,             // line #2000:       void main () {
		OtherMethod,             // line #2001:         Map<int, String> l0 = {};
		OtherMethod,             // line #2002:         Map<int, String> l1 = {3: "hello"};
		OtherMethod,             // line #2003:         Map<int, String> l2 = {"hello": "hello"};
		OtherMethod,             // line #2004:         Map<int, String> l3 = {3: 3};
		OtherMethod,             // line #2005:         Map<int, String> l4 = {3:"hello", "hello": 3};
		OtherMethod,             // line #2006:      }
		OtherMethod,             // line #2007:    ''';
		OtherMethod,             // line #2008:     await assertErrorsInCode(code, [
		OtherMethod,             // line #2009:       error(HintCode.UNUSED_LOCAL_VARIABLE, 46, 2),
		OtherMethod,             // line #2010:       error(HintCode.UNUSED_LOCAL_VARIABLE, 80, 2),
		OtherMethod,             // line #2011:       error(HintCode.UNUSED_LOCAL_VARIABLE, 124, 2),
		OtherMethod,             // line #2012:       error(CompileTimeErrorCode.MAP_KEY_TYPE_NOT_ASSIGNABLE, 130, 7),
		OtherMethod,             // line #2013:       error(HintCode.UNUSED_LOCAL_VARIABLE, 174, 2),
		OtherMethod,             // line #2014:       error(CompileTimeErrorCode.MAP_VALUE_TYPE_NOT_ASSIGNABLE, 183, 1),
		OtherMethod,             // line #2015:       error(HintCode.UNUSED_LOCAL_VARIABLE, 212, 2),
		OtherMethod,             // line #2016:       error(CompileTimeErrorCode.MAP_KEY_TYPE_NOT_ASSIGNABLE, 229, 7),
		OtherMethod,             // line #2017:       error(CompileTimeErrorCode.MAP_VALUE_TYPE_NOT_ASSIGNABLE, 238, 1),
		OtherMethod,             // line #2018:     ]);
		OtherMethod,             // line #2019:
		OtherMethod,             // line #2020:     List<Statement> statements =
		OtherMethod,             // line #2021:         AstFinder.getStatementsInTopLevelFunction(unit, "main");
		OtherMethod,             // line #2022:     DartType literal(int i) {
		OtherMethod,             // line #2023:       VariableDeclarationStatement stmt = statements[i];
		OtherMethod,             // line #2024:       VariableDeclaration decl = stmt.variables.variables[0];
		OtherMethod,             // line #2025:       SetOrMapLiteral exp = decl.initializer;
		OtherMethod,             // line #2026:       return exp.staticType;
		OtherMethod,             // line #2027:     }
		OtherMethod,             // line #2028:
		OtherMethod,             // line #2029:     Asserter<InterfaceType> assertMapOfIntToString =
		OtherMethod,             // line #2030:         _isMapOf(_isInt, _isString);
		OtherMethod,             // line #2031:
		OtherMethod,             // line #2032:     assertMapOfIntToString(literal(0));
		OtherMethod,             // line #2033:     assertMapOfIntToString(literal(1));
		OtherMethod,             // line #2034:     assertMapOfIntToString(literal(2));
		OtherMethod,             // line #2035:     assertMapOfIntToString(literal(3));
		OtherMethod,             // line #2036:   }
		BlankLine,               // line #2037:
		OtherMethod,             // line #2038:   test_mapLiteral_simple_disabled() async {
		OtherMethod,             // line #2039:     String code = r'''
		OtherMethod,             // line #2040:       void main () {
		OtherMethod,             // line #2041:         Map<int, String> l0 = <int, dynamic>{};
		OtherMethod,             // line #2042:         Map<int, String> l1 = <int, dynamic>{3: "hello"};
		OtherMethod,             // line #2043:         Map<int, String> l2 = <int, dynamic>{"hello": "hello"};
		OtherMethod,             // line #2044:         Map<int, String> l3 = <int, dynamic>{3: 3};
		OtherMethod,             // line #2045:      }
		OtherMethod,             // line #2046:    ''';
		OtherMethod,             // line #2047:     await assertErrorsInCode(code, [
		OtherMethod,             // line #2048:       error(HintCode.UNUSED_LOCAL_VARIABLE, 46, 2),
		OtherMethod,             // line #2049:       error(CompileTimeErrorCode.INVALID_CAST_LITERAL_MAP, 51, 16),
		OtherMethod,             // line #2050:       error(HintCode.UNUSED_LOCAL_VARIABLE, 94, 2),
		OtherMethod,             // line #2051:       error(CompileTimeErrorCode.INVALID_CAST_LITERAL_MAP, 99, 26),
		OtherMethod,             // line #2052:       error(HintCode.UNUSED_LOCAL_VARIABLE, 152, 2),
		OtherMethod,             // line #2053:       error(CompileTimeErrorCode.INVALID_CAST_LITERAL_MAP, 157, 32),
		OtherMethod,             // line #2054:       error(CompileTimeErrorCode.MAP_KEY_TYPE_NOT_ASSIGNABLE, 172, 7),
		OtherMethod,             // line #2055:       error(HintCode.UNUSED_LOCAL_VARIABLE, 216, 2),
		OtherMethod,             // line #2056:       error(CompileTimeErrorCode.INVALID_CAST_LITERAL_MAP, 221, 20),
		OtherMethod,             // line #2057:     ]);
		OtherMethod,             // line #2058:
		OtherMethod,             // line #2059:     List<Statement> statements =
		OtherMethod,             // line #2060:         AstFinder.getStatementsInTopLevelFunction(unit, "main");
		OtherMethod,             // line #2061:     DartType literal(int i) {
		OtherMethod,             // line #2062:       VariableDeclarationStatement stmt = statements[i];
		OtherMethod,             // line #2063:       VariableDeclaration decl = stmt.variables.variables[0];
		OtherMethod,             // line #2064:       SetOrMapLiteral exp = decl.initializer;
		OtherMethod,             // line #2065:       return exp.staticType;
		OtherMethod,             // line #2066:     }
		OtherMethod,             // line #2067:
		OtherMethod,             // line #2068:     Asserter<InterfaceType> assertMapOfIntToDynamic =
		OtherMethod,             // line #2069:         _isMapOf(_isInt, _isDynamic);
		OtherMethod,             // line #2070:
		OtherMethod,             // line #2071:     assertMapOfIntToDynamic(literal(0));
		OtherMethod,             // line #2072:     assertMapOfIntToDynamic(literal(1));
		OtherMethod,             // line #2073:     assertMapOfIntToDynamic(literal(2));
		OtherMethod,             // line #2074:     assertMapOfIntToDynamic(literal(3));
		OtherMethod,             // line #2075:   }
		BlankLine,               // line #2076:
		OtherMethod,             // line #2077:   test_methodDeclaration_body_propagation() async {
		OtherMethod,             // line #2078:     String code = r'''
		OtherMethod,             // line #2079:       class A {
		OtherMethod,             // line #2080:         List<String> m0(int x) => ["hello"];
		OtherMethod,             // line #2081:         List<String> m1(int x) {return [3];}
		OtherMethod,             // line #2082:       }
		OtherMethod,             // line #2083:    ''';
		OtherMethod,             // line #2084:     await assertErrorsInCode(code, [
		OtherMethod,             // line #2085:       error(CompileTimeErrorCode.LIST_ELEMENT_TYPE_NOT_ASSIGNABLE, 101, 1),
		OtherMethod,             // line #2086:     ]);
		OtherMethod,             // line #2087:
		OtherMethod,             // line #2088:     Expression methodReturnValue(String methodName) {
		OtherMethod,             // line #2089:       MethodDeclaration method =
		OtherMethod,             // line #2090:           AstFinder.getMethodInClass(unit, "A", methodName);
		OtherMethod,             // line #2091:       FunctionBody body = method.body;
		OtherMethod,             // line #2092:       if (body is ExpressionFunctionBody) {
		OtherMethod,             // line #2093:         return body.expression;
		OtherMethod,             // line #2094:       } else {
		OtherMethod,             // line #2095:         Statement stmt = (body as BlockFunctionBody).block.statements[0];
		OtherMethod,             // line #2096:         return (stmt as ReturnStatement).expression;
		OtherMethod,             // line #2097:       }
		OtherMethod,             // line #2098:     }
		OtherMethod,             // line #2099:
		OtherMethod,             // line #2100:     Asserter<InterfaceType> assertListOfString = _isListOf(_isString);
		OtherMethod,             // line #2101:     assertListOfString(methodReturnValue("m0").staticType);
		OtherMethod,             // line #2102:     assertListOfString(methodReturnValue("m1").staticType);
		OtherMethod,             // line #2103:   }
		BlankLine,               // line #2104:
		OtherMethod,             // line #2105:   test_partialTypes1() async {
		OtherMethod,             // line #2106:     // Test that downwards inference with a partial type
		OtherMethod,             // line #2107:     // correctly uses the partial information to fill in subterm
		OtherMethod,             // line #2108:     // types
		OtherMethod,             // line #2109:     String code = r'''
		OtherMethod,             // line #2110:     typedef To Func1<From, To>(From x);
		OtherMethod,             // line #2111:     S f<S, T>(Func1<S, T> g) => null;
		OtherMethod,             // line #2112:     String test() => f((l) => l.length);
		OtherMethod,             // line #2113:    ''';
		OtherMethod,             // line #2114:     await assertNoErrorsInCode(code);
		OtherMethod,             // line #2115:
		OtherMethod,             // line #2116:     FunctionDeclaration test = AstFinder.getTopLevelFunction(unit, "test");
		OtherMethod,             // line #2117:     ExpressionFunctionBody body = test.functionExpression.body;
		OtherMethod,             // line #2118:     _isString(body.expression.staticType);
		OtherMethod,             // line #2119:     MethodInvocation invoke = body.expression;
		OtherMethod,             // line #2120:     FunctionExpression function = invoke.argumentList.arguments[0];
		OtherMethod,             // line #2121:     ExecutableElement f0 = function.declaredElement;
		OtherMethod,             // line #2122:     FunctionType type = f0.type;
		OtherMethod,             // line #2123:     _isFunction2Of(_isString, _isInt)(type);
		OtherMethod,             // line #2124:   }
		BlankLine,               // line #2125:
		OtherMethod,             // line #2126:   test_pinning_multipleConstraints1() async {
		OtherMethod,             // line #2127:     // Test that downwards inference with two different downwards covariant
		OtherMethod,             // line #2128:     // constraints on the same parameter correctly fails to infer when
		OtherMethod,             // line #2129:     // the types do not share a common subtype
		OtherMethod,             // line #2130:     String code = r'''
		OtherMethod,             // line #2131:     class A<S, T> {
		OtherMethod,             // line #2132:       S s;
		OtherMethod,             // line #2133:       T t;
		OtherMethod,             // line #2134:     }
		OtherMethod,             // line #2135:     class B<S> extends A<S, S> { B(S s); }
		OtherMethod,             // line #2136:     A<int, String> test() => new B(3);
		OtherMethod,             // line #2137:    ''';
		OtherMethod,             // line #2138:     await assertErrorsInCode(code, [
		OtherMethod,             // line #2139:       error(CompileTimeErrorCode.INVALID_CAST_LITERAL, 126, 1),
		OtherMethod,             // line #2140:     ]);
		OtherMethod,             // line #2141:
		OtherMethod,             // line #2142:     FunctionDeclaration test = AstFinder.getTopLevelFunction(unit, "test");
		OtherMethod,             // line #2143:     ExpressionFunctionBody body = test.functionExpression.body;
		OtherMethod,             // line #2144:     DartType type = body.expression.staticType;
		OtherMethod,             // line #2145:
		OtherMethod,             // line #2146:     Element elementB = AstFinder.getClass(unit, "B").declaredElement;
		OtherMethod,             // line #2147:
		OtherMethod,             // line #2148:     _isInstantiationOf(_hasElement(elementB))([_isNull])(type);
		OtherMethod,             // line #2149:   }
		BlankLine,               // line #2150:
		OtherMethod,             // line #2151:   test_pinning_multipleConstraints2() async {
		OtherMethod,             // line #2152:     // Test that downwards inference with two identical downwards covariant
		OtherMethod,             // line #2153:     // constraints on the same parameter correctly infers and pins the type
		OtherMethod,             // line #2154:     String code = r'''
		OtherMethod,             // line #2155:     class A<S, T> {
		OtherMethod,             // line #2156:       S s;
		OtherMethod,             // line #2157:       T t;
		OtherMethod,             // line #2158:     }
		OtherMethod,             // line #2159:     class B<S> extends A<S, S> { B(S s); }
		OtherMethod,             // line #2160:     A<num, num> test() => new B(3);
		OtherMethod,             // line #2161:    ''';
		OtherMethod,             // line #2162:     await assertNoErrorsInCode(code);
		OtherMethod,             // line #2163:
		OtherMethod,             // line #2164:     FunctionDeclaration test = AstFinder.getTopLevelFunction(unit, "test");
		OtherMethod,             // line #2165:     ExpressionFunctionBody body = test.functionExpression.body;
		OtherMethod,             // line #2166:     DartType type = body.expression.staticType;
		OtherMethod,             // line #2167:
		OtherMethod,             // line #2168:     Element elementB = AstFinder.getClass(unit, "B").declaredElement;
		OtherMethod,             // line #2169:
		OtherMethod,             // line #2170:     _isInstantiationOf(_hasElement(elementB))([_isNum])(type);
		OtherMethod,             // line #2171:   }
		BlankLine,               // line #2172:
		OtherMethod,             // line #2173:   test_pinning_multipleConstraints3() async {
		OtherMethod,             // line #2174:     // Test that downwards inference with two different downwards covariant
		OtherMethod,             // line #2175:     // constraints on the same parameter correctly fails to infer when
		OtherMethod,             // line #2176:     // the types do not share a common subtype, but do share a common supertype
		OtherMethod,             // line #2177:     String code = r'''
		OtherMethod,             // line #2178:     class A<S, T> {
		OtherMethod,             // line #2179:       S s;
		OtherMethod,             // line #2180:       T t;
		OtherMethod,             // line #2181:     }
		OtherMethod,             // line #2182:     class B<S> extends A<S, S> { B(S s); }
		OtherMethod,             // line #2183:     A<int, double> test() => new B(3);
		OtherMethod,             // line #2184:    ''';
		OtherMethod,             // line #2185:     await assertErrorsInCode(code, [
		OtherMethod,             // line #2186:       error(CompileTimeErrorCode.INVALID_CAST_LITERAL, 126, 1),
		OtherMethod,             // line #2187:     ]);
		OtherMethod,             // line #2188:
		OtherMethod,             // line #2189:     FunctionDeclaration test = AstFinder.getTopLevelFunction(unit, "test");
		OtherMethod,             // line #2190:     ExpressionFunctionBody body = test.functionExpression.body;
		OtherMethod,             // line #2191:     DartType type = body.expression.staticType;
		OtherMethod,             // line #2192:
		OtherMethod,             // line #2193:     Element elementB = AstFinder.getClass(unit, "B").declaredElement;
		OtherMethod,             // line #2194:
		OtherMethod,             // line #2195:     _isInstantiationOf(_hasElement(elementB))([_isNull])(type);
		OtherMethod,             // line #2196:   }
		BlankLine,               // line #2197:
		OtherMethod,             // line #2198:   test_pinning_multipleConstraints4() async {
		OtherMethod,             // line #2199:     // Test that downwards inference with two subtype related downwards
		OtherMethod,             // line #2200:     // covariant constraints on the same parameter correctly infers and pins
		OtherMethod,             // line #2201:     // the type
		OtherMethod,             // line #2202:     String code = r'''
		OtherMethod,             // line #2203:     class A<S, T> {
		OtherMethod,             // line #2204:       S s;
		OtherMethod,             // line #2205:       T t;
		OtherMethod,             // line #2206:     }
		OtherMethod,             // line #2207:     class B<S> extends A<S, S> {}
		OtherMethod,             // line #2208:     A<int, num> test() => new B();
		OtherMethod,             // line #2209:    ''';
		OtherMethod,             // line #2210:     await assertNoErrorsInCode(code);
		OtherMethod,             // line #2211:
		OtherMethod,             // line #2212:     FunctionDeclaration test = AstFinder.getTopLevelFunction(unit, "test");
		OtherMethod,             // line #2213:     ExpressionFunctionBody body = test.functionExpression.body;
		OtherMethod,             // line #2214:     DartType type = body.expression.staticType;
		OtherMethod,             // line #2215:
		OtherMethod,             // line #2216:     Element elementB = AstFinder.getClass(unit, "B").declaredElement;
		OtherMethod,             // line #2217:
		OtherMethod,             // line #2218:     _isInstantiationOf(_hasElement(elementB))([_isInt])(type);
		OtherMethod,             // line #2219:   }
		BlankLine,               // line #2220:
		OtherMethod,             // line #2221:   test_pinning_multipleConstraints_contravariant1() async {
		OtherMethod,             // line #2222:     // Test that downwards inference with two different downwards contravariant
		OtherMethod,             // line #2223:     // constraints on the same parameter chooses the upper bound
		OtherMethod,             // line #2224:     // when the only supertype is Object
		OtherMethod,             // line #2225:     String code = r'''
		OtherMethod,             // line #2226:     class A<S, T> {
		OtherMethod,             // line #2227:       S s;
		OtherMethod,             // line #2228:       T t;
		OtherMethod,             // line #2229:     }
		OtherMethod,             // line #2230:     class B<S> extends A<S, S> {}
		OtherMethod,             // line #2231:     typedef void Contra1<T>(T x);
		OtherMethod,             // line #2232:     Contra1<A<S, S>> mkA<S>() => (A<S, S> x) {};
		OtherMethod,             // line #2233:     Contra1<A<int, String>> test() => mkA();
		OtherMethod,             // line #2234:    ''';
		OtherMethod,             // line #2235:     await assertNoErrorsInCode(code);
		OtherMethod,             // line #2236:
		OtherMethod,             // line #2237:     FunctionDeclaration test = AstFinder.getTopLevelFunction(unit, "test");
		OtherMethod,             // line #2238:     ExpressionFunctionBody body = test.functionExpression.body;
		OtherMethod,             // line #2239:     FunctionType functionType = body.expression.staticType;
		OtherMethod,             // line #2240:     DartType type = functionType.normalParameterTypes[0];
		OtherMethod,             // line #2241:
		OtherMethod,             // line #2242:     Element elementA = AstFinder.getClass(unit, "A").declaredElement;
		OtherMethod,             // line #2243:
		OtherMethod,             // line #2244:     _isInstantiationOf(_hasElement(elementA))([_isObject, _isObject])(type);
		OtherMethod,             // line #2245:   }
		BlankLine,               // line #2246:
		OtherMethod,             // line #2247:   test_pinning_multipleConstraints_contravariant2() async {
		OtherMethod,             // line #2248:     // Test that downwards inference with two identical downwards contravariant
		OtherMethod,             // line #2249:     // constraints on the same parameter correctly pins the type
		OtherMethod,             // line #2250:     String code = r'''
		OtherMethod,             // line #2251:     class A<S, T> {
		OtherMethod,             // line #2252:       S s;
		OtherMethod,             // line #2253:       T t;
		OtherMethod,             // line #2254:     }
		OtherMethod,             // line #2255:     class B<S> extends A<S, S> {}
		OtherMethod,             // line #2256:     typedef void Contra1<T>(T x);
		OtherMethod,             // line #2257:     Contra1<A<S, S>> mkA<S>() => (A<S, S> x) {};
		OtherMethod,             // line #2258:     Contra1<A<num, num>> test() => mkA();
		OtherMethod,             // line #2259:    ''';
		OtherMethod,             // line #2260:     await assertNoErrorsInCode(code);
		OtherMethod,             // line #2261:
		OtherMethod,             // line #2262:     FunctionDeclaration test = AstFinder.getTopLevelFunction(unit, "test");
		OtherMethod,             // line #2263:     ExpressionFunctionBody body = test.functionExpression.body;
		OtherMethod,             // line #2264:     FunctionType functionType = body.expression.staticType;
		OtherMethod,             // line #2265:     DartType type = functionType.normalParameterTypes[0];
		OtherMethod,             // line #2266:
		OtherMethod,             // line #2267:     Element elementA = AstFinder.getClass(unit, "A").declaredElement;
		OtherMethod,             // line #2268:
		OtherMethod,             // line #2269:     _isInstantiationOf(_hasElement(elementA))([_isNum, _isNum])(type);
		OtherMethod,             // line #2270:   }
		BlankLine,               // line #2271:
		OtherMethod,             // line #2272:   test_pinning_multipleConstraints_contravariant3() async {
		OtherMethod,             // line #2273:     // Test that downwards inference with two different downwards contravariant
		OtherMethod,             // line #2274:     // constraints on the same parameter correctly choose the least upper bound
		OtherMethod,             // line #2275:     // when they share a common supertype
		OtherMethod,             // line #2276:     String code = r'''
		OtherMethod,             // line #2277:     class A<S, T> {
		OtherMethod,             // line #2278:       S s;
		OtherMethod,             // line #2279:       T t;
		OtherMethod,             // line #2280:     }
		OtherMethod,             // line #2281:     class B<S> extends A<S, S> {}
		OtherMethod,             // line #2282:     typedef void Contra1<T>(T x);
		OtherMethod,             // line #2283:     Contra1<A<S, S>> mkA<S>() => (A<S, S> x) {};
		OtherMethod,             // line #2284:     Contra1<A<int, double>> test() => mkA();
		OtherMethod,             // line #2285:    ''';
		OtherMethod,             // line #2286:     await assertNoErrorsInCode(code);
		OtherMethod,             // line #2287:
		OtherMethod,             // line #2288:     FunctionDeclaration test = AstFinder.getTopLevelFunction(unit, "test");
		OtherMethod,             // line #2289:     ExpressionFunctionBody body = test.functionExpression.body;
		OtherMethod,             // line #2290:     FunctionType functionType = body.expression.staticType;
		OtherMethod,             // line #2291:     DartType type = functionType.normalParameterTypes[0];
		OtherMethod,             // line #2292:
		OtherMethod,             // line #2293:     Element elementA = AstFinder.getClass(unit, "A").declaredElement;
		OtherMethod,             // line #2294:
		OtherMethod,             // line #2295:     _isInstantiationOf(_hasElement(elementA))([_isNum, _isNum])(type);
		OtherMethod,             // line #2296:   }
		BlankLine,               // line #2297:
		OtherMethod,             // line #2298:   test_pinning_multipleConstraints_contravariant4() async {
		OtherMethod,             // line #2299:     // Test that downwards inference with two different downwards contravariant
		OtherMethod,             // line #2300:     // constraints on the same parameter correctly choose the least upper bound
		OtherMethod,             // line #2301:     // when one is a subtype of the other
		OtherMethod,             // line #2302:     String code = r'''
		OtherMethod,             // line #2303:     class A<S, T> {
		OtherMethod,             // line #2304:       S s;
		OtherMethod,             // line #2305:       T t;
		OtherMethod,             // line #2306:     }
		OtherMethod,             // line #2307:     class B<S> extends A<S, S> {}
		OtherMethod,             // line #2308:     typedef void Contra1<T>(T x);
		OtherMethod,             // line #2309:     Contra1<A<S, S>> mkA<S>() => (A<S, S> x) {};
		OtherMethod,             // line #2310:     Contra1<A<int, num>> test() => mkA();
		OtherMethod,             // line #2311:    ''';
		OtherMethod,             // line #2312:     await assertNoErrorsInCode(code);
		OtherMethod,             // line #2313:
		OtherMethod,             // line #2314:     FunctionDeclaration test = AstFinder.getTopLevelFunction(unit, "test");
		OtherMethod,             // line #2315:     ExpressionFunctionBody body = test.functionExpression.body;
		OtherMethod,             // line #2316:     FunctionType functionType = body.expression.staticType;
		OtherMethod,             // line #2317:     DartType type = functionType.normalParameterTypes[0];
		OtherMethod,             // line #2318:
		OtherMethod,             // line #2319:     Element elementA = AstFinder.getClass(unit, "A").declaredElement;
		OtherMethod,             // line #2320:
		OtherMethod,             // line #2321:     _isInstantiationOf(_hasElement(elementA))([_isNum, _isNum])(type);
		OtherMethod,             // line #2322:   }
		BlankLine,               // line #2323:
		OtherMethod,             // line #2324:   test_redirectedConstructor_named() async {
		OtherMethod,             // line #2325:     var code = r'''
		OtherMethod,             // line #2326: class A<T, U> implements B<T, U> {
		OtherMethod,             // line #2327:   A.named();
		OtherMethod,             // line #2328: }
		OtherMethod,             // line #2329:
		OtherMethod,             // line #2330: class B<T2, U2> {
		OtherMethod,             // line #2331:   factory B() = A.named;
		OtherMethod,             // line #2332: }
		OtherMethod,             // line #2333:    ''';
		OtherMethod,             // line #2334:     await assertNoErrorsInCode(code);
		OtherMethod,             // line #2335:
		OtherMethod,             // line #2336:     ClassDeclaration b = unit.declarations[1];
		OtherMethod,             // line #2337:     ConstructorDeclaration bConstructor = b.members[0];
		OtherMethod,             // line #2338:     ConstructorName redirected = bConstructor.redirectedConstructor;
		OtherMethod,             // line #2339:
		OtherMethod,             // line #2340:     TypeName typeName = redirected.type;
		OtherMethod,             // line #2341:     assertType(typeName.type, 'A<T2, U2>');
		OtherMethod,             // line #2342:     assertType(typeName.type, 'A<T2, U2>');
		OtherMethod,             // line #2343:
		OtherMethod,             // line #2344:     var constructorMember = redirected.staticElement;
		OtherMethod,             // line #2345:     expect(
		OtherMethod,             // line #2346:       constructorMember.getDisplayString(withNullability: false),
		OtherMethod,             // line #2347:       'A<T2, U2> A.named()',
		OtherMethod,             // line #2348:     );
		OtherMethod,             // line #2349:     expect(redirected.name.staticElement, constructorMember);
		OtherMethod,             // line #2350:   }
		BlankLine,               // line #2351:
		OtherMethod,             // line #2352:   test_redirectedConstructor_self() async {
		OtherMethod,             // line #2353:     await assertNoErrorsInCode(r'''
		OtherMethod,             // line #2354: class A<T> {
		OtherMethod,             // line #2355:   A();
		OtherMethod,             // line #2356:   factory A.redirected() = A;
		OtherMethod,             // line #2357: }
		OtherMethod,             // line #2358: ''');
		OtherMethod,             // line #2359:   }
		BlankLine,               // line #2360:
		OtherMethod,             // line #2361:   test_redirectedConstructor_unnamed() async {
		OtherMethod,             // line #2362:     await assertNoErrorsInCode(r'''
		OtherMethod,             // line #2363: class A<T, U> implements B<T, U> {
		OtherMethod,             // line #2364:   A();
		OtherMethod,             // line #2365: }
		OtherMethod,             // line #2366:
		OtherMethod,             // line #2367: class B<T2, U2> {
		OtherMethod,             // line #2368:   factory B() = A;
		OtherMethod,             // line #2369: }
		OtherMethod,             // line #2370: ''');
		OtherMethod,             // line #2371:
		OtherMethod,             // line #2372:     ClassDeclaration b = result.unit.declarations[1];
		OtherMethod,             // line #2373:     ConstructorDeclaration bConstructor = b.members[0];
		OtherMethod,             // line #2374:     ConstructorName redirected = bConstructor.redirectedConstructor;
		OtherMethod,             // line #2375:
		OtherMethod,             // line #2376:     TypeName typeName = redirected.type;
		OtherMethod,             // line #2377:     assertType(typeName.type, 'A<T2, U2>');
		OtherMethod,             // line #2378:     assertType(typeName.type, 'A<T2, U2>');
		OtherMethod,             // line #2379:
		OtherMethod,             // line #2380:     expect(redirected.name, isNull);
		OtherMethod,             // line #2381:     expect(
		OtherMethod,             // line #2382:       redirected.staticElement.getDisplayString(withNullability: false),
		OtherMethod,             // line #2383:       'A<T2, U2> A()',
		OtherMethod,             // line #2384:     );
		OtherMethod,             // line #2385:   }
		BlankLine,               // line #2386:
		OtherMethod,             // line #2387:   test_redirectingConstructor_propagation() async {
		OtherMethod,             // line #2388:     String code = r'''
		OtherMethod,             // line #2389:       class A {
		OtherMethod,             // line #2390:         A() : this.named([]);
		OtherMethod,             // line #2391:         A.named(List<String> x);
		OtherMethod,             // line #2392:       }
		OtherMethod,             // line #2393:    ''';
		OtherMethod,             // line #2394:     await assertNoErrorsInCode(code);
		OtherMethod,             // line #2395:
		OtherMethod,             // line #2396:     ConstructorDeclaration constructor =
		OtherMethod,             // line #2397:         AstFinder.getConstructorInClass(unit, "A", null);
		OtherMethod,             // line #2398:     RedirectingConstructorInvocation invocation = constructor.initializers[0];
		OtherMethod,             // line #2399:     Expression exp = invocation.argumentList.arguments[0];
		OtherMethod,             // line #2400:     _isListOf(_isString)(exp.staticType);
		OtherMethod,             // line #2401:   }
		BlankLine,               // line #2402:
		OtherMethod,             // line #2403:   test_returnType_variance1() async {
		OtherMethod,             // line #2404:     // Check that downwards inference correctly pins a type parameter
		OtherMethod,             // line #2405:     // when the parameter is constrained in a contravariant position
		OtherMethod,             // line #2406:     String code = r'''
		OtherMethod,             // line #2407:     typedef To Func1<From, To>(From x);
		OtherMethod,             // line #2408:     Func1<T, String> f<T>(T x) => null;
		OtherMethod,             // line #2409:     Func1<num, String> test() => f(42);
		OtherMethod,             // line #2410:    ''';
		OtherMethod,             // line #2411:     await assertNoErrorsInCode(code);
		OtherMethod,             // line #2412:
		OtherMethod,             // line #2413:     FunctionDeclaration test = AstFinder.getTopLevelFunction(unit, "test");
		OtherMethod,             // line #2414:     ExpressionFunctionBody body = test.functionExpression.body;
		OtherMethod,             // line #2415:     MethodInvocation invoke = body.expression;
		OtherMethod,             // line #2416:     _isFunction2Of(_isNum, _isFunction2Of(_isNum, _isString))(
		OtherMethod,             // line #2417:         invoke.staticInvokeType);
		OtherMethod,             // line #2418:   }
		BlankLine,               // line #2419:
		OtherMethod,             // line #2420:   test_returnType_variance2() async {
		OtherMethod,             // line #2421:     // Check that downwards inference correctly pins a type parameter
		OtherMethod,             // line #2422:     // when the parameter is constrained in a covariant position
		OtherMethod,             // line #2423:     String code = r'''
		OtherMethod,             // line #2424:     typedef To Func1<From, To>(From x);
		OtherMethod,             // line #2425:     Func1<String, T> f<T>(T x) => null;
		OtherMethod,             // line #2426:     Func1<String, num> test() => f(42);
		OtherMethod,             // line #2427:    ''';
		OtherMethod,             // line #2428:     await assertNoErrorsInCode(code);
		OtherMethod,             // line #2429:
		OtherMethod,             // line #2430:     FunctionDeclaration test = AstFinder.getTopLevelFunction(unit, "test");
		OtherMethod,             // line #2431:     ExpressionFunctionBody body = test.functionExpression.body;
		OtherMethod,             // line #2432:     MethodInvocation invoke = body.expression;
		OtherMethod,             // line #2433:     _isFunction2Of(_isNum, _isFunction2Of(_isString, _isNum))(
		OtherMethod,             // line #2434:         invoke.staticInvokeType);
		OtherMethod,             // line #2435:   }
		BlankLine,               // line #2436:
		OtherMethod,             // line #2437:   test_returnType_variance3() async {
		OtherMethod,             // line #2438:     // Check that the variance heuristic chooses the most precise type
		OtherMethod,             // line #2439:     // when the return type uses the variable in a contravariant position
		OtherMethod,             // line #2440:     // and there is no downwards constraint.
		OtherMethod,             // line #2441:     String code = r'''
		OtherMethod,             // line #2442:     typedef To Func1<From, To>(From x);
		OtherMethod,             // line #2443:     Func1<T, String> f<T>(T x, g(T x)) => null;
		OtherMethod,             // line #2444:     dynamic test() => f(42, (num x) => x);
		OtherMethod,             // line #2445:    ''';
		OtherMethod,             // line #2446:     await assertNoErrorsInCode(code);
		OtherMethod,             // line #2447:
		OtherMethod,             // line #2448:     FunctionDeclaration test = AstFinder.getTopLevelFunction(unit, "test");
		OtherMethod,             // line #2449:     ExpressionFunctionBody body = test.functionExpression.body;
		OtherMethod,             // line #2450:     FunctionType functionType = body.expression.staticType;
		OtherMethod,             // line #2451:     DartType type = functionType.normalParameterTypes[0];
		OtherMethod,             // line #2452:     _isInt(type);
		OtherMethod,             // line #2453:   }
		BlankLine,               // line #2454:
		OtherMethod,             // line #2455:   test_returnType_variance4() async {
		OtherMethod,             // line #2456:     // Check that the variance heuristic chooses the more precise type
		OtherMethod,             // line #2457:     // when the return type uses the variable in a covariant position
		OtherMethod,             // line #2458:     // and there is no downwards constraint
		OtherMethod,             // line #2459:     String code = r'''
		OtherMethod,             // line #2460:     typedef To Func1<From, To>(From x);
		OtherMethod,             // line #2461:     Func1<String, T> f<T>(T x, g(T x)) => null;
		OtherMethod,             // line #2462:     dynamic test() => f(42, (num x) => x);
		OtherMethod,             // line #2463:    ''';
		OtherMethod,             // line #2464:     await assertNoErrorsInCode(code);
		OtherMethod,             // line #2465:
		OtherMethod,             // line #2466:     FunctionDeclaration test = AstFinder.getTopLevelFunction(unit, "test");
		OtherMethod,             // line #2467:     ExpressionFunctionBody body = test.functionExpression.body;
		OtherMethod,             // line #2468:     FunctionType functionType = body.expression.staticType;
		OtherMethod,             // line #2469:     DartType type = functionType.returnType;
		OtherMethod,             // line #2470:     _isInt(type);
		OtherMethod,             // line #2471:   }
		BlankLine,               // line #2472:
		OtherMethod,             // line #2473:   test_returnType_variance5() async {
		OtherMethod,             // line #2474:     // Check that pinning works correctly with a partial type
		OtherMethod,             // line #2475:     // when the return type uses the variable in a contravariant position
		OtherMethod,             // line #2476:     String code = r'''
		OtherMethod,             // line #2477:     typedef To Func1<From, To>(From x);
		OtherMethod,             // line #2478:     Func1<T, String> f<T>(T x) => null;
		OtherMethod,             // line #2479:     T g<T, S>(Func1<T, S> f) => null;
		OtherMethod,             // line #2480:     num test() => g(f(3));
		OtherMethod,             // line #2481:    ''';
		OtherMethod,             // line #2482:     await assertNoErrorsInCode(code);
		OtherMethod,             // line #2483:
		OtherMethod,             // line #2484:     FunctionDeclaration test = AstFinder.getTopLevelFunction(unit, "test");
		OtherMethod,             // line #2485:     ExpressionFunctionBody body = test.functionExpression.body;
		OtherMethod,             // line #2486:     MethodInvocation call = body.expression;
		OtherMethod,             // line #2487:     _isNum(call.staticType);
		OtherMethod,             // line #2488:     _isFunction2Of(_isFunction2Of(_isNum, _isString), _isNum)(
		OtherMethod,             // line #2489:         call.staticInvokeType);
		OtherMethod,             // line #2490:   }
		BlankLine,               // line #2491:
		OtherMethod,             // line #2492:   test_returnType_variance6() async {
		OtherMethod,             // line #2493:     // Check that pinning works correctly with a partial type
		OtherMethod,             // line #2494:     // when the return type uses the variable in a covariant position
		OtherMethod,             // line #2495:     String code = r'''
		OtherMethod,             // line #2496:     typedef To Func1<From, To>(From x);
		OtherMethod,             // line #2497:     Func1<String, T> f<T>(T x) => null;
		OtherMethod,             // line #2498:     T g<T, S>(Func1<S, T> f) => null;
		OtherMethod,             // line #2499:     num test() => g(f(3));
		OtherMethod,             // line #2500:    ''';
		OtherMethod,             // line #2501:     await assertNoErrorsInCode(code);
		OtherMethod,             // line #2502:
		OtherMethod,             // line #2503:     FunctionDeclaration test = AstFinder.getTopLevelFunction(unit, "test");
		OtherMethod,             // line #2504:     ExpressionFunctionBody body = test.functionExpression.body;
		OtherMethod,             // line #2505:     MethodInvocation call = body.expression;
		OtherMethod,             // line #2506:     _isNum(call.staticType);
		OtherMethod,             // line #2507:     _isFunction2Of(_isFunction2Of(_isString, _isNum), _isNum)(
		OtherMethod,             // line #2508:         call.staticInvokeType);
		OtherMethod,             // line #2509:   }
		BlankLine,               // line #2510:
		OtherMethod,             // line #2511:   test_superConstructorInvocation_propagation() async {
		OtherMethod,             // line #2512:     String code = r'''
		OtherMethod,             // line #2513:       class B {
		OtherMethod,             // line #2514:         B(List<String> p);
		OtherMethod,             // line #2515:       }
		OtherMethod,             // line #2516:       class A extends B {
		OtherMethod,             // line #2517:         A() : super([]);
		OtherMethod,             // line #2518:       }
		OtherMethod,             // line #2519:    ''';
		OtherMethod,             // line #2520:     await assertNoErrorsInCode(code);
		OtherMethod,             // line #2521:
		OtherMethod,             // line #2522:     ConstructorDeclaration constructor =
		OtherMethod,             // line #2523:         AstFinder.getConstructorInClass(unit, "A", null);
		OtherMethod,             // line #2524:     SuperConstructorInvocation invocation = constructor.initializers[0];
		OtherMethod,             // line #2525:     Expression exp = invocation.argumentList.arguments[0];
		OtherMethod,             // line #2526:     _isListOf(_isString)(exp.staticType);
		OtherMethod,             // line #2527:   }
		BlankLine,               // line #2528:
		OtherMethod,             // line #2529:   /// Verifies the result has [CompileTimeErrorCode.COULD_NOT_INFER] with
		OtherMethod,             // line #2530:   /// the expected [errorMessage].
		OtherMethod,             // line #2531:   void _expectInferenceError(String errorMessage) {
		OtherMethod,             // line #2532:     var errors = result.errors
		OtherMethod,             // line #2533:         .where((e) => e.errorCode == CompileTimeErrorCode.COULD_NOT_INFER)
		OtherMethod,             // line #2534:         .map((e) => e.message)
		OtherMethod,             // line #2535:         .toList();
		OtherMethod,             // line #2536:     expect(errors.length, 1);
		OtherMethod,             // line #2537:     var actual = errors[0];
		OtherMethod,             // line #2538:     expect(actual,
		OtherMethod,             // line #2539:         errorMessage, // Print the literal error message for easy copy+paste:
		OtherMethod,             // line #2540:         reason: 'Actual error did not match expected error:\n$actual');
		OtherMethod,             // line #2541:   }
		BlankLine,               // line #2542:
		OtherMethod,             // line #2543:   /// Helper method for testing `FutureOr<T>`.
		OtherMethod,             // line #2544:   ///
		OtherMethod,             // line #2545:   /// Validates that [code] produces [errors]. It should define a function
		OtherMethod,             // line #2546:   /// "test", whose body is an expression that invokes a method. Returns that
		OtherMethod,             // line #2547:   /// invocation.
		OtherMethod,             // line #2548:   Future<MethodInvocation> _testFutureOr(String code,
		OtherMethod,             // line #2549:       {List<ExpectedError> expectedErrors = const []}) async {
		OtherMethod,             // line #2550:     var fullCode = """
		OtherMethod,             // line #2551: import "dart:async";
		OtherMethod,             // line #2552:
		OtherMethod,             // line #2553: $code
		OtherMethod,             // line #2554: """;
		OtherMethod,             // line #2555:     await assertErrorsInCode(fullCode, expectedErrors);
		OtherMethod,             // line #2556:
		OtherMethod,             // line #2557:     FunctionDeclaration test = AstFinder.getTopLevelFunction(unit, "test");
		OtherMethod,             // line #2558:     ExpressionFunctionBody body = test.functionExpression.body;
		OtherMethod,             // line #2559:     return body.expression;
		OtherMethod,             // line #2560:   }
		BlankLine,               // line #2561:
	}

	runParsePhase(t, opts, source, want)
	// runFullStylizer(t, opts, source, wantSource, want)
}

func TestStrongMode_Class2(t *testing.T) {
	source := strong_mode_dart_txt[91956 : 129066+12] // extra utf8 bytes
	// wantSource := strong_mode_want_txt[560:769]

	if source[0] != '@' {
		t.Fatalf("source[0] = %c, want @", source[0])
	}
	if source[len(source)-1] != '\n' {
		t.Fatalf("source[%v] = %c, want newline", len(source)-1, source[len(source)-1])
	}

	opts := &Options{
		MemberOrdering: defaultMemberOrdering,
		// Verbose:        true,
	}

	want := []EntityType{
		Unknown,     // line #2: {
		OtherMethod, // line #3:   void expectStaticInvokeType(String search, String expected) {
		OtherMethod, // line #4:     var invocation = findNode.simple(search).parent as MethodInvocation;
		OtherMethod, // line #5:     assertInvokeType(invocation, expected);
		OtherMethod, // line #6:   }
		BlankLine,   // line #7:
		OtherMethod, // line #8:   test_dynamicObjectGetter_hashCode() async {
		OtherMethod, // line #9:     await assertErrorsInCode(r'''
		OtherMethod, // line #10: main() {
		OtherMethod, // line #11:   dynamic a = null;
		OtherMethod, // line #12:   var foo = a.hashCode;
		OtherMethod, // line #13: }
		OtherMethod, // line #14: ''', [
		OtherMethod, // line #15:       error(HintCode.UNUSED_LOCAL_VARIABLE, 35, 3),
		OtherMethod, // line #16:     ]);
		OtherMethod, // line #17:     expectInitializerType('foo', 'int');
		OtherMethod, // line #18:   }
		BlankLine,   // line #19:
		OtherMethod, // line #20:   test_futureOr_promotion1() async {
		OtherMethod, // line #21:     // Test that promotion from FutureOr<T> to T works for concrete types
		OtherMethod, // line #22:     String code = r'''
		OtherMethod, // line #23:     import "dart:async";
		OtherMethod, // line #24:     dynamic test(FutureOr<int> x) => (x is int) && (x.abs() == 0);
		OtherMethod, // line #25:    ''';
		OtherMethod, // line #26:     await assertNoErrorsInCode(code);
		OtherMethod, // line #27:   }
		BlankLine,   // line #28:
		OtherMethod, // line #29:   test_futureOr_promotion2() async {
		OtherMethod, // line #30:     // Test that promotion from FutureOr<T> to Future<T> works for concrete
		OtherMethod, // line #31:     // types
		OtherMethod, // line #32:     String code = r'''
		OtherMethod, // line #33:     import "dart:async";
		OtherMethod, // line #34:     dynamic test(FutureOr<int> x) => (x is Future<int>) &&
		OtherMethod, // line #35:                                      (x.then((x) => x) == null);
		OtherMethod, // line #36:    ''';
		OtherMethod, // line #37:     await assertNoErrorsInCode(code);
		OtherMethod, // line #38:   }
		BlankLine,   // line #39:
		OtherMethod, // line #40:   test_futureOr_promotion3() async {
		OtherMethod, // line #41:     // Test that promotion from FutureOr<T> to T works for type
		OtherMethod, // line #42:     // parameters T
		OtherMethod, // line #43:     String code = r'''
		OtherMethod, // line #44:     import "dart:async";
		OtherMethod, // line #45:     dynamic test<T extends num>(FutureOr<T> x) => (x is T) &&
		OtherMethod, // line #46:                                                   (x.abs() == 0);
		OtherMethod, // line #47:    ''';
		OtherMethod, // line #48:     await assertNoErrorsInCode(code);
		OtherMethod, // line #49:   }
		BlankLine,   // line #50:
		OtherMethod, // line #51:   test_futureOr_promotion4() async {
		OtherMethod, // line #52:     // Test that promotion from FutureOr<T> to Future<T> works for type
		OtherMethod, // line #53:     // parameters T
		OtherMethod, // line #54:     String code = r'''
		OtherMethod, // line #55:     import "dart:async";
		OtherMethod, // line #56:     dynamic test<T extends num>(FutureOr<T> x) => (x is Future<T>) &&
		OtherMethod, // line #57:                                                   (x.then((x) => x) == null);
		OtherMethod, // line #58:    ''';
		OtherMethod, // line #59:     await assertNoErrorsInCode(code);
		OtherMethod, // line #60:   }
		BlankLine,   // line #61:
		OtherMethod, // line #62:   test_generalizedVoid_assignToVoidOk() async {
		OtherMethod, // line #63:     await assertErrorsInCode(r'''
		OtherMethod, // line #64: void main() {
		OtherMethod, // line #65:   void x;
		OtherMethod, // line #66:   x = 42;
		OtherMethod, // line #67: }
		OtherMethod, // line #68: ''', [
		OtherMethod, // line #69:       error(HintCode.UNUSED_LOCAL_VARIABLE, 21, 1),
		OtherMethod, // line #70:     ]);
		OtherMethod, // line #71:   }
		BlankLine,   // line #72:
		OtherMethod, // line #73:   test_genericFunction() async {
		OtherMethod, // line #74:     await assertNoErrorsInCode(r'T f<T>(T x) => null;');
		OtherMethod, // line #75:     expectFunctionType('f', 'T Function<T>(T)', typeFormals: '[T]');
		OtherMethod, // line #76:     SimpleIdentifier f = findNode.simple('f');
		OtherMethod, // line #77:     FunctionElementImpl e = f.staticElement;
		OtherMethod, // line #78:     FunctionType ft = e.type.instantiate([typeProvider.stringType]);
		OtherMethod, // line #79:     assertType(ft, 'String Function(String)');
		OtherMethod, // line #80:   }
		BlankLine,   // line #81:
		OtherMethod, // line #82:   test_genericFunction_bounds() async {
		OtherMethod, // line #83:     await assertNoErrorsInCode(r'T f<T extends num>(T x) => null;');
		OtherMethod, // line #84:     expectFunctionType('f', 'T Function<T extends num>(T)',
		OtherMethod, // line #85:         typeFormals: '[T extends num]');
		OtherMethod, // line #86:   }
		BlankLine,   // line #87:
		OtherMethod, // line #88:   test_genericFunction_parameter() async {
		OtherMethod, // line #89:     await assertNoErrorsInCode(r'''
		OtherMethod, // line #90: void g(T f<T>(T x)) {}
		OtherMethod, // line #91: ''');
		OtherMethod, // line #92:     var type = expectFunctionType2('f', 'T Function<T>(T)');
		OtherMethod, // line #93:     FunctionType ft = type.instantiate([typeProvider.stringType]);
		OtherMethod, // line #94:     assertType(ft, 'String Function(String)');
		OtherMethod, // line #95:   }
		BlankLine,   // line #96:
		OtherMethod, // line #97:   test_genericFunction_static() async {
		OtherMethod, // line #98:     await assertNoErrorsInCode(r'''
		OtherMethod, // line #99: class C<E> {
		OtherMethod, // line #100:   static T f<T>(T x) => null;
		OtherMethod, // line #101: }
		OtherMethod, // line #102: ''');
		OtherMethod, // line #103:     expectFunctionType('f', 'T Function<T>(T)', typeFormals: '[T]');
		OtherMethod, // line #104:     SimpleIdentifier f = findNode.simple('f');
		OtherMethod, // line #105:     MethodElementImpl e = f.staticElement;
		OtherMethod, // line #106:     FunctionType ft = e.type.instantiate([typeProvider.stringType]);
		OtherMethod, // line #107:     assertType(ft, 'String Function(String)');
		OtherMethod, // line #108:   }
		BlankLine,   // line #109:
		OtherMethod, // line #110:   test_genericFunction_typedef() async {
		OtherMethod, // line #111:     String code = r'''
		OtherMethod, // line #112: typedef T F<T>(T x);
		OtherMethod, // line #113: F f0;
		OtherMethod, // line #114:
		OtherMethod, // line #115: class C {
		OtherMethod, // line #116:   static F f1;
		OtherMethod, // line #117:   F f2;
		OtherMethod, // line #118:   void g(F f3) { // C
		OtherMethod, // line #119:     F f4;
		OtherMethod, // line #120:     f0(3);
		OtherMethod, // line #121:     f1(3);
		OtherMethod, // line #122:     f2(3);
		OtherMethod, // line #123:     f3(3);
		OtherMethod, // line #124:     f4(3);
		OtherMethod, // line #125:   }
		OtherMethod, // line #126: }
		OtherMethod, // line #127:
		OtherMethod, // line #128: class D<S> {
		OtherMethod, // line #129:   static F f1;
		OtherMethod, // line #130:   F f2;
		OtherMethod, // line #131:   void g(F f3) { // D
		OtherMethod, // line #132:     F f4;
		OtherMethod, // line #133:     f0(3);
		OtherMethod, // line #134:     f1(3);
		OtherMethod, // line #135:     f2(3);
		OtherMethod, // line #136:     f3(3);
		OtherMethod, // line #137:     f4(3);
		OtherMethod, // line #138:   }
		OtherMethod, // line #139: }
		OtherMethod, // line #140: ''';
		OtherMethod, // line #141:     await assertNoErrorsInCode(code);
		OtherMethod, // line #142:
		OtherMethod, // line #143:     checkBody(String className) {
		OtherMethod, // line #144:       var statements = findNode.block('{ // $className').statements;
		OtherMethod, // line #145:
		OtherMethod, // line #146:       for (int i = 1; i <= 5; i++) {
		OtherMethod, // line #147:         Expression exp = (statements[i] as ExpressionStatement).expression;
		OtherMethod, // line #148:         expect(exp.staticType, typeProvider.dynamicType);
		OtherMethod, // line #149:       }
		OtherMethod, // line #150:     }
		OtherMethod, // line #151:
		OtherMethod, // line #152:     checkBody("C");
		OtherMethod, // line #153:     checkBody("D");
		OtherMethod, // line #154:   }
		BlankLine,   // line #155:
		OtherMethod, // line #156:   test_genericFunction_upwardsAndDownwards() async {
		OtherMethod, // line #157:     // Regression tests for https://github.com/dart-lang/sdk/issues/27586.
		OtherMethod, // line #158:     await assertNoErrorsInCode(r'List<num> x = [1, 2];');
		OtherMethod, // line #159:     expectInitializerType('x', 'List<num>');
		OtherMethod, // line #160:   }
		BlankLine,   // line #161:
		OtherMethod, // line #162:   test_genericFunction_upwardsAndDownwards_Object() async {
		OtherMethod, // line #163:     // Regression tests for https://github.com/dart-lang/sdk/issues/27625.
		OtherMethod, // line #164:     await assertNoErrorsInCode(r'''
		OtherMethod, // line #165: List<Object> aaa = [];
		OtherMethod, // line #166: List<Object> bbb = [1, 2, 3];
		OtherMethod, // line #167: List<Object> ccc = [null];
		OtherMethod, // line #168: List<Object> ddd = [1 as dynamic];
		OtherMethod, // line #169: List<Object> eee = [new Object()];
		OtherMethod, // line #170:     ''');
		OtherMethod, // line #171:     expectInitializerType('aaa', 'List<Object>');
		OtherMethod, // line #172:     expectInitializerType('bbb', 'List<Object>');
		OtherMethod, // line #173:     expectInitializerType('ccc', 'List<Object>');
		OtherMethod, // line #174:     expectInitializerType('ddd', 'List<Object>');
		OtherMethod, // line #175:     expectInitializerType('eee', 'List<Object>');
		OtherMethod, // line #176:   }
		BlankLine,   // line #177:
		OtherMethod, // line #178:   test_genericMethod() async {
		OtherMethod, // line #179:     await assertErrorsInCode(r'''
		OtherMethod, // line #180: class C<E> {
		OtherMethod, // line #181:   List<T> f<T>(E e) => null;
		OtherMethod, // line #182: }
		OtherMethod, // line #183: main() {
		OtherMethod, // line #184:   C<String> cOfString;
		OtherMethod, // line #185: }
		OtherMethod, // line #186: ''', [
		OtherMethod, // line #187:       error(HintCode.UNUSED_LOCAL_VARIABLE, 65, 9),
		OtherMethod, // line #188:     ]);
		OtherMethod, // line #189:     assertType(findElement.method('f').type, 'List<T> Function<T>(E)');
		OtherMethod, // line #190:
		OtherMethod, // line #191:     var cOfString = findElement.localVar('cOfString');
		OtherMethod, // line #192:     var ft = (cOfString.type as InterfaceType).getMethod('f').type;
		OtherMethod, // line #193:     assertType(ft, 'List<T> Function<T>(String)');
		OtherMethod, // line #194:     assertType(
		OtherMethod, // line #195:         ft.instantiate([typeProvider.intType]), 'List<int> Function(String)');
		OtherMethod, // line #196:   }
		BlankLine,   // line #197:
		OtherMethod, // line #198:   test_genericMethod_explicitTypeParams() async {
		OtherMethod, // line #199:     await assertErrorsInCode(r'''
		OtherMethod, // line #200: class C<E> {
		OtherMethod, // line #201:   List<T> f<T>(E e) => null;
		OtherMethod, // line #202: }
		OtherMethod, // line #203: main() {
		OtherMethod, // line #204:   C<String> cOfString;
		OtherMethod, // line #205:   var x = cOfString.f<int>('hi');
		OtherMethod, // line #206: }
		OtherMethod, // line #207: ''', [
		OtherMethod, // line #208:       error(HintCode.UNUSED_LOCAL_VARIABLE, 82, 1),
		OtherMethod, // line #209:     ]);
		OtherMethod, // line #210:     MethodInvocation f = findNode.simple('f<int>').parent;
		OtherMethod, // line #211:     FunctionType ft = f.staticInvokeType;
		OtherMethod, // line #212:     assertType(ft, 'List<int> Function(String)');
		OtherMethod, // line #213:
		OtherMethod, // line #214:     var x = findElement.localVar('x');
		OtherMethod, // line #215:     expect(x.type, typeProvider.listType2(typeProvider.intType));
		OtherMethod, // line #216:   }
		BlankLine,   // line #217:
		OtherMethod, // line #218:   test_genericMethod_functionExpressionInvocation_explicit() async {
		OtherMethod, // line #219:     await assertErrorsInCode(r'''
		OtherMethod, // line #220: class C<E> {
		OtherMethod, // line #221:   T f<T>(T e) => null;
		OtherMethod, // line #222:   static T g<T>(T e) => null;
		OtherMethod, // line #223:   static T Function<T>(T) h = null;
		OtherMethod, // line #224: }
		OtherMethod, // line #225:
		OtherMethod, // line #226: T topF<T>(T e) => null;
		OtherMethod, // line #227: var topG = topF;
		OtherMethod, // line #228: void test<S>(T Function<T>(T) pf) {
		OtherMethod, // line #229:   var c = new C<int>();
		OtherMethod, // line #230:   T lf<T>(T e) => null;
		OtherMethod, // line #231:
		OtherMethod, // line #232:   var lambdaCall = (<E>(E e) => e)<int>(3);
		OtherMethod, // line #233:   var methodCall = (c.f)<int>(3);
		OtherMethod, // line #234:   var staticCall = (C.g)<int>(3);
		OtherMethod, // line #235:   var staticFieldCall = (C.h)<int>(3);
		OtherMethod, // line #236:   var topFunCall = (topF)<int>(3);
		OtherMethod, // line #237:   var topFieldCall = (topG)<int>(3);
		OtherMethod, // line #238:   var localCall = (lf)<int>(3);
		OtherMethod, // line #239:   var paramCall = (pf)<int>(3);
		OtherMethod, // line #240: }
		OtherMethod, // line #241: ''', [
		OtherMethod, // line #242:       error(HintCode.UNUSED_LOCAL_VARIABLE, 237, 10),
		OtherMethod, // line #243:       error(HintCode.UNUSED_LOCAL_VARIABLE, 281, 10),
		OtherMethod, // line #244:       error(HintCode.UNUSED_LOCAL_VARIABLE, 315, 10),
		OtherMethod, // line #245:       error(HintCode.UNUSED_LOCAL_VARIABLE, 349, 15),
		OtherMethod, // line #246:       error(HintCode.UNUSED_LOCAL_VARIABLE, 388, 10),
		OtherMethod, // line #247:       error(HintCode.UNUSED_LOCAL_VARIABLE, 423, 12),
		OtherMethod, // line #248:       error(HintCode.UNUSED_LOCAL_VARIABLE, 460, 9),
		OtherMethod, // line #249:       error(HintCode.UNUSED_LOCAL_VARIABLE, 492, 9),
		OtherMethod, // line #250:     ]);
		OtherMethod, // line #251:     _assertLocalVarType('lambdaCall', "int");
		OtherMethod, // line #252:     _assertLocalVarType('methodCall', "int");
		OtherMethod, // line #253:     _assertLocalVarType('staticCall', "int");
		OtherMethod, // line #254:     _assertLocalVarType('staticFieldCall', "int");
		OtherMethod, // line #255:     _assertLocalVarType('topFunCall', "int");
		OtherMethod, // line #256:     _assertLocalVarType('topFieldCall', "int");
		OtherMethod, // line #257:     _assertLocalVarType('localCall', "int");
		OtherMethod, // line #258:     _assertLocalVarType('paramCall', "int");
		OtherMethod, // line #259:   }
		BlankLine,   // line #260:
		OtherMethod, // line #261:   test_genericMethod_functionExpressionInvocation_functionTypedParameter_explicit() async {
		OtherMethod, // line #262:     await assertErrorsInCode(r'''
		OtherMethod, // line #263: void test<S>(T pf<T>(T e)) {
		OtherMethod, // line #264:   var paramCall = (pf)<int>(3);
		OtherMethod, // line #265: }
		OtherMethod, // line #266: ''', [
		OtherMethod, // line #267:       error(HintCode.UNUSED_LOCAL_VARIABLE, 35, 9),
		OtherMethod, // line #268:     ]);
		OtherMethod, // line #269:     _assertLocalVarType('paramCall', "int");
		OtherMethod, // line #270:   }
		BlankLine,   // line #271:
		OtherMethod, // line #272:   test_genericMethod_functionExpressionInvocation_functionTypedParameter_inferred() async {
		OtherMethod, // line #273:     await assertErrorsInCode(r'''
		OtherMethod, // line #274: void test<S>(T pf<T>(T e)) {
		OtherMethod, // line #275:   var paramCall = (pf)(3);
		OtherMethod, // line #276: }
		OtherMethod, // line #277: ''', [
		OtherMethod, // line #278:       error(HintCode.UNUSED_LOCAL_VARIABLE, 35, 9),
		OtherMethod, // line #279:     ]);
		OtherMethod, // line #280:     _assertLocalVarType('paramCall', "int");
		OtherMethod, // line #281:   }
		BlankLine,   // line #282:
		OtherMethod, // line #283:   test_genericMethod_functionExpressionInvocation_inferred() async {
		OtherMethod, // line #284:     await assertErrorsInCode(r'''
		OtherMethod, // line #285: class C<E> {
		OtherMethod, // line #286:   T f<T>(T e) => null;
		OtherMethod, // line #287:   static T g<T>(T e) => null;
		OtherMethod, // line #288:   static T Function<T>(T) h = null;
		OtherMethod, // line #289: }
		OtherMethod, // line #290:
		OtherMethod, // line #291: T topF<T>(T e) => null;
		OtherMethod, // line #292: var topG = topF;
		OtherMethod, // line #293: void test<S>(T Function<T>(T) pf) {
		OtherMethod, // line #294:   var c = new C<int>();
		OtherMethod, // line #295:   T lf<T>(T e) => null;
		OtherMethod, // line #296:
		OtherMethod, // line #297:   var lambdaCall = (<E>(E e) => e)(3);
		OtherMethod, // line #298:   var methodCall = (c.f)(3);
		OtherMethod, // line #299:   var staticCall = (C.g)(3);
		OtherMethod, // line #300:   var staticFieldCall = (C.h)(3);
		OtherMethod, // line #301:   var topFunCall = (topF)(3);
		OtherMethod, // line #302:   var topFieldCall = (topG)(3);
		OtherMethod, // line #303:   var localCall = (lf)(3);
		OtherMethod, // line #304:   var paramCall = (pf)(3);
		OtherMethod, // line #305: }
		OtherMethod, // line #306: ''', [
		OtherMethod, // line #307:       error(HintCode.UNUSED_LOCAL_VARIABLE, 237, 10),
		OtherMethod, // line #308:       error(HintCode.UNUSED_LOCAL_VARIABLE, 276, 10),
		OtherMethod, // line #309:       error(HintCode.UNUSED_LOCAL_VARIABLE, 305, 10),
		OtherMethod, // line #310:       error(HintCode.UNUSED_LOCAL_VARIABLE, 334, 15),
		OtherMethod, // line #311:       error(HintCode.UNUSED_LOCAL_VARIABLE, 368, 10),
		OtherMethod, // line #312:       error(HintCode.UNUSED_LOCAL_VARIABLE, 398, 12),
		OtherMethod, // line #313:       error(HintCode.UNUSED_LOCAL_VARIABLE, 430, 9),
		OtherMethod, // line #314:       error(HintCode.UNUSED_LOCAL_VARIABLE, 457, 9),
		OtherMethod, // line #315:     ]);
		OtherMethod, // line #316:     _assertLocalVarType('lambdaCall', "int");
		OtherMethod, // line #317:     _assertLocalVarType('methodCall', "int");
		OtherMethod, // line #318:     _assertLocalVarType('staticCall', "int");
		OtherMethod, // line #319:     _assertLocalVarType('staticFieldCall', "int");
		OtherMethod, // line #320:     _assertLocalVarType('topFunCall', "int");
		OtherMethod, // line #321:     _assertLocalVarType('topFieldCall', "int");
		OtherMethod, // line #322:     _assertLocalVarType('localCall', "int");
		OtherMethod, // line #323:     _assertLocalVarType('paramCall', "int");
		OtherMethod, // line #324:   }
		BlankLine,   // line #325:
		OtherMethod, // line #326:   test_genericMethod_functionInvocation_explicit() async {
		OtherMethod, // line #327:     await assertErrorsInCode(r'''
		OtherMethod, // line #328: class C<E> {
		OtherMethod, // line #329:   T f<T>(T e) => null;
		OtherMethod, // line #330:   static T g<T>(T e) => null;
		OtherMethod, // line #331:   static T Function<T>(T) h = null;
		OtherMethod, // line #332: }
		OtherMethod, // line #333:
		OtherMethod, // line #334: T topF<T>(T e) => null;
		OtherMethod, // line #335: var topG = topF;
		OtherMethod, // line #336: void test<S>(T Function<T>(T) pf) {
		OtherMethod, // line #337:   var c = new C<int>();
		OtherMethod, // line #338:   T lf<T>(T e) => null;
		OtherMethod, // line #339:   var methodCall = c.f<int>(3);
		OtherMethod, // line #340:   var staticCall = C.g<int>(3);
		OtherMethod, // line #341:   var staticFieldCall = C.h<int>(3);
		OtherMethod, // line #342:   var topFunCall = topF<int>(3);
		OtherMethod, // line #343:   var topFieldCall = topG<int>(3);
		OtherMethod, // line #344:   var localCall = lf<int>(3);
		OtherMethod, // line #345:   var paramCall = pf<int>(3);
		OtherMethod, // line #346: }
		OtherMethod, // line #347: ''', [
		OtherMethod, // line #348:       error(HintCode.UNUSED_LOCAL_VARIABLE, 236, 10),
		OtherMethod, // line #349:       error(HintCode.UNUSED_LOCAL_VARIABLE, 268, 10),
		OtherMethod, // line #350:       error(HintCode.UNUSED_LOCAL_VARIABLE, 300, 15),
		OtherMethod, // line #351:       error(HintCode.UNUSED_LOCAL_VARIABLE, 337, 10),
		OtherMethod, // line #352:       error(HintCode.UNUSED_LOCAL_VARIABLE, 370, 12),
		OtherMethod, // line #353:       error(HintCode.UNUSED_LOCAL_VARIABLE, 405, 9),
		OtherMethod, // line #354:       error(HintCode.UNUSED_LOCAL_VARIABLE, 435, 9),
		OtherMethod, // line #355:     ]);
		OtherMethod, // line #356:     _assertLocalVarType('methodCall', "int");
		OtherMethod, // line #357:     _assertLocalVarType('staticCall', "int");
		OtherMethod, // line #358:     _assertLocalVarType('staticFieldCall', "int");
		OtherMethod, // line #359:     _assertLocalVarType('topFunCall', "int");
		OtherMethod, // line #360:     _assertLocalVarType('topFieldCall', "int");
		OtherMethod, // line #361:     _assertLocalVarType('localCall', "int");
		OtherMethod, // line #362:     _assertLocalVarType('paramCall', "int");
		OtherMethod, // line #363:   }
		BlankLine,   // line #364:
		OtherMethod, // line #365:   test_genericMethod_functionInvocation_functionTypedParameter_explicit() async {
		OtherMethod, // line #366:     await assertErrorsInCode(r'''
		OtherMethod, // line #367: void test<S>(T pf<T>(T e)) {
		OtherMethod, // line #368:   var paramCall = pf<int>(3);
		OtherMethod, // line #369: }
		OtherMethod, // line #370: ''', [
		OtherMethod, // line #371:       error(HintCode.UNUSED_LOCAL_VARIABLE, 35, 9),
		OtherMethod, // line #372:     ]);
		OtherMethod, // line #373:     _assertLocalVarType('paramCall', "int");
		OtherMethod, // line #374:   }
		BlankLine,   // line #375:
		OtherMethod, // line #376:   test_genericMethod_functionInvocation_functionTypedParameter_inferred() async {
		OtherMethod, // line #377:     await assertErrorsInCode(r'''
		OtherMethod, // line #378: void test<S>(T pf<T>(T e)) {
		OtherMethod, // line #379:   var paramCall = pf(3);
		OtherMethod, // line #380: }
		OtherMethod, // line #381: ''', [
		OtherMethod, // line #382:       error(HintCode.UNUSED_LOCAL_VARIABLE, 35, 9),
		OtherMethod, // line #383:     ]);
		OtherMethod, // line #384:     _assertLocalVarType('paramCall', "int");
		OtherMethod, // line #385:   }
		BlankLine,   // line #386:
		OtherMethod, // line #387:   test_genericMethod_functionInvocation_inferred() async {
		OtherMethod, // line #388:     await assertErrorsInCode(r'''
		OtherMethod, // line #389: class C<E> {
		OtherMethod, // line #390:   T f<T>(T e) => null;
		OtherMethod, // line #391:   static T g<T>(T e) => null;
		OtherMethod, // line #392:   static T Function<T>(T) h = null;
		OtherMethod, // line #393: }
		OtherMethod, // line #394:
		OtherMethod, // line #395: T topF<T>(T e) => null;
		OtherMethod, // line #396: var topG = topF;
		OtherMethod, // line #397: void test<S>(T Function<T>(T) pf) {
		OtherMethod, // line #398:   var c = new C<int>();
		OtherMethod, // line #399:   T lf<T>(T e) => null;
		OtherMethod, // line #400:   var methodCall = c.f(3);
		OtherMethod, // line #401:   var staticCall = C.g(3);
		OtherMethod, // line #402:   var staticFieldCall = C.h(3);
		OtherMethod, // line #403:   var topFunCall = topF(3);
		OtherMethod, // line #404:   var topFieldCall = topG(3);
		OtherMethod, // line #405:   var localCall = lf(3);
		OtherMethod, // line #406:   var paramCall = pf(3);
		OtherMethod, // line #407: }
		OtherMethod, // line #408: ''', [
		OtherMethod, // line #409:       error(HintCode.UNUSED_LOCAL_VARIABLE, 236, 10),
		OtherMethod, // line #410:       error(HintCode.UNUSED_LOCAL_VARIABLE, 263, 10),
		OtherMethod, // line #411:       error(HintCode.UNUSED_LOCAL_VARIABLE, 290, 15),
		OtherMethod, // line #412:       error(HintCode.UNUSED_LOCAL_VARIABLE, 322, 10),
		OtherMethod, // line #413:       error(HintCode.UNUSED_LOCAL_VARIABLE, 350, 12),
		OtherMethod, // line #414:       error(HintCode.UNUSED_LOCAL_VARIABLE, 380, 9),
		OtherMethod, // line #415:       error(HintCode.UNUSED_LOCAL_VARIABLE, 405, 9),
		OtherMethod, // line #416:     ]);
		OtherMethod, // line #417:     _assertLocalVarType('methodCall', "int");
		OtherMethod, // line #418:     _assertLocalVarType('staticCall', "int");
		OtherMethod, // line #419:     _assertLocalVarType('staticFieldCall', "int");
		OtherMethod, // line #420:     _assertLocalVarType('topFunCall', "int");
		OtherMethod, // line #421:     _assertLocalVarType('topFieldCall', "int");
		OtherMethod, // line #422:     _assertLocalVarType('localCall', "int");
		OtherMethod, // line #423:     _assertLocalVarType('paramCall', "int");
		OtherMethod, // line #424:   }
		BlankLine,   // line #425:
		OtherMethod, // line #426:   test_genericMethod_functionTypedParameter() async {
		OtherMethod, // line #427:     await assertErrorsInCode(r'''
		OtherMethod, // line #428: class C<E> {
		OtherMethod, // line #429:   List<T> f<T>(T f(E e)) => null;
		OtherMethod, // line #430: }
		OtherMethod, // line #431: main() {
		OtherMethod, // line #432:   C<String> cOfString;
		OtherMethod, // line #433: }
		OtherMethod, // line #434: ''', [
		OtherMethod, // line #435:       error(HintCode.UNUSED_LOCAL_VARIABLE, 70, 9),
		OtherMethod, // line #436:     ]);
		OtherMethod, // line #437:     assertType(
		OtherMethod, // line #438:         findElement.method('f').type, 'List<T> Function<T>(T Function(E))');
		OtherMethod, // line #439:
		OtherMethod, // line #440:     var cOfString = findElement.localVar('cOfString');
		OtherMethod, // line #441:     var ft = (cOfString.type as InterfaceType).getMethod('f').type;
		OtherMethod, // line #442:     assertType(ft, 'List<T> Function<T>(T Function(String))');
		OtherMethod, // line #443:     assertType(ft.instantiate([typeProvider.intType]),
		OtherMethod, // line #444:         'List<int> Function(int Function(String))');
		OtherMethod, // line #445:   }
		BlankLine,   // line #446:
		OtherMethod, // line #447:   test_genericMethod_functionTypedParameter_tearoff() async {
		OtherMethod, // line #448:     await assertErrorsInCode(r'''
		OtherMethod, // line #449: void test<S>(T pf<T>(T e)) {
		OtherMethod, // line #450:   var paramTearOff = pf;
		OtherMethod, // line #451: }
		OtherMethod, // line #452: ''', [
		OtherMethod, // line #453:       error(HintCode.UNUSED_LOCAL_VARIABLE, 35, 12),
		OtherMethod, // line #454:     ]);
		OtherMethod, // line #455:     _assertLocalVarType('paramTearOff', "T Function<T>(T)");
		OtherMethod, // line #456:   }
		BlankLine,   // line #457:
		OtherMethod, // line #458:   test_genericMethod_implicitDynamic() async {
		OtherMethod, // line #459:     // Regression test for:
		OtherMethod, // line #460:     // https://github.com/dart-lang/sdk/issues/25100#issuecomment-162047588
		OtherMethod, // line #461:     // These should not cause any hints or warnings.
		OtherMethod, // line #462:     await assertNoErrorsInCode(r'''
		OtherMethod, // line #463: class List<E> {
		OtherMethod, // line #464:   T map<T>(T f(E e)) => null;
		OtherMethod, // line #465: }
		OtherMethod, // line #466: void foo() {
		OtherMethod, // line #467:   List list = null;
		OtherMethod, // line #468:   list.map((e) => e);
		OtherMethod, // line #469:   list.map((e) => 3);
		OtherMethod, // line #470: }''');
		OtherMethod, // line #471:     expectIdentifierType(
		OtherMethod, // line #472:         'map((e) => e);', 'T Function<T>(T Function(dynamic))');
		OtherMethod, // line #473:     expectIdentifierType(
		OtherMethod, // line #474:         'map((e) => 3);', 'T Function<T>(T Function(dynamic))');
		OtherMethod, // line #475:
		OtherMethod, // line #476:     MethodInvocation m1 = findNode.methodInvocation('map((e) => e);');
		OtherMethod, // line #477:     assertInvokeType(m1, 'dynamic Function(dynamic Function(dynamic))');
		OtherMethod, // line #478:     MethodInvocation m2 = findNode.methodInvocation('map((e) => 3);');
		OtherMethod, // line #479:     assertInvokeType(m2, 'int Function(int Function(dynamic))');
		OtherMethod, // line #480:   }
		BlankLine,   // line #481:
		OtherMethod, // line #482:   test_genericMethod_max_doubleDouble() async {
		OtherMethod, // line #483:     await assertErrorsInCode(r'''
		OtherMethod, // line #484: import 'dart:math';
		OtherMethod, // line #485: main() {
		OtherMethod, // line #486:   var foo = max(1.0, 2.0);
		OtherMethod, // line #487: }
		OtherMethod, // line #488: ''', [
		OtherMethod, // line #489:       error(HintCode.UNUSED_LOCAL_VARIABLE, 35, 3),
		OtherMethod, // line #490:     ]);
		OtherMethod, // line #491:     expectInitializerType('foo', 'double');
		OtherMethod, // line #492:   }
		BlankLine,   // line #493:
		OtherMethod, // line #494:   test_genericMethod_max_doubleDouble_prefixed() async {
		OtherMethod, // line #495:     await assertErrorsInCode(r'''
		OtherMethod, // line #496: import 'dart:math' as math;
		OtherMethod, // line #497: main() {
		OtherMethod, // line #498:   var foo = math.max(1.0, 2.0);
		OtherMethod, // line #499: }
		OtherMethod, // line #500: ''', [
		OtherMethod, // line #501:       error(HintCode.UNUSED_LOCAL_VARIABLE, 43, 3),
		OtherMethod, // line #502:     ]);
		OtherMethod, // line #503:     expectInitializerType('foo', 'double');
		OtherMethod, // line #504:   }
		BlankLine,   // line #505:
		OtherMethod, // line #506:   test_genericMethod_max_doubleInt() async {
		OtherMethod, // line #507:     await assertErrorsInCode(r'''
		OtherMethod, // line #508: import 'dart:math';
		OtherMethod, // line #509: main() {
		OtherMethod, // line #510:   var foo = max(1.0, 2);
		OtherMethod, // line #511: }
		OtherMethod, // line #512: ''', [
		OtherMethod, // line #513:       error(HintCode.UNUSED_LOCAL_VARIABLE, 35, 3),
		OtherMethod, // line #514:     ]);
		OtherMethod, // line #515:     expectInitializerType('foo', 'num');
		OtherMethod, // line #516:   }
		BlankLine,   // line #517:
		OtherMethod, // line #518:   test_genericMethod_max_intDouble() async {
		OtherMethod, // line #519:     await assertErrorsInCode(r'''
		OtherMethod, // line #520: import 'dart:math';
		OtherMethod, // line #521: main() {
		OtherMethod, // line #522:   var foo = max(1, 2.0);
		OtherMethod, // line #523: }
		OtherMethod, // line #524: ''', [
		OtherMethod, // line #525:       error(HintCode.UNUSED_LOCAL_VARIABLE, 35, 3),
		OtherMethod, // line #526:     ]);
		OtherMethod, // line #527:     expectInitializerType('foo', 'num');
		OtherMethod, // line #528:   }
		BlankLine,   // line #529:
		OtherMethod, // line #530:   test_genericMethod_max_intInt() async {
		OtherMethod, // line #531:     await assertErrorsInCode(r'''
		OtherMethod, // line #532: import 'dart:math';
		OtherMethod, // line #533: main() {
		OtherMethod, // line #534:   var foo = max(1, 2);
		OtherMethod, // line #535: }
		OtherMethod, // line #536: ''', [
		OtherMethod, // line #537:       error(HintCode.UNUSED_LOCAL_VARIABLE, 35, 3),
		OtherMethod, // line #538:     ]);
		OtherMethod, // line #539:     expectInitializerType('foo', 'int');
		OtherMethod, // line #540:   }
		BlankLine,   // line #541:
		OtherMethod, // line #542:   test_genericMethod_nestedBound() async {
		OtherMethod, // line #543:     // Just validate that there is no warning on the call to `.abs()`.
		OtherMethod, // line #544:     await assertNoErrorsInCode(r'''
		OtherMethod, // line #545: class Foo<T extends num> {
		OtherMethod, // line #546:   void method<U extends T>(U u) {
		OtherMethod, // line #547:     u.abs();
		OtherMethod, // line #548:   }
		OtherMethod, // line #549: }
		OtherMethod, // line #550: ''');
		OtherMethod, // line #551:   }
		BlankLine,   // line #552:
		OtherMethod, // line #553:   test_genericMethod_nestedCapture() async {
		OtherMethod, // line #554:     await assertNoErrorsInCode(r'''
		OtherMethod, // line #555: class C<T> {
		OtherMethod, // line #556:   T f<S>(S x) {
		OtherMethod, // line #557:     new C<S>().f<int>(3);
		OtherMethod, // line #558:     new C<S>().f; // tear-off
		OtherMethod, // line #559:     return null;
		OtherMethod, // line #560:   }
		OtherMethod, // line #561: }
		OtherMethod, // line #562: ''');
		OtherMethod, // line #563:     MethodInvocation f = findNode.methodInvocation('f<int>(3);');
		OtherMethod, // line #564:     assertInvokeType(f, 'S Function(int)');
		OtherMethod, // line #565:
		OtherMethod, // line #566:     expectIdentifierType('f;', 'S Function<S₀>(S₀)');
		OtherMethod, // line #567:   }
		BlankLine,   // line #568:
		OtherMethod, // line #569:   @FailingTest(issue: 'https://github.com/dart-lang/sdk/issues/30236')
		OtherMethod, // line #570:   test_genericMethod_nestedCaptureBounds() async {
		OtherMethod, // line #571:     await assertNoErrorsInCode(r'''
		OtherMethod, // line #572: class C<T> {
		OtherMethod, // line #573:   T f<S extends T>(S x) {
		OtherMethod, // line #574:     new C<S>().f<int>(3);
		OtherMethod, // line #575:     new C<S>().f; // tear-off
		OtherMethod, // line #576:     return null;
		OtherMethod, // line #577:   }
		OtherMethod, // line #578: }
		OtherMethod, // line #579: ''');
		OtherMethod, // line #580:     MethodInvocation f = findNode.methodInvocation('f<int>(3);');
		OtherMethod, // line #581:     assertInvokeType(f, 'S Function(int)');
		OtherMethod, // line #582:     FunctionType ft = f.staticInvokeType;
		OtherMethod, // line #583:     expect('${ft.typeArguments}', '[S, int]');
		OtherMethod, // line #584:
		OtherMethod, // line #585:     expectIdentifierType('f;', 'S Function<S₀ extends S>(S₀)');
		OtherMethod, // line #586:   }
		BlankLine,   // line #587:
		OtherMethod, // line #588:   test_genericMethod_nestedFunctions() async {
		OtherMethod, // line #589:     await assertErrorsInCode(r'''
		OtherMethod, // line #590: S f<S>(S x) {
		OtherMethod, // line #591:   g<S>(S x) => f;
		OtherMethod, // line #592:   return null;
		OtherMethod, // line #593: }
		OtherMethod, // line #594: ''', [
		OtherMethod, // line #595:       error(HintCode.UNUSED_ELEMENT, 16, 1),
		OtherMethod, // line #596:     ]);
		OtherMethod, // line #597:     assertType(findElement.topFunction('f').type, 'S Function<S>(S)');
		OtherMethod, // line #598:     assertType(findElement.localFunction('g').type,
		OtherMethod, // line #599:         'S Function<S>(S) Function<S₀>(S₀)');
		OtherMethod, // line #600:   }
		BlankLine,   // line #601:
		OtherMethod, // line #602:   test_genericMethod_override() async {
		OtherMethod, // line #603:     await assertNoErrorsInCode(r'''
		OtherMethod, // line #604: class C {
		OtherMethod, // line #605:   T f<T>(T x) => null;
		OtherMethod, // line #606: }
		OtherMethod, // line #607: class D extends C {
		OtherMethod, // line #608:   T f<T>(T x) => null; // from D
		OtherMethod, // line #609: }
		OtherMethod, // line #610: ''');
		OtherMethod, // line #611:     expectFunctionType('f<T>(T x) => null; // from D', 'T Function<T>(T)',
		OtherMethod, // line #612:         typeFormals: '[T]');
		OtherMethod, // line #613:     SimpleIdentifier f = findNode.simple('f<T>(T x) => null; // from D');
		OtherMethod, // line #614:     MethodElementImpl e = f.staticElement;
		OtherMethod, // line #615:     FunctionType ft = e.type.instantiate([typeProvider.stringType]);
		OtherMethod, // line #616:     assertType(ft, 'String Function(String)');
		OtherMethod, // line #617:   }
		BlankLine,   // line #618:
		OtherMethod, // line #619:   test_genericMethod_override_bounds() async {
		OtherMethod, // line #620:     await assertNoErrorsInCode(r'''
		OtherMethod, // line #621: class A {}
		OtherMethod, // line #622: class B {
		OtherMethod, // line #623:   T f<T extends A>(T x) => null;
		OtherMethod, // line #624: }
		OtherMethod, // line #625: // override with the same bound is OK
		OtherMethod, // line #626: class C extends B {
		OtherMethod, // line #627:   T f<T extends A>(T x) => null;
		OtherMethod, // line #628: }
		OtherMethod, // line #629: // override with new name and the same bound is OK
		OtherMethod, // line #630: class D extends B {
		OtherMethod, // line #631:   Q f<Q extends A>(Q x) => null;
		OtherMethod, // line #632: }
		OtherMethod, // line #633: ''');
		OtherMethod, // line #634:   }
		BlankLine,   // line #635:
		OtherMethod, // line #636:   test_genericMethod_override_covariant_field() async {
		OtherMethod, // line #637:     await assertNoErrorsInCode(r'''
		OtherMethod, // line #638: abstract class A {
		OtherMethod, // line #639:   num get x;
		OtherMethod, // line #640:   set x(covariant num _);
		OtherMethod, // line #641: }
		OtherMethod, // line #642:
		OtherMethod, // line #643: class B extends A {
		OtherMethod, // line #644:   int x;
		OtherMethod, // line #645: }
		OtherMethod, // line #646: ''');
		OtherMethod, // line #647:   }
		BlankLine,   // line #648:
		OtherMethod, // line #649:   test_genericMethod_override_differentContextsSameBounds() async {
		OtherMethod, // line #650:     await assertNoErrorsInCode(r'''
		OtherMethod, // line #651:         class GenericMethodBounds<T> {
		OtherMethod, // line #652:   Type get t => T;
		OtherMethod, // line #653:   GenericMethodBounds<E> foo<E extends T>() => new GenericMethodBounds<E>();
		OtherMethod, // line #654:   GenericMethodBounds<E> bar<E extends void Function(T)>() =>
		OtherMethod, // line #655:       new GenericMethodBounds<E>();
		OtherMethod, // line #656: }
		OtherMethod, // line #657:
		OtherMethod, // line #658: class GenericMethodBoundsDerived extends GenericMethodBounds<num> {
		OtherMethod, // line #659:   GenericMethodBounds<E> foo<E extends num>() => new GenericMethodBounds<E>();
		OtherMethod, // line #660:   GenericMethodBounds<E> bar<E extends void Function(num)>() =>
		OtherMethod, // line #661:       new GenericMethodBounds<E>();
		OtherMethod, // line #662: }
		OtherMethod, // line #663: ''');
		OtherMethod, // line #664:   }
		BlankLine,   // line #665:
		OtherMethod, // line #666:   test_genericMethod_override_invalidContravariantTypeParamBounds() async {
		OtherMethod, // line #667:     await assertErrorsInCode(r'''
		OtherMethod, // line #668: class A {}
		OtherMethod, // line #669: class B extends A {}
		OtherMethod, // line #670: class C {
		OtherMethod, // line #671:   T f<T extends A>(T x) => null;
		OtherMethod, // line #672: }
		OtherMethod, // line #673: class D extends C {
		OtherMethod, // line #674:   T f<T extends B>(T x) => null;
		OtherMethod, // line #675: }''', [
		OtherMethod, // line #676:       error(CompileTimeErrorCode.INVALID_OVERRIDE, 101, 1),
		OtherMethod, // line #677:     ]);
		OtherMethod, // line #678:   }
		BlankLine,   // line #679:
		OtherMethod, // line #680:   test_genericMethod_override_invalidCovariantTypeParamBounds() async {
		OtherMethod, // line #681:     await assertErrorsInCode(r'''
		OtherMethod, // line #682: class A {}
		OtherMethod, // line #683: class B extends A {}
		OtherMethod, // line #684: class C {
		OtherMethod, // line #685:   T f<T extends B>(T x) => null;
		OtherMethod, // line #686: }
		OtherMethod, // line #687: class D extends C {
		OtherMethod, // line #688:   T f<T extends A>(T x) => null;
		OtherMethod, // line #689: }''', [
		OtherMethod, // line #690:       error(CompileTimeErrorCode.INVALID_OVERRIDE, 101, 1),
		OtherMethod, // line #691:     ]);
		OtherMethod, // line #692:   }
		BlankLine,   // line #693:
		OtherMethod, // line #694:   test_genericMethod_override_invalidReturnType() async {
		OtherMethod, // line #695:     await assertErrorsInCode(r'''
		OtherMethod, // line #696: class C {
		OtherMethod, // line #697:   Iterable<T> f<T>(T x) => null;
		OtherMethod, // line #698: }
		OtherMethod, // line #699: class D extends C {
		OtherMethod, // line #700:   String f<S>(S x) => null;
		OtherMethod, // line #701: }''', [
		OtherMethod, // line #702:       error(CompileTimeErrorCode.INVALID_OVERRIDE, 74, 1),
		OtherMethod, // line #703:     ]);
		OtherMethod, // line #704:   }
		BlankLine,   // line #705:
		OtherMethod, // line #706:   test_genericMethod_override_invalidTypeParamCount() async {
		OtherMethod, // line #707:     await assertErrorsInCode(r'''
		OtherMethod, // line #708: class C {
		OtherMethod, // line #709:   T f<T>(T x) => null;
		OtherMethod, // line #710: }
		OtherMethod, // line #711: class D extends C {
		OtherMethod, // line #712:   S f<T, S>(T x) => null;
		OtherMethod, // line #713: }''', [
		OtherMethod, // line #714:       error(CompileTimeErrorCode.INVALID_OVERRIDE, 59, 1),
		OtherMethod, // line #715:     ]);
		OtherMethod, // line #716:   }
		BlankLine,   // line #717:
		OtherMethod, // line #718:   test_genericMethod_propagatedType_promotion() async {
		OtherMethod, // line #719:     // Regression test for:
		OtherMethod, // line #720:     // https://github.com/dart-lang/sdk/issues/25340
		OtherMethod, // line #721:
		OtherMethod, // line #722:     // Note, after https://github.com/dart-lang/sdk/issues/25486 the original
		OtherMethod, // line #723:     // example won't work, as we now compute a static type and therefore discard
		OtherMethod, // line #724:     // the propagated type. So a new test was created that doesn't run under
		OtherMethod, // line #725:     // strong mode.
		OtherMethod, // line #726:     await assertErrorsInCode(r'''
		OtherMethod, // line #727: abstract class Iter {
		OtherMethod, // line #728:   List<S> map<S>(S f(x));
		OtherMethod, // line #729: }
		OtherMethod, // line #730: class C {}
		OtherMethod, // line #731: C toSpan(dynamic element) {
		OtherMethod, // line #732:   if (element is Iter) {
		OtherMethod, // line #733:     var y = element.map(toSpan);
		OtherMethod, // line #734:   }
		OtherMethod, // line #735:   return null;
		OtherMethod, // line #736: }
		OtherMethod, // line #737: ''', [
		OtherMethod, // line #738:       error(HintCode.UNUSED_LOCAL_VARIABLE, 122, 1),
		OtherMethod, // line #739:     ]);
		OtherMethod, // line #740:     _assertLocalVarType('y', 'List<C>');
		OtherMethod, // line #741:   }
		BlankLine,   // line #742:
		OtherMethod, // line #743:   test_genericMethod_tearoff() async {
		OtherMethod, // line #744:     await assertErrorsInCode(r'''
		OtherMethod, // line #745: class C<E> {
		OtherMethod, // line #746:   T f<T>(E e) => null;
		OtherMethod, // line #747:   static T g<T>(T e) => null;
		OtherMethod, // line #748:   static T Function<T>(T) h = null;
		OtherMethod, // line #749: }
		OtherMethod, // line #750:
		OtherMethod, // line #751: T topF<T>(T e) => null;
		OtherMethod, // line #752: var topG = topF;
		OtherMethod, // line #753: void test<S>(T Function<T>(T) pf) {
		OtherMethod, // line #754:   var c = new C<int>();
		OtherMethod, // line #755:   T lf<T>(T e) => null;
		OtherMethod, // line #756:   var methodTearOff = c.f;
		OtherMethod, // line #757:   var staticTearOff = C.g;
		OtherMethod, // line #758:   var staticFieldTearOff = C.h;
		OtherMethod, // line #759:   var topFunTearOff = topF;
		OtherMethod, // line #760:   var topFieldTearOff = topG;
		OtherMethod, // line #761:   var localTearOff = lf;
		OtherMethod, // line #762:   var paramTearOff = pf;
		OtherMethod, // line #763: }
		OtherMethod, // line #764: ''', [
		OtherMethod, // line #765:       error(HintCode.UNUSED_LOCAL_VARIABLE, 236, 13),
		OtherMethod, // line #766:       error(HintCode.UNUSED_LOCAL_VARIABLE, 263, 13),
		OtherMethod, // line #767:       error(HintCode.UNUSED_LOCAL_VARIABLE, 290, 18),
		OtherMethod, // line #768:       error(HintCode.UNUSED_LOCAL_VARIABLE, 322, 13),
		OtherMethod, // line #769:       error(HintCode.UNUSED_LOCAL_VARIABLE, 350, 15),
		OtherMethod, // line #770:       error(HintCode.UNUSED_LOCAL_VARIABLE, 380, 12),
		OtherMethod, // line #771:       error(HintCode.UNUSED_LOCAL_VARIABLE, 405, 12),
		OtherMethod, // line #772:     ]);
		OtherMethod, // line #773:     _assertLocalVarType('methodTearOff', "T Function<T>(int)");
		OtherMethod, // line #774:     _assertLocalVarType('staticTearOff', "T Function<T>(T)");
		OtherMethod, // line #775:     _assertLocalVarType('staticFieldTearOff', "T Function<T>(T)");
		OtherMethod, // line #776:     _assertLocalVarType('topFunTearOff', "T Function<T>(T)");
		OtherMethod, // line #777:     _assertLocalVarType('topFieldTearOff', "T Function<T>(T)");
		OtherMethod, // line #778:     _assertLocalVarType('localTearOff', "T Function<T>(T)");
		OtherMethod, // line #779:     _assertLocalVarType('paramTearOff', "T Function<T>(T)");
		OtherMethod, // line #780:   }
		BlankLine,   // line #781:
		OtherMethod, // line #782:   @failingTest
		OtherMethod, // line #783:   test_genericMethod_tearoff_instantiated() async {
		OtherMethod, // line #784:     await assertNoErrorsInCode(r'''
		OtherMethod, // line #785: class C<E> {
		OtherMethod, // line #786:   T f<T>(E e) => null;
		OtherMethod, // line #787:   static T g<T>(T e) => null;
		OtherMethod, // line #788:   static T Function<T>(T) h = null;
		OtherMethod, // line #789: }
		OtherMethod, // line #790:
		OtherMethod, // line #791: T topF<T>(T e) => null;
		OtherMethod, // line #792: var topG = topF;
		OtherMethod, // line #793: void test<S>(T pf<T>(T e)) {
		OtherMethod, // line #794:   var c = new C<int>();
		OtherMethod, // line #795:   T lf<T>(T e) => null;
		OtherMethod, // line #796:   var methodTearOffInst = c.f<int>;
		OtherMethod, // line #797:   var staticTearOffInst = C.g<int>;
		OtherMethod, // line #798:   var staticFieldTearOffInst = C.h<int>;
		OtherMethod, // line #799:   var topFunTearOffInst = topF<int>;
		OtherMethod, // line #800:   var topFieldTearOffInst = topG<int>;
		OtherMethod, // line #801:   var localTearOffInst = lf<int>;
		OtherMethod, // line #802:   var paramTearOffInst = pf<int>;
		OtherMethod, // line #803: }
		OtherMethod, // line #804: ''');
		OtherMethod, // line #805:     expectIdentifierType('methodTearOffInst', "int Function(int)");
		OtherMethod, // line #806:     expectIdentifierType('staticTearOffInst', "int Function(int)");
		OtherMethod, // line #807:     expectIdentifierType('staticFieldTearOffInst', "int Function(int)");
		OtherMethod, // line #808:     expectIdentifierType('topFunTearOffInst', "int Function(int)");
		OtherMethod, // line #809:     expectIdentifierType('topFieldTearOffInst', "int Function(int)");
		OtherMethod, // line #810:     expectIdentifierType('localTearOffInst', "int Function(int)");
		OtherMethod, // line #811:     expectIdentifierType('paramTearOffInst', "int Function(int)");
		OtherMethod, // line #812:   }
		BlankLine,   // line #813:
		OtherMethod, // line #814:   test_genericMethod_then() async {
		OtherMethod, // line #815:     await assertErrorsInCode(r'''
		OtherMethod, // line #816: String toString(int x) => x.toString();
		OtherMethod, // line #817: main() {
		OtherMethod, // line #818:   Future<int> bar = null;
		OtherMethod, // line #819:   var foo = bar.then(toString);
		OtherMethod, // line #820: }
		OtherMethod, // line #821: ''', [
		OtherMethod, // line #822:       error(HintCode.UNUSED_LOCAL_VARIABLE, 81, 3),
		OtherMethod, // line #823:     ]);
		OtherMethod, // line #824:
		OtherMethod, // line #825:     expectInitializerType('foo', 'Future<String>');
		OtherMethod, // line #826:   }
		BlankLine,   // line #827:
		OtherMethod, // line #828:   test_genericMethod_then_prefixed() async {
		OtherMethod, // line #829:     await assertErrorsInCode(r'''
		OtherMethod, // line #830: import 'dart:async' as async;
		OtherMethod, // line #831: String toString(int x) => x.toString();
		OtherMethod, // line #832: main() {
		OtherMethod, // line #833:   async.Future<int> bar = null;
		OtherMethod, // line #834:   var foo = bar.then(toString);
		OtherMethod, // line #835: }
		OtherMethod, // line #836: ''', [
		OtherMethod, // line #837:       error(HintCode.UNUSED_LOCAL_VARIABLE, 117, 3),
		OtherMethod, // line #838:     ]);
		OtherMethod, // line #839:     expectInitializerType('foo', 'Future<String>');
		OtherMethod, // line #840:   }
		BlankLine,   // line #841:
		OtherMethod, // line #842:   test_genericMethod_then_propagatedType() async {
		OtherMethod, // line #843:     // Regression test for https://github.com/dart-lang/sdk/issues/25482.
		OtherMethod, // line #844:     await assertErrorsInCode(r'''
		OtherMethod, // line #845: void main() {
		OtherMethod, // line #846:   Future<String> p;
		OtherMethod, // line #847:   var foo = p.then((r) => new Future<String>.value(3));
		OtherMethod, // line #848: }
		OtherMethod, // line #849: ''', [
		OtherMethod, // line #850:       error(HintCode.UNUSED_LOCAL_VARIABLE, 40, 3),
		OtherMethod, // line #851:       error(CompileTimeErrorCode.ARGUMENT_TYPE_NOT_ASSIGNABLE, 85, 1),
		OtherMethod, // line #852:     ]);
		OtherMethod, // line #853:     // Note: this correctly reports the error
		OtherMethod, // line #854:     // CompileTimeErrorCode.ARGUMENT_TYPE_NOT_ASSIGNABLE when run with the driver;
		OtherMethod, // line #855:     // when run without the driver, it reports no errors.  So we don't bother
		OtherMethod, // line #856:     // checking whether the correct errors were reported.
		OtherMethod, // line #857:     expectInitializerType('foo', 'Future<String>');
		OtherMethod, // line #858:   }
		BlankLine,   // line #859:
		OtherMethod, // line #860:   test_genericMethod_toplevel_field_staticTearoff() async {
		OtherMethod, // line #861:     await assertErrorsInCode(r'''
		OtherMethod, // line #862: class C<E> {
		OtherMethod, // line #863:   static T g<T>(T e) => null;
		OtherMethod, // line #864:   static T Function<T>(T) h = null;
		OtherMethod, // line #865: }
		OtherMethod, // line #866:
		OtherMethod, // line #867: void test() {
		OtherMethod, // line #868:   var fieldRead = C.h;
		OtherMethod, // line #869: }
		OtherMethod, // line #870: ''', [
		OtherMethod, // line #871:       error(HintCode.UNUSED_LOCAL_VARIABLE, 102, 9),
		OtherMethod, // line #872:     ]);
		OtherMethod, // line #873:     _assertLocalVarType('fieldRead', "T Function<T>(T)");
		OtherMethod, // line #874:   }
		BlankLine,   // line #875:
		OtherMethod, // line #876:   test_implicitBounds() async {
		OtherMethod, // line #877:     await assertErrorsInCode(r'''
		OtherMethod, // line #878: class A<T> {}
		OtherMethod, // line #879:
		OtherMethod, // line #880: class B<T extends num> {}
		OtherMethod, // line #881:
		OtherMethod, // line #882: class C<S extends int, T extends B<S>, U extends A> {}
		OtherMethod, // line #883:
		OtherMethod, // line #884: void test() {
		OtherMethod, // line #885:   A ai;
		OtherMethod, // line #886:   B bi;
		OtherMethod, // line #887:   C ci;
		OtherMethod, // line #888:   var aa = new A();
		OtherMethod, // line #889:   var bb = new B();
		OtherMethod, // line #890:   var cc = new C();
		OtherMethod, // line #891: }
		OtherMethod, // line #892: ''', [
		OtherMethod, // line #893:       error(HintCode.UNUSED_LOCAL_VARIABLE, 116, 2),
		OtherMethod, // line #894:       error(HintCode.UNUSED_LOCAL_VARIABLE, 124, 2),
		OtherMethod, // line #895:       error(HintCode.UNUSED_LOCAL_VARIABLE, 132, 2),
		OtherMethod, // line #896:       error(HintCode.UNUSED_LOCAL_VARIABLE, 142, 2),
		OtherMethod, // line #897:       error(HintCode.UNUSED_LOCAL_VARIABLE, 162, 2),
		OtherMethod, // line #898:       error(HintCode.UNUSED_LOCAL_VARIABLE, 182, 2),
		OtherMethod, // line #899:     ]);
		OtherMethod, // line #900:     _assertLocalVarType('ai', "A<dynamic>");
		OtherMethod, // line #901:     _assertLocalVarType('bi', "B<num>");
		OtherMethod, // line #902:     _assertLocalVarType('ci', "C<int, B<int>, A<dynamic>>");
		OtherMethod, // line #903:     _assertLocalVarType('aa', "A<dynamic>");
		OtherMethod, // line #904:     _assertLocalVarType('bb', "B<num>");
		OtherMethod, // line #905:     _assertLocalVarType('cc', "C<int, B<int>, A<dynamic>>");
		OtherMethod, // line #906:   }
		BlankLine,   // line #907:
		OtherMethod, // line #908:   test_instantiateToBounds_class_error_extension_malbounded() async {
		OtherMethod, // line #909:     // Test that superclasses are strictly checked for malbounded default
		OtherMethod, // line #910:     // types
		OtherMethod, // line #911:     await assertErrorsInCode(r'''
		OtherMethod, // line #912: class C<T0 extends List<T1>, T1 extends List<T0>> {}
		OtherMethod, // line #913: class D extends C {}
		OtherMethod, // line #914: ''', [
		OtherMethod, // line #915:       error(CompileTimeErrorCode.TYPE_ARGUMENT_NOT_MATCHING_BOUNDS, 69, 1),
		OtherMethod, // line #916:     ]);
		OtherMethod, // line #917:   }
		BlankLine,   // line #918:
		OtherMethod, // line #919:   test_instantiateToBounds_class_error_instantiation_malbounded() async {
		OtherMethod, // line #920:     // Test that instance creations are strictly checked for malbounded default
		OtherMethod, // line #921:     // types
		OtherMethod, // line #922:     await assertErrorsInCode(r'''
		OtherMethod, // line #923: class C<T0 extends List<T1>, T1 extends List<T0>> {}
		OtherMethod, // line #924: void test() {
		OtherMethod, // line #925:   var c = new C();
		OtherMethod, // line #926: }
		OtherMethod, // line #927: ''', [
		OtherMethod, // line #928:       error(HintCode.UNUSED_LOCAL_VARIABLE, 73, 1),
		OtherMethod, // line #929:       error(CompileTimeErrorCode.TYPE_ARGUMENT_NOT_MATCHING_BOUNDS, 81, 1),
		OtherMethod, // line #930:       error(CompileTimeErrorCode.COULD_NOT_INFER, 81, 1),
		OtherMethod, // line #931:     ]);
		OtherMethod, // line #932:     _assertLocalVarType('c', 'C<List<dynamic>, List<List<dynamic>>>');
		OtherMethod, // line #933:   }
		BlankLine,   // line #934:
		OtherMethod, // line #935:   test_instantiateToBounds_class_error_recursion() async {
		OtherMethod, // line #936:     await assertNoErrorsInCode(r'''
		OtherMethod, // line #937: class C<T0 extends List<T1>, T1 extends List<T0>> {}
		OtherMethod, // line #938: C c;
		OtherMethod, // line #939: ''');
		OtherMethod, // line #940:     _assertTopVarType('c', 'C<List<dynamic>, List<dynamic>>');
		OtherMethod, // line #941:   }
		BlankLine,   // line #942:
		OtherMethod, // line #943:   test_instantiateToBounds_class_error_recursion_self() async {
		OtherMethod, // line #944:     await assertNoErrorsInCode(r'''
		OtherMethod, // line #945: class C<T extends C<T>> {}
		OtherMethod, // line #946: C c;
		OtherMethod, // line #947: ''');
		OtherMethod, // line #948:     _assertTopVarType('c', 'C<C<dynamic>>');
		OtherMethod, // line #949:   }
		BlankLine,   // line #950:
		OtherMethod, // line #951:   test_instantiateToBounds_class_error_recursion_self2() async {
		OtherMethod, // line #952:     await assertNoErrorsInCode(r'''
		OtherMethod, // line #953: class A<E> {}
		OtherMethod, // line #954: class C<T extends A<T>> {}
		OtherMethod, // line #955: C c;
		OtherMethod, // line #956: ''');
		OtherMethod, // line #957:     _assertTopVarType('c', 'C<A<dynamic>>');
		OtherMethod, // line #958:   }
		BlankLine,   // line #959:
		OtherMethod, // line #960:   test_instantiateToBounds_class_error_typedef() async {
		OtherMethod, // line #961:     await assertNoErrorsInCode(r'''
		OtherMethod, // line #962: typedef T F<T>(T x);
		OtherMethod, // line #963: class C<T extends F<T>> {}
		OtherMethod, // line #964: C c;
		OtherMethod, // line #965: ''');
		OtherMethod, // line #966:     _assertTopVarType('c', 'C<dynamic Function(dynamic)>');
		OtherMethod, // line #967:   }
		BlankLine,   // line #968:
		OtherMethod, // line #969:   test_instantiateToBounds_class_ok_implicitDynamic_multi() async {
		OtherMethod, // line #970:     await assertNoErrorsInCode(r'''
		OtherMethod, // line #971: class C<T0 extends Map<T1, T2>, T1 extends List, T2 extends int> {}
		OtherMethod, // line #972: C c;
		OtherMethod, // line #973: ''');
		OtherMethod, // line #974:     _assertTopVarType('c', 'C<Map<List<dynamic>, int>, List<dynamic>, int>');
		OtherMethod, // line #975:   }
		BlankLine,   // line #976:
		OtherMethod, // line #977:   test_instantiateToBounds_class_ok_referenceOther_after() async {
		OtherMethod, // line #978:     await assertNoErrorsInCode(r'''
		OtherMethod, // line #979: class C<T0 extends T1, T1 extends int> {}
		OtherMethod, // line #980: C c;
		OtherMethod, // line #981: ''');
		OtherMethod, // line #982:     _assertTopVarType('c', 'C<int, int>');
		OtherMethod, // line #983:   }
		BlankLine,   // line #984:
		OtherMethod, // line #985:   test_instantiateToBounds_class_ok_referenceOther_after2() async {
		OtherMethod, // line #986:     await assertNoErrorsInCode(r'''
		OtherMethod, // line #987: class C<T0 extends Map<T1, T1>, T1 extends int> {}
		OtherMethod, // line #988: C c;
		OtherMethod, // line #989: ''');
		OtherMethod, // line #990:     _assertTopVarType('c', 'C<Map<int, int>, int>');
		OtherMethod, // line #991:   }
		BlankLine,   // line #992:
		OtherMethod, // line #993:   test_instantiateToBounds_class_ok_referenceOther_before() async {
		OtherMethod, // line #994:     await assertNoErrorsInCode(r'''
		OtherMethod, // line #995: class C<T0 extends int, T1 extends T0> {}
		OtherMethod, // line #996: C c;
		OtherMethod, // line #997: ''');
		OtherMethod, // line #998:     _assertTopVarType('c', 'C<int, int>');
		OtherMethod, // line #999:   }
		BlankLine,   // line #1000:
		OtherMethod, // line #1001:   test_instantiateToBounds_class_ok_referenceOther_multi() async {
		OtherMethod, // line #1002:     await assertNoErrorsInCode(r'''
		OtherMethod, // line #1003: class C<T0 extends Map<T1, T2>, T1 extends List<T2>, T2 extends int> {}
		OtherMethod, // line #1004: C c;
		OtherMethod, // line #1005: ''');
		OtherMethod, // line #1006:     _assertTopVarType('c', 'C<Map<List<int>, int>, List<int>, int>');
		OtherMethod, // line #1007:   }
		BlankLine,   // line #1008:
		OtherMethod, // line #1009:   test_instantiateToBounds_class_ok_simpleBounds() async {
		OtherMethod, // line #1010:     await assertErrorsInCode(r'''
		OtherMethod, // line #1011: class A<T> {}
		OtherMethod, // line #1012: class B<T extends num> {}
		OtherMethod, // line #1013: class C<T extends List<int>> {}
		OtherMethod, // line #1014: class D<T extends A> {}
		OtherMethod, // line #1015: void main() {
		OtherMethod, // line #1016:   A a;
		OtherMethod, // line #1017:   B b;
		OtherMethod, // line #1018:   C c;
		OtherMethod, // line #1019:   D d;
		OtherMethod, // line #1020: }
		OtherMethod, // line #1021: ''', [
		OtherMethod, // line #1022:       error(HintCode.UNUSED_LOCAL_VARIABLE, 114, 1),
		OtherMethod, // line #1023:       error(HintCode.UNUSED_LOCAL_VARIABLE, 121, 1),
		OtherMethod, // line #1024:       error(HintCode.UNUSED_LOCAL_VARIABLE, 128, 1),
		OtherMethod, // line #1025:       error(HintCode.UNUSED_LOCAL_VARIABLE, 135, 1),
		OtherMethod, // line #1026:     ]);
		OtherMethod, // line #1027:     _assertLocalVarType('a', 'A<dynamic>');
		OtherMethod, // line #1028:     _assertLocalVarType('b', 'B<num>');
		OtherMethod, // line #1029:     _assertLocalVarType('c', 'C<List<int>>');
		OtherMethod, // line #1030:     _assertLocalVarType('d', 'D<A<dynamic>>');
		OtherMethod, // line #1031:   }
		BlankLine,   // line #1032:
		OtherMethod, // line #1033:   test_instantiateToBounds_generic_function_error_malbounded() async {
		OtherMethod, // line #1034:     // Test that generic methods are strictly checked for malbounded default
		OtherMethod, // line #1035:     // types
		OtherMethod, // line #1036:     await assertErrorsInCode(r'''
		OtherMethod, // line #1037: T0 f<T0 extends List<T1>, T1 extends List<T0>>() {}
		OtherMethod, // line #1038: void g() {
		OtherMethod, // line #1039:   var c = f();
		OtherMethod, // line #1040:   return;
		OtherMethod, // line #1041: }
		OtherMethod, // line #1042: ''', [
		OtherMethod, // line #1043:       error(HintCode.MISSING_RETURN, 3, 1),
		OtherMethod, // line #1044:       error(HintCode.UNUSED_LOCAL_VARIABLE, 69, 1),
		OtherMethod, // line #1045:       error(CompileTimeErrorCode.COULD_NOT_INFER, 73, 1),
		OtherMethod, // line #1046:     ]);
		OtherMethod, // line #1047:     _assertLocalVarType('c', 'List<dynamic>');
		OtherMethod, // line #1048:   }
		BlankLine,   // line #1049:
		OtherMethod, // line #1050:   test_instantiateToBounds_method_ok_referenceOther_before() async {
		OtherMethod, // line #1051:     await assertNoErrorsInCode(r'''
		OtherMethod, // line #1052: class C<T> {
		OtherMethod, // line #1053:   void m<S0 extends T, S1 extends List<S0>>(S0 p0, S1 p1) {}
		OtherMethod, // line #1054:
		OtherMethod, // line #1055:   void main() {
		OtherMethod, // line #1056:     m(null, null);
		OtherMethod, // line #1057:   }
		OtherMethod, // line #1058: }
		OtherMethod, // line #1059: ''');
		OtherMethod, // line #1060:
		OtherMethod, // line #1061:     expectStaticInvokeType('m(null', 'void Function(Null, Null)');
		OtherMethod, // line #1062:   }
		BlankLine,   // line #1063:
		OtherMethod, // line #1064:   test_instantiateToBounds_method_ok_referenceOther_before2() async {
		OtherMethod, // line #1065:     await assertNoErrorsInCode(r'''
		OtherMethod, // line #1066: class C<T> {
		OtherMethod, // line #1067:   Map<S0, S1> m<S0 extends T, S1 extends List<S0>>() => null;
		OtherMethod, // line #1068:
		OtherMethod, // line #1069:   void main() {
		OtherMethod, // line #1070:     m();
		OtherMethod, // line #1071:   }
		OtherMethod, // line #1072: }
		OtherMethod, // line #1073: ''');
		OtherMethod, // line #1074:
		OtherMethod, // line #1075:     expectStaticInvokeType('m();', 'Map<T, List<T>> Function()');
		OtherMethod, // line #1076:   }
		BlankLine,   // line #1077:
		OtherMethod, // line #1078:   test_instantiateToBounds_method_ok_simpleBounds() async {
		OtherMethod, // line #1079:     await assertNoErrorsInCode(r'''
		OtherMethod, // line #1080: class C<T> {
		OtherMethod, // line #1081:   void m<S extends T>(S p0) {}
		OtherMethod, // line #1082:
		OtherMethod, // line #1083:   void main() {
		OtherMethod, // line #1084:     m(null);
		OtherMethod, // line #1085:   }
		OtherMethod, // line #1086: }
		OtherMethod, // line #1087: ''');
		OtherMethod, // line #1088:
		OtherMethod, // line #1089:     expectStaticInvokeType('m(null)', 'void Function(Null)');
		OtherMethod, // line #1090:   }
		BlankLine,   // line #1091:
		OtherMethod, // line #1092:   test_instantiateToBounds_method_ok_simpleBounds2() async {
		OtherMethod, // line #1093:     await assertNoErrorsInCode(r'''
		OtherMethod, // line #1094: class C<T> {
		OtherMethod, // line #1095:   S m<S extends T>() => null;
		OtherMethod, // line #1096:
		OtherMethod, // line #1097:   void main() {
		OtherMethod, // line #1098:     m();
		OtherMethod, // line #1099:   }
		OtherMethod, // line #1100: }
		OtherMethod, // line #1101: ''');
		OtherMethod, // line #1102:
		OtherMethod, // line #1103:     expectStaticInvokeType('m();', 'T Function()');
		OtherMethod, // line #1104:   }
		BlankLine,   // line #1105:
		OtherMethod, // line #1106:   test_issue32396() async {
		OtherMethod, // line #1107:     await assertNoErrorsInCode(r'''
		OtherMethod, // line #1108: class C<E> {
		OtherMethod, // line #1109:   static T g<T>(T e) => null;
		OtherMethod, // line #1110:   static final h = g;
		OtherMethod, // line #1111: }
		OtherMethod, // line #1112: ''');
		OtherMethod, // line #1113:   }
		BlankLine,   // line #1114:
		OtherMethod, // line #1115:   test_objectMethodOnFunctions_Anonymous() async {
		OtherMethod, // line #1116:     await _objectMethodOnFunctions_helper2(r'''
		OtherMethod, // line #1117: void main() {
		OtherMethod, // line #1118:   var f = (x) => 3;
		OtherMethod, // line #1119:   // No errors, correct type
		OtherMethod, // line #1120:   var t0 = f.toString();
		OtherMethod, // line #1121:   var t1 = f.toString;
		OtherMethod, // line #1122:   var t2 = f.hashCode;
		OtherMethod, // line #1123:
		OtherMethod, // line #1124:   // Expressions, no errors, correct type
		OtherMethod, // line #1125:   var t3 = (f).toString();
		OtherMethod, // line #1126:   var t4 = (f).toString;
		OtherMethod, // line #1127:   var t5 = (f).hashCode;
		OtherMethod, // line #1128:
		OtherMethod, // line #1129:   // Cascades, no errors
		OtherMethod, // line #1130:   f..toString();
		OtherMethod, // line #1131:   f..toString;
		OtherMethod, // line #1132:   f..hashCode;
		OtherMethod, // line #1133:
		OtherMethod, // line #1134:   // Expression cascades, no errors
		OtherMethod, // line #1135:   (f)..toString();
		OtherMethod, // line #1136:   (f)..toString;
		OtherMethod, // line #1137:   (f)..hashCode;
		OtherMethod, // line #1138: }
		OtherMethod, // line #1139: ''', [
		OtherMethod, // line #1140:       error(HintCode.UNUSED_LOCAL_VARIABLE, 69, 2),
		OtherMethod, // line #1141:       error(HintCode.UNUSED_LOCAL_VARIABLE, 94, 2),
		OtherMethod, // line #1142:       error(HintCode.UNUSED_LOCAL_VARIABLE, 117, 2),
		OtherMethod, // line #1143:       error(HintCode.UNUSED_LOCAL_VARIABLE, 183, 2),
		OtherMethod, // line #1144:       error(HintCode.UNUSED_LOCAL_VARIABLE, 210, 2),
		OtherMethod, // line #1145:       error(HintCode.UNUSED_LOCAL_VARIABLE, 235, 2),
		OtherMethod, // line #1146:     ]);
		OtherMethod, // line #1147:   }
		BlankLine,   // line #1148:
		OtherMethod, // line #1149:   test_objectMethodOnFunctions_Function() async {
		OtherMethod, // line #1150:     await _objectMethodOnFunctions_helper2(r'''
		OtherMethod, // line #1151: void main() {
		OtherMethod, // line #1152:   Function f;
		OtherMethod, // line #1153:   // No errors, correct type
		OtherMethod, // line #1154:   var t0 = f.toString();
		OtherMethod, // line #1155:   var t1 = f.toString;
		OtherMethod, // line #1156:   var t2 = f.hashCode;
		OtherMethod, // line #1157:
		OtherMethod, // line #1158:   // Expressions, no errors, correct type
		OtherMethod, // line #1159:   var t3 = (f).toString();
		OtherMethod, // line #1160:   var t4 = (f).toString;
		OtherMethod, // line #1161:   var t5 = (f).hashCode;
		OtherMethod, // line #1162:
		OtherMethod, // line #1163:   // Cascades, no errors
		OtherMethod, // line #1164:   f..toString();
		OtherMethod, // line #1165:   f..toString;
		OtherMethod, // line #1166:   f..hashCode;
		OtherMethod, // line #1167:
		OtherMethod, // line #1168:   // Expression cascades, no errors
		OtherMethod, // line #1169:   (f)..toString();
		OtherMethod, // line #1170:   (f)..toString;
		OtherMethod, // line #1171:   (f)..hashCode;
		OtherMethod, // line #1172: }
		OtherMethod, // line #1173: ''', [
		OtherMethod, // line #1174:       error(HintCode.UNUSED_LOCAL_VARIABLE, 63, 2),
		OtherMethod, // line #1175:       error(HintCode.UNUSED_LOCAL_VARIABLE, 88, 2),
		OtherMethod, // line #1176:       error(HintCode.UNUSED_LOCAL_VARIABLE, 111, 2),
		OtherMethod, // line #1177:       error(HintCode.UNUSED_LOCAL_VARIABLE, 177, 2),
		OtherMethod, // line #1178:       error(HintCode.UNUSED_LOCAL_VARIABLE, 204, 2),
		OtherMethod, // line #1179:       error(HintCode.UNUSED_LOCAL_VARIABLE, 229, 2),
		OtherMethod, // line #1180:     ]);
		OtherMethod, // line #1181:   }
		BlankLine,   // line #1182:
		OtherMethod, // line #1183:   test_objectMethodOnFunctions_Static() async {
		OtherMethod, // line #1184:     await _objectMethodOnFunctions_helper2(r'''
		OtherMethod, // line #1185: int f(int x) => null;
		OtherMethod, // line #1186: void main() {
		OtherMethod, // line #1187:   // No errors, correct type
		OtherMethod, // line #1188:   var t0 = f.toString();
		OtherMethod, // line #1189:   var t1 = f.toString;
		OtherMethod, // line #1190:   var t2 = f.hashCode;
		OtherMethod, // line #1191:
		OtherMethod, // line #1192:   // Expressions, no errors, correct type
		OtherMethod, // line #1193:   var t3 = (f).toString();
		OtherMethod, // line #1194:   var t4 = (f).toString;
		OtherMethod, // line #1195:   var t5 = (f).hashCode;
		OtherMethod, // line #1196:
		OtherMethod, // line #1197:   // Cascades, no errors
		OtherMethod, // line #1198:   f..toString();
		OtherMethod, // line #1199:   f..toString;
		OtherMethod, // line #1200:   f..hashCode;
		OtherMethod, // line #1201:
		OtherMethod, // line #1202:   // Expression cascades, no errors
		OtherMethod, // line #1203:   (f)..toString();
		OtherMethod, // line #1204:   (f)..toString;
		OtherMethod, // line #1205:   (f)..hashCode;
		OtherMethod, // line #1206: }
		OtherMethod, // line #1207: ''', [
		OtherMethod, // line #1208:       error(HintCode.UNUSED_LOCAL_VARIABLE, 71, 2),
		OtherMethod, // line #1209:       error(HintCode.UNUSED_LOCAL_VARIABLE, 96, 2),
		OtherMethod, // line #1210:       error(HintCode.UNUSED_LOCAL_VARIABLE, 119, 2),
		OtherMethod, // line #1211:       error(HintCode.UNUSED_LOCAL_VARIABLE, 185, 2),
		OtherMethod, // line #1212:       error(HintCode.UNUSED_LOCAL_VARIABLE, 212, 2),
		OtherMethod, // line #1213:       error(HintCode.UNUSED_LOCAL_VARIABLE, 237, 2),
		OtherMethod, // line #1214:     ]);
		OtherMethod, // line #1215:   }
		BlankLine,   // line #1216:
		OtherMethod, // line #1217:   test_objectMethodOnFunctions_Typedef() async {
		OtherMethod, // line #1218:     await _objectMethodOnFunctions_helper2(r'''
		OtherMethod, // line #1219: typedef bool Predicate<T>(T object);
		OtherMethod, // line #1220:
		OtherMethod, // line #1221: void main() {
		OtherMethod, // line #1222:   Predicate<int> f;
		OtherMethod, // line #1223:   // No errors, correct type
		OtherMethod, // line #1224:   var t0 = f.toString();
		OtherMethod, // line #1225:   var t1 = f.toString;
		OtherMethod, // line #1226:   var t2 = f.hashCode;
		OtherMethod, // line #1227:
		OtherMethod, // line #1228:   // Expressions, no errors, correct type
		OtherMethod, // line #1229:   var t3 = (f).toString();
		OtherMethod, // line #1230:   var t4 = (f).toString;
		OtherMethod, // line #1231:   var t5 = (f).hashCode;
		OtherMethod, // line #1232:
		OtherMethod, // line #1233:   // Cascades, no errors
		OtherMethod, // line #1234:   f..toString();
		OtherMethod, // line #1235:   f..toString;
		OtherMethod, // line #1236:   f..hashCode;
		OtherMethod, // line #1237:
		OtherMethod, // line #1238:   // Expression cascades, no errors
		OtherMethod, // line #1239:   (f)..toString();
		OtherMethod, // line #1240:   (f)..toString;
		OtherMethod, // line #1241:   (f)..hashCode;
		OtherMethod, // line #1242: }
		OtherMethod, // line #1243: ''', [
		OtherMethod, // line #1244:       error(HintCode.UNUSED_LOCAL_VARIABLE, 107, 2),
		OtherMethod, // line #1245:       error(HintCode.UNUSED_LOCAL_VARIABLE, 132, 2),
		OtherMethod, // line #1246:       error(HintCode.UNUSED_LOCAL_VARIABLE, 155, 2),
		OtherMethod, // line #1247:       error(HintCode.UNUSED_LOCAL_VARIABLE, 221, 2),
		OtherMethod, // line #1248:       error(HintCode.UNUSED_LOCAL_VARIABLE, 248, 2),
		OtherMethod, // line #1249:       error(HintCode.UNUSED_LOCAL_VARIABLE, 273, 2),
		OtherMethod, // line #1250:     ]);
		OtherMethod, // line #1251:   }
		BlankLine,   // line #1252:
		OtherMethod, // line #1253:   test_returnOfInvalidType_object_void() async {
		OtherMethod, // line #1254:     await assertErrorsInCode(
		OtherMethod, // line #1255:         "Object f() { void voidFn() => null; return voidFn(); }", [
		OtherMethod, // line #1256:       error(CompileTimeErrorCode.RETURN_OF_INVALID_TYPE_FROM_FUNCTION, 43, 8),
		OtherMethod, // line #1257:     ]);
		OtherMethod, // line #1258:   }
		BlankLine,   // line #1259:
		OtherMethod, // line #1260:   test_setterWithDynamicTypeIsError() async {
		OtherMethod, // line #1261:     await assertErrorsInCode(r'''
		OtherMethod, // line #1262: class A {
		OtherMethod, // line #1263:   dynamic set f(String s) => null;
		OtherMethod, // line #1264: }
		OtherMethod, // line #1265: dynamic set g(int x) => null;
		OtherMethod, // line #1266: ''', [
		OtherMethod, // line #1267:       error(CompileTimeErrorCode.NON_VOID_RETURN_FOR_SETTER, 12, 7),
		OtherMethod, // line #1268:       error(CompileTimeErrorCode.NON_VOID_RETURN_FOR_SETTER, 47, 7),
		OtherMethod, // line #1269:     ]);
		OtherMethod, // line #1270:   }
		BlankLine,   // line #1271:
		OtherMethod, // line #1272:   test_setterWithExplicitVoidType_returningVoid() async {
		OtherMethod, // line #1273:     await assertNoErrorsInCode(r'''
		OtherMethod, // line #1274: void returnsVoid() {}
		OtherMethod, // line #1275: class A {
		OtherMethod, // line #1276:   void set f(String s) => returnsVoid();
		OtherMethod, // line #1277: }
		OtherMethod, // line #1278: void set g(int x) => returnsVoid();
		OtherMethod, // line #1279: ''');
		OtherMethod, // line #1280:   }
		BlankLine,   // line #1281:
		OtherMethod, // line #1282:   test_setterWithNoVoidType() async {
		OtherMethod, // line #1283:     await assertErrorsInCode(r'''
		OtherMethod, // line #1284: class A {
		OtherMethod, // line #1285:   set f(String s) {
		OtherMethod, // line #1286:     return '42';
		OtherMethod, // line #1287:   }
		OtherMethod, // line #1288: }
		OtherMethod, // line #1289: set g(int x) => 42;
		OtherMethod, // line #1290: ''', [
		OtherMethod, // line #1291:       error(CompileTimeErrorCode.RETURN_OF_INVALID_TYPE_FROM_FUNCTION, 41, 4),
		OtherMethod, // line #1292:     ]);
		OtherMethod, // line #1293:   }
		BlankLine,   // line #1294:
		OtherMethod, // line #1295:   test_setterWithNoVoidType_returningVoid() async {
		OtherMethod, // line #1296:     await assertNoErrorsInCode(r'''
		OtherMethod, // line #1297: void returnsVoid() {}
		OtherMethod, // line #1298: class A {
		OtherMethod, // line #1299:   set f(String s) => returnsVoid();
		OtherMethod, // line #1300: }
		OtherMethod, // line #1301: set g(int x) => returnsVoid();
		OtherMethod, // line #1302: ''');
		OtherMethod, // line #1303:   }
		BlankLine,   // line #1304:
		OtherMethod, // line #1305:   test_setterWithOtherTypeIsError() async {
		OtherMethod, // line #1306:     await assertErrorsInCode(r'''
		OtherMethod, // line #1307: class A {
		OtherMethod, // line #1308:   String set f(String s) => null;
		OtherMethod, // line #1309: }
		OtherMethod, // line #1310: Object set g(x) => null;
		OtherMethod, // line #1311: ''', [
		OtherMethod, // line #1312:       error(CompileTimeErrorCode.NON_VOID_RETURN_FOR_SETTER, 12, 6),
		OtherMethod, // line #1313:       error(CompileTimeErrorCode.NON_VOID_RETURN_FOR_SETTER, 46, 6),
		OtherMethod, // line #1314:     ]);
		OtherMethod, // line #1315:   }
		BlankLine,   // line #1316:
		OtherMethod, // line #1317:   test_ternaryOperator_null_left() async {
		OtherMethod, // line #1318:     await assertErrorsInCode(r'''
		OtherMethod, // line #1319: main() {
		OtherMethod, // line #1320:   var foo = (true) ? null : 3;
		OtherMethod, // line #1321: }
		OtherMethod, // line #1322: ''', [
		OtherMethod, // line #1323:       error(HintCode.UNUSED_LOCAL_VARIABLE, 15, 3),
		OtherMethod, // line #1324:     ]);
		OtherMethod, // line #1325:     expectInitializerType('foo', 'int');
		OtherMethod, // line #1326:   }
		BlankLine,   // line #1327:
		OtherMethod, // line #1328:   test_ternaryOperator_null_right() async {
		OtherMethod, // line #1329:     await assertErrorsInCode(r'''
		OtherMethod, // line #1330: main() {
		OtherMethod, // line #1331:   var foo = (true) ? 3 : null;
		OtherMethod, // line #1332: }
		OtherMethod, // line #1333: ''', [
		OtherMethod, // line #1334:       error(HintCode.UNUSED_LOCAL_VARIABLE, 15, 3),
		OtherMethod, // line #1335:     ]);
		OtherMethod, // line #1336:     expectInitializerType('foo', 'int');
		OtherMethod, // line #1337:   }
		BlankLine,   // line #1338:
		OtherMethod, // line #1339:   void _assertLocalVarType(String name, String expectedType) {
		OtherMethod, // line #1340:     var element = findElement.localVar(name);
		OtherMethod, // line #1341:     assertType(element.type, expectedType);
		OtherMethod, // line #1342:   }
		BlankLine,   // line #1343:
		OtherMethod, // line #1344:   void _assertTopVarType(String name, String expectedType) {
		OtherMethod, // line #1345:     var element = findElement.topVar(name);
		OtherMethod, // line #1346:     assertType(element.type, expectedType);
		OtherMethod, // line #1347:   }
		BlankLine,   // line #1348:
		OtherMethod, // line #1349:   Future<void> _objectMethodOnFunctions_helper2(
		OtherMethod, // line #1350:       String code, List<ExpectedError> expectedErrors) async {
		OtherMethod, // line #1351:     await assertErrorsInCode(code, expectedErrors);
		OtherMethod, // line #1352:     _assertLocalVarType('t0', "String");
		OtherMethod, // line #1353:     _assertLocalVarType('t1', "String Function()");
		OtherMethod, // line #1354:     _assertLocalVarType('t2', "int");
		OtherMethod, // line #1355:     _assertLocalVarType('t3', "String");
		OtherMethod, // line #1356:     _assertLocalVarType('t4', "String Function()");
		OtherMethod, // line #1357:     _assertLocalVarType('t5', "int");
		OtherMethod, // line #1358:   }
		BlankLine,   // line #1359:
	}

	runParsePhase(t, opts, source, want)
	// runFullStylizer(t, opts, source, wantSource, want)
}

func TestStrongMode_Class3(t *testing.T) {
	source := strong_mode_dart_txt[129066+12 : 134946+12]
	// wantSource := strong_mode_want_txt[1027:2248]

	if source[0] != '@' {
		t.Fatalf("source[0] = %c, want @", source[0])
	}
	if source[len(source)-1] != '\n' {
		t.Fatalf("source[%v] = %c, want newline", len(source)-1, source[len(source)-1])
	}

	opts := &Options{
		MemberOrdering: defaultMemberOrdering,
	}

	want := []EntityType{
		Unknown,     // line #2: {
		OtherMethod, // line #3:   test_foreachInference_dynamic_disabled() async {
		OtherMethod, // line #4:     await resolveTestCode(r'''
		OtherMethod, // line #5: main() {
		OtherMethod, // line #6:   var list = <int>[];
		OtherMethod, // line #7:   for (dynamic v in list) {
		OtherMethod, // line #8:     v; // marker
		OtherMethod, // line #9:   }
		OtherMethod, // line #10: }''');
		OtherMethod, // line #11:     assertTypeDynamic(findElement.localVar('v').type);
		OtherMethod, // line #12:     assertTypeDynamic(findNode.simple('v; // marker'));
		OtherMethod, // line #13:   }
		BlankLine,   // line #14:
		OtherMethod, // line #15:   test_foreachInference_reusedVar_disabled() async {
		OtherMethod, // line #16:     await resolveTestCode(r'''
		OtherMethod, // line #17: main() {
		OtherMethod, // line #18:   var list = <int>[];
		OtherMethod, // line #19:   var v;
		OtherMethod, // line #20:   for (v in list) {
		OtherMethod, // line #21:     v; // marker
		OtherMethod, // line #22:   }
		OtherMethod, // line #23: }''');
		OtherMethod, // line #24:     assertTypeDynamic(findNode.simple('v in'));
		OtherMethod, // line #25:     assertTypeDynamic(findNode.simple('v; // marker'));
		OtherMethod, // line #26:   }
		BlankLine,   // line #27:
		OtherMethod, // line #28:   test_foreachInference_var() async {
		OtherMethod, // line #29:     await resolveTestCode(r'''
		OtherMethod, // line #30: main() {
		OtherMethod, // line #31:   var list = <int>[];
		OtherMethod, // line #32:   for (var v in list) {
		OtherMethod, // line #33:     v; // marker
		OtherMethod, // line #34:   }
		OtherMethod, // line #35: }''');
		OtherMethod, // line #36:     assertType(findElement.localVar('v').type, 'int');
		OtherMethod, // line #37:     assertType(findNode.simple('v; // marker'), 'int');
		OtherMethod, // line #38:   }
		BlankLine,   // line #39:
		OtherMethod, // line #40:   test_foreachInference_var_iterable() async {
		OtherMethod, // line #41:     await resolveTestCode(r'''
		OtherMethod, // line #42: main() {
		OtherMethod, // line #43:   Iterable<int> list = <int>[];
		OtherMethod, // line #44:   for (var v in list) {
		OtherMethod, // line #45:     v; // marker
		OtherMethod, // line #46:   }
		OtherMethod, // line #47: }''');
		OtherMethod, // line #48:     assertType(findElement.localVar('v').type, 'int');
		OtherMethod, // line #49:     assertType(findNode.simple('v; // marker'), 'int');
		OtherMethod, // line #50:   }
		BlankLine,   // line #51:
		OtherMethod, // line #52:   test_foreachInference_var_stream() async {
		OtherMethod, // line #53:     await resolveTestCode(r'''
		OtherMethod, // line #54: main() async {
		OtherMethod, // line #55:   Stream<int> stream = null;
		OtherMethod, // line #56:   await for (var v in stream) {
		OtherMethod, // line #57:     v; // marker
		OtherMethod, // line #58:   }
		OtherMethod, // line #59: }''');
		OtherMethod, // line #60:     assertType(findElement.localVar('v').type, 'int');
		OtherMethod, // line #61:     assertType(findNode.simple('v; // marker'), 'int');
		OtherMethod, // line #62:   }
		BlankLine,   // line #63:
		OtherMethod, // line #64:   test_inconsistentMethodInheritance_inferFunctionTypeFromTypedef() async {
		OtherMethod, // line #65:     await assertNoErrorsInCode(r'''
		OtherMethod, // line #66: typedef bool F<E>(E argument);
		OtherMethod, // line #67:
		OtherMethod, // line #68: abstract class Base {
		OtherMethod, // line #69:   f<E extends int>(F<int> x);
		OtherMethod, // line #70: }
		OtherMethod, // line #71:
		OtherMethod, // line #72: abstract class BaseCopy extends Base {
		OtherMethod, // line #73: }
		OtherMethod, // line #74:
		OtherMethod, // line #75: abstract class Override implements Base, BaseCopy {
		OtherMethod, // line #76:   f<E extends int>(x) => null;
		OtherMethod, // line #77: }
		OtherMethod, // line #78:
		OtherMethod, // line #79: class C extends Override implements Base {}
		OtherMethod, // line #80: ''');
		OtherMethod, // line #81:   }
		BlankLine,   // line #82:
		OtherMethod, // line #83:   test_localVariableInference_bottom_disabled() async {
		OtherMethod, // line #84:     await resolveTestCode(r'''
		OtherMethod, // line #85: main() {
		OtherMethod, // line #86:   var v = null;
		OtherMethod, // line #87:   v; // marker
		OtherMethod, // line #88: }''');
		OtherMethod, // line #89:     assertTypeDynamic(findElement.localVar('v').type);
		OtherMethod, // line #90:     assertTypeDynamic(findNode.simple('v; // marker'));
		OtherMethod, // line #91:   }
		BlankLine,   // line #92:
		OtherMethod, // line #93:   test_localVariableInference_constant() async {
		OtherMethod, // line #94:     await resolveTestCode(r'''
		OtherMethod, // line #95: main() {
		OtherMethod, // line #96:   var v = 3;
		OtherMethod, // line #97:   v; // marker
		OtherMethod, // line #98: }''');
		OtherMethod, // line #99:     assertType(findElement.localVar('v').type, 'int');
		OtherMethod, // line #100:     assertType(findNode.simple('v; // marker'), 'int');
		OtherMethod, // line #101:   }
		BlankLine,   // line #102:
		OtherMethod, // line #103:   test_localVariableInference_declaredType_disabled() async {
		OtherMethod, // line #104:     await resolveTestCode(r'''
		OtherMethod, // line #105: main() {
		OtherMethod, // line #106:   dynamic v = 3;
		OtherMethod, // line #107:   v; // marker
		OtherMethod, // line #108: }''');
		OtherMethod, // line #109:     assertTypeDynamic(findElement.localVar('v').type);
		OtherMethod, // line #110:     assertTypeDynamic(findNode.simple('v; // marker'));
		OtherMethod, // line #111:   }
		BlankLine,   // line #112:
		OtherMethod, // line #113:   test_localVariableInference_noInitializer_disabled() async {
		OtherMethod, // line #114:     await resolveTestCode(r'''
		OtherMethod, // line #115: main() {
		OtherMethod, // line #116:   var v;
		OtherMethod, // line #117:   v = 3;
		OtherMethod, // line #118:   v; // marker
		OtherMethod, // line #119: }''');
		OtherMethod, // line #120:     if (hasAssignmentLeftResolution) {
		OtherMethod, // line #121:       assertTypeDynamic(findNode.simple('v ='));
		OtherMethod, // line #122:     } else {
		OtherMethod, // line #123:       assertTypeNull(findNode.simple('v ='));
		OtherMethod, // line #124:     }
		OtherMethod, // line #125:     assertTypeDynamic(findNode.simple('v; // marker'));
		OtherMethod, // line #126:   }
		BlankLine,   // line #127:
		OtherMethod, // line #128:   test_localVariableInference_transitive_field_inferred_lexical() async {
		OtherMethod, // line #129:     await resolveTestCode(r'''
		OtherMethod, // line #130: class A {
		OtherMethod, // line #131:   final x = 3;
		OtherMethod, // line #132:   f() {
		OtherMethod, // line #133:     var v = x;
		OtherMethod, // line #134:     return v; // marker
		OtherMethod, // line #135:   }
		OtherMethod, // line #136: }
		OtherMethod, // line #137: main() {
		OtherMethod, // line #138: }
		OtherMethod, // line #139: ''');
		OtherMethod, // line #140:     assertType(findElement.localVar('v').type, 'int');
		OtherMethod, // line #141:     assertType(findNode.simple('v; // marker'), 'int');
		OtherMethod, // line #142:   }
		BlankLine,   // line #143:
		OtherMethod, // line #144:   test_localVariableInference_transitive_field_inferred_reversed() async {
		OtherMethod, // line #145:     await resolveTestCode(r'''
		OtherMethod, // line #146: class A {
		OtherMethod, // line #147:   f() {
		OtherMethod, // line #148:     var v = x;
		OtherMethod, // line #149:     return v; // marker
		OtherMethod, // line #150:   }
		OtherMethod, // line #151:   final x = 3;
		OtherMethod, // line #152: }
		OtherMethod, // line #153: main() {
		OtherMethod, // line #154: }
		OtherMethod, // line #155: ''');
		OtherMethod, // line #156:     assertType(findElement.localVar('v').type, 'int');
		OtherMethod, // line #157:     assertType(findNode.simple('v; // marker'), 'int');
		OtherMethod, // line #158:   }
		BlankLine,   // line #159:
		OtherMethod, // line #160:   test_localVariableInference_transitive_field_lexical() async {
		OtherMethod, // line #161:     await resolveTestCode(r'''
		OtherMethod, // line #162: class A {
		OtherMethod, // line #163:   int x = 3;
		OtherMethod, // line #164:   f() {
		OtherMethod, // line #165:     var v = x;
		OtherMethod, // line #166:     return v; // marker
		OtherMethod, // line #167:   }
		OtherMethod, // line #168: }
		OtherMethod, // line #169: main() {
		OtherMethod, // line #170: }
		OtherMethod, // line #171: ''');
		OtherMethod, // line #172:     assertType(findElement.localVar('v').type, 'int');
		OtherMethod, // line #173:     assertType(findNode.simple('v; // marker'), 'int');
		OtherMethod, // line #174:   }
		BlankLine,   // line #175:
		OtherMethod, // line #176:   test_localVariableInference_transitive_field_reversed() async {
		OtherMethod, // line #177:     await resolveTestCode(r'''
		OtherMethod, // line #178: class A {
		OtherMethod, // line #179:   f() {
		OtherMethod, // line #180:     var v = x;
		OtherMethod, // line #181:     return v; // marker
		OtherMethod, // line #182:   }
		OtherMethod, // line #183:   int x = 3;
		OtherMethod, // line #184: }
		OtherMethod, // line #185: main() {
		OtherMethod, // line #186: }
		OtherMethod, // line #187: ''');
		OtherMethod, // line #188:     assertType(findElement.localVar('v').type, 'int');
		OtherMethod, // line #189:     assertType(findNode.simple('v; // marker'), 'int');
		OtherMethod, // line #190:   }
		BlankLine,   // line #191:
		OtherMethod, // line #192:   test_localVariableInference_transitive_list_local() async {
		OtherMethod, // line #193:     await resolveTestCode(r'''
		OtherMethod, // line #194: main() {
		OtherMethod, // line #195:   var x = <int>[3];
		OtherMethod, // line #196:   var v = x[0];
		OtherMethod, // line #197:   v; // marker
		OtherMethod, // line #198: }''');
		OtherMethod, // line #199:     assertType(findElement.localVar('v').type, 'int');
		OtherMethod, // line #200:     assertType(findNode.simple('v; // marker'), 'int');
		OtherMethod, // line #201:   }
		BlankLine,   // line #202:
		OtherMethod, // line #203:   test_localVariableInference_transitive_local() async {
		OtherMethod, // line #204:     await resolveTestCode(r'''
		OtherMethod, // line #205: main() {
		OtherMethod, // line #206:   var x = 3;
		OtherMethod, // line #207:   var v = x;
		OtherMethod, // line #208:   v; // marker
		OtherMethod, // line #209: }''');
		OtherMethod, // line #210:     assertType(findElement.localVar('v').type, 'int');
		OtherMethod, // line #211:     assertType(findNode.simple('v; // marker'), 'int');
		OtherMethod, // line #212:   }
		BlankLine,   // line #213:
		OtherMethod, // line #214:   test_localVariableInference_transitive_topLevel_inferred_lexical() async {
		OtherMethod, // line #215:     await resolveTestCode(r'''
		OtherMethod, // line #216: final x = 3;
		OtherMethod, // line #217: main() {
		OtherMethod, // line #218:   var v = x;
		OtherMethod, // line #219:   v; // marker
		OtherMethod, // line #220: }
		OtherMethod, // line #221: ''');
		OtherMethod, // line #222:     assertType(findElement.localVar('v').type, 'int');
		OtherMethod, // line #223:     assertType(findNode.simple('v; // marker'), 'int');
		OtherMethod, // line #224:   }
		BlankLine,   // line #225:
		OtherMethod, // line #226:   test_localVariableInference_transitive_toplevel_inferred_reversed() async {
		OtherMethod, // line #227:     await resolveTestCode(r'''
		OtherMethod, // line #228: main() {
		OtherMethod, // line #229:   var v = x;
		OtherMethod, // line #230:   v; // marker
		OtherMethod, // line #231: }
		OtherMethod, // line #232: final x = 3;
		OtherMethod, // line #233: ''');
		OtherMethod, // line #234:     assertType(findElement.localVar('v').type, 'int');
		OtherMethod, // line #235:     assertType(findNode.simple('v; // marker'), 'int');
		OtherMethod, // line #236:   }
		BlankLine,   // line #237:
		OtherMethod, // line #238:   test_localVariableInference_transitive_topLevel_lexical() async {
		OtherMethod, // line #239:     await resolveTestCode(r'''
		OtherMethod, // line #240: int x = 3;
		OtherMethod, // line #241: main() {
		OtherMethod, // line #242:   var v = x;
		OtherMethod, // line #243:   v; // marker
		OtherMethod, // line #244: }
		OtherMethod, // line #245: ''');
		OtherMethod, // line #246:     assertType(findElement.localVar('v').type, 'int');
		OtherMethod, // line #247:     assertType(findNode.simple('v; // marker'), 'int');
		OtherMethod, // line #248:   }
		BlankLine,   // line #249:
		OtherMethod, // line #250:   test_localVariableInference_transitive_topLevel_reversed() async {
		OtherMethod, // line #251:     await resolveTestCode(r'''
		OtherMethod, // line #252: main() {
		OtherMethod, // line #253:   var v = x;
		OtherMethod, // line #254:   v; // marker
		OtherMethod, // line #255: }
		OtherMethod, // line #256: int x = 3;
		OtherMethod, // line #257: ''');
		OtherMethod, // line #258:     assertType(findElement.localVar('v').type, 'int');
		OtherMethod, // line #259:     assertType(findNode.simple('v; // marker'), 'int');
		OtherMethod, // line #260:   }
		BlankLine,   // line #261:
	}

	runParsePhase(t, opts, source, want)
	// runFullStylizer(t, opts, source, wantSource, want)
}
