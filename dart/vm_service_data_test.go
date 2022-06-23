/*
Copyright © 2022 Glenn M. Lewis

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

var wantVMServiceClass01 = []EntityType{
	Unknown,     // line #244: {
	OtherMethod, // line #245:   /// Returns the stream for a given stream id.
	OtherMethod, // line #246:   ///
	OtherMethod, // line #247:   /// This is not a part of the spec, but is needed for both the client and
	OtherMethod, // line #248:   /// server to get access to the real event streams.
	OtherMethod, // line #249:   Stream<Event> onEvent(String streamId);
	BlankLine,   // line #250:
	OtherMethod, // line #251:   /// Handler for calling extra service extensions.
	OtherMethod, // line #252:   Future<Response> callServiceExtension(String method,
	OtherMethod, // line #253:       {String isolateId, Map args});
	BlankLine,   // line #254:
	OtherMethod, // line #255:   /// The `addBreakpoint` RPC is used to add a breakpoint at a specific line of
	OtherMethod, // line #256:   /// some script.
	OtherMethod, // line #257:   ///
	OtherMethod, // line #258:   /// The `scriptId` parameter is used to specify the target script.
	OtherMethod, // line #259:   ///
	OtherMethod, // line #260:   /// The `line` parameter is used to specify the target line for the
	OtherMethod, // line #261:   /// breakpoint. If there are multiple possible breakpoints on the target line,
	OtherMethod, // line #262:   /// then the VM will place the breakpoint at the location which would execute
	OtherMethod, // line #263:   /// soonest. If it is not possible to set a breakpoint at the target line, the
	OtherMethod, // line #264:   /// breakpoint will be added at the next possible breakpoint location within
	OtherMethod, // line #265:   /// the same function.
	OtherMethod, // line #266:   ///
	OtherMethod, // line #267:   /// The `column` parameter may be optionally specified. This is useful for
	OtherMethod, // line #268:   /// targeting a specific breakpoint on a line with multiple possible
	OtherMethod, // line #269:   /// breakpoints.
	OtherMethod, // line #270:   ///
	OtherMethod, // line #271:   /// If no breakpoint is possible at that line, the `102` (Cannot add
	OtherMethod, // line #272:   /// breakpoint) [RPC error] code is returned.
	OtherMethod, // line #273:   ///
	OtherMethod, // line #274:   /// Note that breakpoints are added and removed on a per-isolate basis.
	OtherMethod, // line #275:   ///
	OtherMethod, // line #276:   /// If `isolateId` refers to an isolate which has exited, then the `Collected`
	OtherMethod, // line #277:   /// [Sentinel] is returned.
	OtherMethod, // line #278:   ///
	OtherMethod, // line #279:   /// See [Breakpoint].
	OtherMethod, // line #280:   ///
	OtherMethod, // line #281:   /// This method will throw a [SentinelException] in the case a [Sentinel] is
	OtherMethod, // line #282:   /// returned.
	OtherMethod, // line #283:   Future<Breakpoint> addBreakpoint(
	OtherMethod, // line #284:     String isolateId,
	OtherMethod, // line #285:     String scriptId,
	OtherMethod, // line #286:     int line, {
	OtherMethod, // line #287:     int column,
	OtherMethod, // line #288:   });
	BlankLine,   // line #289:
	OtherMethod, // line #290:   /// The `addBreakpoint` RPC is used to add a breakpoint at a specific line of
	OtherMethod, // line #291:   /// some script. This RPC is useful when a script has not yet been assigned an
	OtherMethod, // line #292:   /// id, for example, if a script is in a deferred library which has not yet
	OtherMethod, // line #293:   /// been loaded.
	OtherMethod, // line #294:   ///
	OtherMethod, // line #295:   /// The `scriptUri` parameter is used to specify the target script.
	OtherMethod, // line #296:   ///
	OtherMethod, // line #297:   /// The `line` parameter is used to specify the target line for the
	OtherMethod, // line #298:   /// breakpoint. If there are multiple possible breakpoints on the target line,
	OtherMethod, // line #299:   /// then the VM will place the breakpoint at the location which would execute
	OtherMethod, // line #300:   /// soonest. If it is not possible to set a breakpoint at the target line, the
	OtherMethod, // line #301:   /// breakpoint will be added at the next possible breakpoint location within
	OtherMethod, // line #302:   /// the same function.
	OtherMethod, // line #303:   ///
	OtherMethod, // line #304:   /// The `column` parameter may be optionally specified. This is useful for
	OtherMethod, // line #305:   /// targeting a specific breakpoint on a line with multiple possible
	OtherMethod, // line #306:   /// breakpoints.
	OtherMethod, // line #307:   ///
	OtherMethod, // line #308:   /// If no breakpoint is possible at that line, the `102` (Cannot add
	OtherMethod, // line #309:   /// breakpoint) [RPC error] code is returned.
	OtherMethod, // line #310:   ///
	OtherMethod, // line #311:   /// Note that breakpoints are added and removed on a per-isolate basis.
	OtherMethod, // line #312:   ///
	OtherMethod, // line #313:   /// If `isolateId` refers to an isolate which has exited, then the `Collected`
	OtherMethod, // line #314:   /// [Sentinel] is returned.
	OtherMethod, // line #315:   ///
	OtherMethod, // line #316:   /// See [Breakpoint].
	OtherMethod, // line #317:   ///
	OtherMethod, // line #318:   /// This method will throw a [SentinelException] in the case a [Sentinel] is
	OtherMethod, // line #319:   /// returned.
	OtherMethod, // line #320:   Future<Breakpoint> addBreakpointWithScriptUri(
	OtherMethod, // line #321:     String isolateId,
	OtherMethod, // line #322:     String scriptUri,
	OtherMethod, // line #323:     int line, {
	OtherMethod, // line #324:     int column,
	OtherMethod, // line #325:   });
	BlankLine,   // line #326:
	OtherMethod, // line #327:   /// The `addBreakpointAtEntry` RPC is used to add a breakpoint at the
	OtherMethod, // line #328:   /// entrypoint of some function.
	OtherMethod, // line #329:   ///
	OtherMethod, // line #330:   /// If no breakpoint is possible at the function entry, the `102` (Cannot add
	OtherMethod, // line #331:   /// breakpoint) [RPC error] code is returned.
	OtherMethod, // line #332:   ///
	OtherMethod, // line #333:   /// If `isolateId` refers to an isolate which has exited, then the `Collected`
	OtherMethod, // line #334:   /// [Sentinel] is returned.
	OtherMethod, // line #335:   ///
	OtherMethod, // line #336:   /// See [Breakpoint].
	OtherMethod, // line #337:   ///
	OtherMethod, // line #338:   /// Note that breakpoints are added and removed on a per-isolate basis.
	OtherMethod, // line #339:   ///
	OtherMethod, // line #340:   /// This method will throw a [SentinelException] in the case a [Sentinel] is
	OtherMethod, // line #341:   /// returned.
	OtherMethod, // line #342:   Future<Breakpoint> addBreakpointAtEntry(String isolateId, String functionId);
	BlankLine,   // line #343:
	OtherMethod, // line #344:   /// Clears all CPU profiling samples.
	OtherMethod, // line #345:   ///
	OtherMethod, // line #346:   /// If `isolateId` refers to an isolate which has exited, then the `Collected`
	OtherMethod, // line #347:   /// [Sentinel] is returned.
	OtherMethod, // line #348:   ///
	OtherMethod, // line #349:   /// See [Success].
	OtherMethod, // line #350:   ///
	OtherMethod, // line #351:   /// This method will throw a [SentinelException] in the case a [Sentinel] is
	OtherMethod, // line #352:   /// returned.
	OtherMethod, // line #353:   Future<Success> clearCpuSamples(String isolateId);
	BlankLine,   // line #354:
	OtherMethod, // line #355:   /// Clears all VM timeline events.
	OtherMethod, // line #356:   ///
	OtherMethod, // line #357:   /// See [Success].
	OtherMethod, // line #358:   Future<Success> clearVMTimeline();
	BlankLine,   // line #359:
	OtherMethod, // line #360:   /// The `invoke` RPC is used to perform regular method invocation on some
	OtherMethod, // line #361:   /// receiver, as if by dart:mirror's ObjectMirror.invoke. Note this does not
	OtherMethod, // line #362:   /// provide a way to perform getter, setter or constructor invocation.
	OtherMethod, // line #363:   ///
	OtherMethod, // line #364:   /// `targetId` may refer to a [Library], [Class], or [Instance].
	OtherMethod, // line #365:   ///
	OtherMethod, // line #366:   /// Each elements of `argumentId` may refer to an [Instance].
	OtherMethod, // line #367:   ///
	OtherMethod, // line #368:   /// If `disableBreakpoints` is provided and set to true, any breakpoints hit
	OtherMethod, // line #369:   /// as a result of this invocation are ignored, including pauses resulting
	OtherMethod, // line #370:   /// from a call to `debugger()` from `dart:developer`. Defaults to false if
	OtherMethod, // line #371:   /// not provided.
	OtherMethod, // line #372:   ///
	OtherMethod, // line #373:   /// If `targetId` or any element of `argumentIds` is a temporary id which has
	OtherMethod, // line #374:   /// expired, then the `Expired` [Sentinel] is returned.
	OtherMethod, // line #375:   ///
	OtherMethod, // line #376:   /// If `targetId` or any element of `argumentIds` refers to an object which
	OtherMethod, // line #377:   /// has been collected by the VM's garbage collector, then the `Collected`
	OtherMethod, // line #378:   /// [Sentinel] is returned.
	OtherMethod, // line #379:   ///
	OtherMethod, // line #380:   /// If `isolateId` refers to an isolate which has exited, then the `Collected`
	OtherMethod, // line #381:   /// [Sentinel] is returned.
	OtherMethod, // line #382:   ///
	OtherMethod, // line #383:   /// If invocation triggers a failed compilation then [RPC error] 113
	OtherMethod, // line #384:   /// "Expression compilation error" is returned.
	OtherMethod, // line #385:   ///
	OtherMethod, // line #386:   /// If a runtime error occurs while evaluating the invocation, an [ErrorRef]
	OtherMethod, // line #387:   /// reference will be returned.
	OtherMethod, // line #388:   ///
	OtherMethod, // line #389:   /// If the invocation is evaluated successfully, an [InstanceRef] reference
	OtherMethod, // line #390:   /// will be returned.
	OtherMethod, // line #391:   ///
	OtherMethod, // line #392:   /// The return value can be one of [InstanceRef] or [ErrorRef].
	OtherMethod, // line #393:   ///
	OtherMethod, // line #394:   /// This method will throw a [SentinelException] in the case a [Sentinel] is
	OtherMethod, // line #395:   /// returned.
	OtherMethod, // line #396:   Future<Response> invoke(
	OtherMethod, // line #397:     String isolateId,
	OtherMethod, // line #398:     String targetId,
	OtherMethod, // line #399:     String selector,
	OtherMethod, // line #400:     List<String> argumentIds, {
	OtherMethod, // line #401:     bool disableBreakpoints,
	OtherMethod, // line #402:   });
	BlankLine,   // line #403:
	OtherMethod, // line #404:   /// The `evaluate` RPC is used to evaluate an expression in the context of
	OtherMethod, // line #405:   /// some target.
	OtherMethod, // line #406:   ///
	OtherMethod, // line #407:   /// `targetId` may refer to a [Library], [Class], or [Instance].
	OtherMethod, // line #408:   ///
	OtherMethod, // line #409:   /// If `targetId` is a temporary id which has expired, then the `Expired`
	OtherMethod, // line #410:   /// [Sentinel] is returned.
	OtherMethod, // line #411:   ///
	OtherMethod, // line #412:   /// If `targetId` refers to an object which has been collected by the VM's
	OtherMethod, // line #413:   /// garbage collector, then the `Collected` [Sentinel] is returned.
	OtherMethod, // line #414:   ///
	OtherMethod, // line #415:   /// If `isolateId` refers to an isolate which has exited, then the `Collected`
	OtherMethod, // line #416:   /// [Sentinel] is returned.
	OtherMethod, // line #417:   ///
	OtherMethod, // line #418:   /// If `scope` is provided, it should be a map from identifiers to object ids.
	OtherMethod, // line #419:   /// These bindings will be added to the scope in which the expression is
	OtherMethod, // line #420:   /// evaluated, which is a child scope of the class or library for
	OtherMethod, // line #421:   /// instance/class or library targets respectively. This means bindings
	OtherMethod, // line #422:   /// provided in `scope` may shadow instance members, class members and
	OtherMethod, // line #423:   /// top-level members.
	OtherMethod, // line #424:   ///
	OtherMethod, // line #425:   /// If `disableBreakpoints` is provided and set to true, any breakpoints hit
	OtherMethod, // line #426:   /// as a result of this evaluation are ignored. Defaults to false if not
	OtherMethod, // line #427:   /// provided.
	OtherMethod, // line #428:   ///
	OtherMethod, // line #429:   /// If the expression fails to parse and compile, then [RPC error] 113
	OtherMethod, // line #430:   /// "Expression compilation error" is returned.
	OtherMethod, // line #431:   ///
	OtherMethod, // line #432:   /// If an error occurs while evaluating the expression, an [ErrorRef]
	OtherMethod, // line #433:   /// reference will be returned.
	OtherMethod, // line #434:   ///
	OtherMethod, // line #435:   /// If the expression is evaluated successfully, an [InstanceRef] reference
	OtherMethod, // line #436:   /// will be returned.
	OtherMethod, // line #437:   ///
	OtherMethod, // line #438:   /// The return value can be one of [InstanceRef] or [ErrorRef].
	OtherMethod, // line #439:   ///
	OtherMethod, // line #440:   /// This method will throw a [SentinelException] in the case a [Sentinel] is
	OtherMethod, // line #441:   /// returned.
	OtherMethod, // line #442:   Future<Response> evaluate(
	OtherMethod, // line #443:     String isolateId,
	OtherMethod, // line #444:     String targetId,
	OtherMethod, // line #445:     String expression, {
	OtherMethod, // line #446:     Map<String, String> scope,
	OtherMethod, // line #447:     bool disableBreakpoints,
	OtherMethod, // line #448:   });
	BlankLine,   // line #449:
	OtherMethod, // line #450:   /// The `evaluateInFrame` RPC is used to evaluate an expression in the context
	OtherMethod, // line #451:   /// of a particular stack frame. `frameIndex` is the index of the desired
	OtherMethod, // line #452:   /// [Frame], with an index of `0` indicating the top (most recent) frame.
	OtherMethod, // line #453:   ///
	OtherMethod, // line #454:   /// If `scope` is provided, it should be a map from identifiers to object ids.
	OtherMethod, // line #455:   /// These bindings will be added to the scope in which the expression is
	OtherMethod, // line #456:   /// evaluated, which is a child scope of the frame's current scope. This means
	OtherMethod, // line #457:   /// bindings provided in `scope` may shadow instance members, class members,
	OtherMethod, // line #458:   /// top-level members, parameters and locals.
	OtherMethod, // line #459:   ///
	OtherMethod, // line #460:   /// If `disableBreakpoints` is provided and set to true, any breakpoints hit
	OtherMethod, // line #461:   /// as a result of this evaluation are ignored. Defaults to false if not
	OtherMethod, // line #462:   /// provided.
	OtherMethod, // line #463:   ///
	OtherMethod, // line #464:   /// If the expression fails to parse and compile, then [RPC error] 113
	OtherMethod, // line #465:   /// "Expression compilation error" is returned.
	OtherMethod, // line #466:   ///
	OtherMethod, // line #467:   /// If an error occurs while evaluating the expression, an [ErrorRef]
	OtherMethod, // line #468:   /// reference will be returned.
	OtherMethod, // line #469:   ///
	OtherMethod, // line #470:   /// If the expression is evaluated successfully, an [InstanceRef] reference
	OtherMethod, // line #471:   /// will be returned.
	OtherMethod, // line #472:   ///
	OtherMethod, // line #473:   /// If `isolateId` refers to an isolate which has exited, then the `Collected`
	OtherMethod, // line #474:   /// [Sentinel] is returned.
	OtherMethod, // line #475:   ///
	OtherMethod, // line #476:   /// The return value can be one of [InstanceRef] or [ErrorRef].
	OtherMethod, // line #477:   ///
	OtherMethod, // line #478:   /// This method will throw a [SentinelException] in the case a [Sentinel] is
	OtherMethod, // line #479:   /// returned.
	OtherMethod, // line #480:   Future<Response> evaluateInFrame(
	OtherMethod, // line #481:     String isolateId,
	OtherMethod, // line #482:     int frameIndex,
	OtherMethod, // line #483:     String expression, {
	OtherMethod, // line #484:     Map<String, String> scope,
	OtherMethod, // line #485:     bool disableBreakpoints,
	OtherMethod, // line #486:   });
	BlankLine,   // line #487:
	OtherMethod, // line #488:   /// The `getAllocationProfile` RPC is used to retrieve allocation information
	OtherMethod, // line #489:   /// for a given isolate.
	OtherMethod, // line #490:   ///
	OtherMethod, // line #491:   /// If `reset` is provided and is set to true, the allocation accumulators
	OtherMethod, // line #492:   /// will be reset before collecting allocation information.
	OtherMethod, // line #493:   ///
	OtherMethod, // line #494:   /// If `gc` is provided and is set to true, a garbage collection will be
	OtherMethod, // line #495:   /// attempted before collecting allocation information. There is no guarantee
	OtherMethod, // line #496:   /// that a garbage collection will be actually be performed.
	OtherMethod, // line #497:   ///
	OtherMethod, // line #498:   /// If `isolateId` refers to an isolate which has exited, then the `Collected`
	OtherMethod, // line #499:   /// [Sentinel] is returned.
	OtherMethod, // line #500:   ///
	OtherMethod, // line #501:   /// This method will throw a [SentinelException] in the case a [Sentinel] is
	OtherMethod, // line #502:   /// returned.
	OtherMethod, // line #503:   Future<AllocationProfile> getAllocationProfile(String isolateId,
	OtherMethod, // line #504:       {bool reset, bool gc});
	BlankLine,   // line #505:
	OtherMethod, // line #506:   /// The `getClassList` RPC is used to retrieve a `ClassList` containing all
	OtherMethod, // line #507:   /// classes for an isolate based on the isolate's `isolateId`.
	OtherMethod, // line #508:   ///
	OtherMethod, // line #509:   /// If `isolateId` refers to an isolate which has exited, then the `Collected`
	OtherMethod, // line #510:   /// [Sentinel] is returned.
	OtherMethod, // line #511:   ///
	OtherMethod, // line #512:   /// See [ClassList].
	OtherMethod, // line #513:   ///
	OtherMethod, // line #514:   /// This method will throw a [SentinelException] in the case a [Sentinel] is
	OtherMethod, // line #515:   /// returned.
	OtherMethod, // line #516:   Future<ClassList> getClassList(String isolateId);
	BlankLine,   // line #517:
	OtherMethod, // line #518:   /// The `getCpuSamples` RPC is used to retrieve samples collected by the CPU
	OtherMethod, // line #519:   /// profiler. Only samples collected in the time range `[timeOriginMicros,
	OtherMethod, // line #520:   /// timeOriginMicros + timeExtentMicros]` will be reported.
	OtherMethod, // line #521:   ///
	OtherMethod, // line #522:   /// If the profiler is disabled, an [RPC error] response will be returned.
	OtherMethod, // line #523:   ///
	OtherMethod, // line #524:   /// If `isolateId` refers to an isolate which has exited, then the `Collected`
	OtherMethod, // line #525:   /// [Sentinel] is returned.
	OtherMethod, // line #526:   ///
	OtherMethod, // line #527:   /// See [CpuSamples].
	OtherMethod, // line #528:   ///
	OtherMethod, // line #529:   /// This method will throw a [SentinelException] in the case a [Sentinel] is
	OtherMethod, // line #530:   /// returned.
	OtherMethod, // line #531:   Future<CpuSamples> getCpuSamples(
	OtherMethod, // line #532:       String isolateId, int timeOriginMicros, int timeExtentMicros);
	BlankLine,   // line #533:
	OtherMethod, // line #534:   /// The `getFlagList` RPC returns a list of all command line flags in the VM
	OtherMethod, // line #535:   /// along with their current values.
	OtherMethod, // line #536:   ///
	OtherMethod, // line #537:   /// See [FlagList].
	OtherMethod, // line #538:   Future<FlagList> getFlagList();
	BlankLine,   // line #539:
	OtherMethod, // line #540:   /// Returns a set of inbound references to the object specified by `targetId`.
	OtherMethod, // line #541:   /// Up to `limit` references will be returned.
	OtherMethod, // line #542:   ///
	OtherMethod, // line #543:   /// The order of the references is undefined (i.e., not related to allocation
	OtherMethod, // line #544:   /// order) and unstable (i.e., multiple invocations of this method against the
	OtherMethod, // line #545:   /// same object can give different answers even if no Dart code has executed
	OtherMethod, // line #546:   /// between the invocations).
	OtherMethod, // line #547:   ///
	OtherMethod, // line #548:   /// The references may include multiple `objectId`s that designate the same
	OtherMethod, // line #549:   /// object.
	OtherMethod, // line #550:   ///
	OtherMethod, // line #551:   /// The references may include objects that are unreachable but have not yet
	OtherMethod, // line #552:   /// been garbage collected.
	OtherMethod, // line #553:   ///
	OtherMethod, // line #554:   /// If `targetId` is a temporary id which has expired, then the `Expired`
	OtherMethod, // line #555:   /// [Sentinel] is returned.
	OtherMethod, // line #556:   ///
	OtherMethod, // line #557:   /// If `targetId` refers to an object which has been collected by the VM's
	OtherMethod, // line #558:   /// garbage collector, then the `Collected` [Sentinel] is returned.
	OtherMethod, // line #559:   ///
	OtherMethod, // line #560:   /// If `isolateId` refers to an isolate which has exited, then the `Collected`
	OtherMethod, // line #561:   /// [Sentinel] is returned.
	OtherMethod, // line #562:   ///
	OtherMethod, // line #563:   /// See [InboundReferences].
	OtherMethod, // line #564:   ///
	OtherMethod, // line #565:   /// This method will throw a [SentinelException] in the case a [Sentinel] is
	OtherMethod, // line #566:   /// returned.
	OtherMethod, // line #567:   Future<InboundReferences> getInboundReferences(
	OtherMethod, // line #568:       String isolateId, String targetId, int limit);
	BlankLine,   // line #569:
	OtherMethod, // line #570:   /// The `getInstances` RPC is used to retrieve a set of instances which are of
	OtherMethod, // line #571:   /// a specific class. This does not include instances of subclasses of the
	OtherMethod, // line #572:   /// given class.
	OtherMethod, // line #573:   ///
	OtherMethod, // line #574:   /// The order of the instances is undefined (i.e., not related to allocation
	OtherMethod, // line #575:   /// order) and unstable (i.e., multiple invocations of this method against the
	OtherMethod, // line #576:   /// same class can give different answers even if no Dart code has executed
	OtherMethod, // line #577:   /// between the invocations).
	OtherMethod, // line #578:   ///
	OtherMethod, // line #579:   /// The set of instances may include objects that are unreachable but have not
	OtherMethod, // line #580:   /// yet been garbage collected.
	OtherMethod, // line #581:   ///
	OtherMethod, // line #582:   /// `objectId` is the ID of the `Class` to retrieve instances for. `objectId`
	OtherMethod, // line #583:   /// must be the ID of a `Class`, otherwise an [RPC error] is returned.
	OtherMethod, // line #584:   ///
	OtherMethod, // line #585:   /// `limit` is the maximum number of instances to be returned.
	OtherMethod, // line #586:   ///
	OtherMethod, // line #587:   /// If `isolateId` refers to an isolate which has exited, then the `Collected`
	OtherMethod, // line #588:   /// [Sentinel] is returned.
	OtherMethod, // line #589:   ///
	OtherMethod, // line #590:   /// See [InstanceSet].
	OtherMethod, // line #591:   ///
	OtherMethod, // line #592:   /// This method will throw a [SentinelException] in the case a [Sentinel] is
	OtherMethod, // line #593:   /// returned.
	OtherMethod, // line #594:   Future<InstanceSet> getInstances(
	OtherMethod, // line #595:       String isolateId, String objectId, int limit);
	BlankLine,   // line #596:
	OtherMethod, // line #597:   /// The `getIsolate` RPC is used to lookup an `Isolate` object by its `id`.
	OtherMethod, // line #598:   ///
	OtherMethod, // line #599:   /// If `isolateId` refers to an isolate which has exited, then the `Collected`
	OtherMethod, // line #600:   /// [Sentinel] is returned.
	OtherMethod, // line #601:   ///
	OtherMethod, // line #602:   /// See [Isolate].
	OtherMethod, // line #603:   ///
	OtherMethod, // line #604:   /// This method will throw a [SentinelException] in the case a [Sentinel] is
	OtherMethod, // line #605:   /// returned.
	OtherMethod, // line #606:   Future<Isolate> getIsolate(String isolateId);
	BlankLine,   // line #607:
	OtherMethod, // line #608:   /// The `getIsolateGroup` RPC is used to lookup an `IsolateGroup` object by
	OtherMethod, // line #609:   /// its `id`.
	OtherMethod, // line #610:   ///
	OtherMethod, // line #611:   /// If `isolateGroupId` refers to an isolate group which has exited, then the
	OtherMethod, // line #612:   /// `Expired` [Sentinel] is returned.
	OtherMethod, // line #613:   ///
	OtherMethod, // line #614:   /// `IsolateGroup` `id` is an opaque identifier that can be fetched from an
	OtherMethod, // line #615:   /// `IsolateGroup`. List of active `IsolateGroup`'s, for example, is available
	OtherMethod, // line #616:   /// on `VM` object.
	OtherMethod, // line #617:   ///
	OtherMethod, // line #618:   /// See [IsolateGroup], [VM].
	OtherMethod, // line #619:   ///
	OtherMethod, // line #620:   /// This method will throw a [SentinelException] in the case a [Sentinel] is
	OtherMethod, // line #621:   /// returned.
	OtherMethod, // line #622:   Future<IsolateGroup> getIsolateGroup(String isolateGroupId);
	BlankLine,   // line #623:
	OtherMethod, // line #624:   /// The `getMemoryUsage` RPC is used to lookup an isolate's memory usage
	OtherMethod, // line #625:   /// statistics by its `id`.
	OtherMethod, // line #626:   ///
	OtherMethod, // line #627:   /// If `isolateId` refers to an isolate which has exited, then the `Collected`
	OtherMethod, // line #628:   /// [Sentinel] is returned.
	OtherMethod, // line #629:   ///
	OtherMethod, // line #630:   /// See [Isolate].
	OtherMethod, // line #631:   ///
	OtherMethod, // line #632:   /// This method will throw a [SentinelException] in the case a [Sentinel] is
	OtherMethod, // line #633:   /// returned.
	OtherMethod, // line #634:   Future<MemoryUsage> getMemoryUsage(String isolateId);
	BlankLine,   // line #635:
	OtherMethod, // line #636:   /// The `getIsolateGroupMemoryUsage` RPC is used to lookup an isolate group's
	OtherMethod, // line #637:   /// memory usage statistics by its `id`.
	OtherMethod, // line #638:   ///
	OtherMethod, // line #639:   /// If `isolateGroupId` refers to an isolate group which has exited, then the
	OtherMethod, // line #640:   /// `Expired` [Sentinel] is returned.
	OtherMethod, // line #641:   ///
	OtherMethod, // line #642:   /// See [IsolateGroup].
	OtherMethod, // line #643:   ///
	OtherMethod, // line #644:   /// This method will throw a [SentinelException] in the case a [Sentinel] is
	OtherMethod, // line #645:   /// returned.
	OtherMethod, // line #646:   Future<MemoryUsage> getIsolateGroupMemoryUsage(String isolateGroupId);
	BlankLine,   // line #647:
	OtherMethod, // line #648:   /// The `getScripts` RPC is used to retrieve a `ScriptList` containing all
	OtherMethod, // line #649:   /// scripts for an isolate based on the isolate's `isolateId`.
	OtherMethod, // line #650:   ///
	OtherMethod, // line #651:   /// If `isolateId` refers to an isolate which has exited, then the `Collected`
	OtherMethod, // line #652:   /// [Sentinel] is returned.
	OtherMethod, // line #653:   ///
	OtherMethod, // line #654:   /// See [ScriptList].
	OtherMethod, // line #655:   ///
	OtherMethod, // line #656:   /// This method will throw a [SentinelException] in the case a [Sentinel] is
	OtherMethod, // line #657:   /// returned.
	OtherMethod, // line #658:   Future<ScriptList> getScripts(String isolateId);
	BlankLine,   // line #659:
	OtherMethod, // line #660:   /// The `getObject` RPC is used to lookup an `object` from some isolate by its
	OtherMethod, // line #661:   /// `id`.
	OtherMethod, // line #662:   ///
	OtherMethod, // line #663:   /// If `objectId` is a temporary id which has expired, then the `Expired`
	OtherMethod, // line #664:   /// [Sentinel] is returned.
	OtherMethod, // line #665:   ///
	OtherMethod, // line #666:   /// If `isolateId` refers to an isolate which has exited, then the `Collected`
	OtherMethod, // line #667:   /// [Sentinel] is returned.
	OtherMethod, // line #668:   ///
	OtherMethod, // line #669:   /// If `objectId` refers to a heap object which has been collected by the VM's
	OtherMethod, // line #670:   /// garbage collector, then the `Collected` [Sentinel] is returned.
	OtherMethod, // line #671:   ///
	OtherMethod, // line #672:   /// If `objectId` refers to a non-heap object which has been deleted, then the
	OtherMethod, // line #673:   /// `Collected` [Sentinel] is returned.
	OtherMethod, // line #674:   ///
	OtherMethod, // line #675:   /// If the object handle has not expired and the object has not been
	OtherMethod, // line #676:   /// collected, then an [Obj] will be returned.
	OtherMethod, // line #677:   ///
	OtherMethod, // line #678:   /// The `offset` and `count` parameters are used to request subranges of
	OtherMethod, // line #679:   /// Instance objects with the kinds: String, List, Map, Uint8ClampedList,
	OtherMethod, // line #680:   /// Uint8List, Uint16List, Uint32List, Uint64List, Int8List, Int16List,
	OtherMethod, // line #681:   /// Int32List, Int64List, Flooat32List, Float64List, Inst32x3List,
	OtherMethod, // line #682:   /// Float32x4List, and Float64x2List. These parameters are otherwise ignored.
	OtherMethod, // line #683:   ///
	OtherMethod, // line #684:   /// This method will throw a [SentinelException] in the case a [Sentinel] is
	OtherMethod, // line #685:   /// returned.
	OtherMethod, // line #686:   Future<Obj> getObject(
	OtherMethod, // line #687:     String isolateId,
	OtherMethod, // line #688:     String objectId, {
	OtherMethod, // line #689:     int offset,
	OtherMethod, // line #690:     int count,
	OtherMethod, // line #691:   });
	BlankLine,   // line #692:
	OtherMethod, // line #693:   /// The `getPorts` RPC is used to retrieve the list of `ReceivePort` instances
	OtherMethod, // line #694:   /// for a given isolate.
	OtherMethod, // line #695:   ///
	OtherMethod, // line #696:   /// See [PortList].
	OtherMethod, // line #697:   Future<PortList> getPorts(String isolateId);
	BlankLine,   // line #698:
	OtherMethod, // line #699:   /// The `getRetainingPath` RPC is used to lookup a path from an object
	OtherMethod, // line #700:   /// specified by `targetId` to a GC root (i.e., the object which is preventing
	OtherMethod, // line #701:   /// this object from being garbage collected).
	OtherMethod, // line #702:   ///
	OtherMethod, // line #703:   /// If `isolateId` refers to an isolate which has exited, then the `Collected`
	OtherMethod, // line #704:   /// [Sentinel] is returned.
	OtherMethod, // line #705:   ///
	OtherMethod, // line #706:   /// If `targetId` refers to a heap object which has been collected by the VM's
	OtherMethod, // line #707:   /// garbage collector, then the `Collected` [Sentinel] is returned.
	OtherMethod, // line #708:   ///
	OtherMethod, // line #709:   /// If `targetId` refers to a non-heap object which has been deleted, then the
	OtherMethod, // line #710:   /// `Collected` [Sentinel] is returned.
	OtherMethod, // line #711:   ///
	OtherMethod, // line #712:   /// If the object handle has not expired and the object has not been
	OtherMethod, // line #713:   /// collected, then an [RetainingPath] will be returned.
	OtherMethod, // line #714:   ///
	OtherMethod, // line #715:   /// The `limit` parameter specifies the maximum path length to be reported as
	OtherMethod, // line #716:   /// part of the retaining path. If a path is longer than `limit`, it will be
	OtherMethod, // line #717:   /// truncated at the root end of the path.
	OtherMethod, // line #718:   ///
	OtherMethod, // line #719:   /// See [RetainingPath].
	OtherMethod, // line #720:   ///
	OtherMethod, // line #721:   /// This method will throw a [SentinelException] in the case a [Sentinel] is
	OtherMethod, // line #722:   /// returned.
	OtherMethod, // line #723:   Future<RetainingPath> getRetainingPath(
	OtherMethod, // line #724:       String isolateId, String targetId, int limit);
	BlankLine,   // line #725:
	OtherMethod, // line #726:   /// Returns a description of major uses of memory known to the VM.
	OtherMethod, // line #727:   ///
	OtherMethod, // line #728:   /// Adding or removing buckets is considered a backwards-compatible change for
	OtherMethod, // line #729:   /// the purposes of versioning. A client must gracefully handle the removal or
	OtherMethod, // line #730:   /// addition of any bucket.
	OtherMethod, // line #731:   Future<ProcessMemoryUsage> getProcessMemoryUsage();
	BlankLine,   // line #732:
	OtherMethod, // line #733:   /// The `getStack` RPC is used to retrieve the current execution stack and
	OtherMethod, // line #734:   /// message queue for an isolate. The isolate does not need to be paused.
	OtherMethod, // line #735:   ///
	OtherMethod, // line #736:   /// If `limit` is provided, up to `limit` frames from the top of the stack
	OtherMethod, // line #737:   /// will be returned. If the stack depth is smaller than `limit` the entire
	OtherMethod, // line #738:   /// stack is returned. Note: this limit also applies to the
	OtherMethod, // line #739:   /// `asyncCausalFrames` and `awaiterFrames` stack representations in the
	OtherMethod, // line #740:   /// `Stack` response.
	OtherMethod, // line #741:   ///
	OtherMethod, // line #742:   /// If `isolateId` refers to an isolate which has exited, then the `Collected`
	OtherMethod, // line #743:   /// [Sentinel] is returned.
	OtherMethod, // line #744:   ///
	OtherMethod, // line #745:   /// See [Stack].
	OtherMethod, // line #746:   ///
	OtherMethod, // line #747:   /// This method will throw a [SentinelException] in the case a [Sentinel] is
	OtherMethod, // line #748:   /// returned.
	OtherMethod, // line #749:   Future<Stack> getStack(String isolateId, {int limit});
	BlankLine,   // line #750:
	OtherMethod, // line #751:   /// The `getSupportedProtocols` RPC is used to determine which protocols are
	OtherMethod, // line #752:   /// supported by the current server.
	OtherMethod, // line #753:   ///
	OtherMethod, // line #754:   /// The result of this call should be intercepted by any middleware that
	OtherMethod, // line #755:   /// extends the core VM service protocol and should add its own protocol to
	OtherMethod, // line #756:   /// the list of protocols before forwarding the response to the client.
	OtherMethod, // line #757:   ///
	OtherMethod, // line #758:   /// See [ProtocolList].
	OtherMethod, // line #759:   Future<ProtocolList> getSupportedProtocols();
	BlankLine,   // line #760:
	OtherMethod, // line #761:   /// The `getSourceReport` RPC is used to generate a set of reports tied to
	OtherMethod, // line #762:   /// source locations in an isolate.
	OtherMethod, // line #763:   ///
	OtherMethod, // line #764:   /// The `reports` parameter is used to specify which reports should be
	OtherMethod, // line #765:   /// generated. The `reports` parameter is a list, which allows multiple
	OtherMethod, // line #766:   /// reports to be generated simultaneously from a consistent isolate state.
	OtherMethod, // line #767:   /// The `reports` parameter is allowed to be empty (this might be used to
	OtherMethod, // line #768:   /// force compilation of a particular subrange of some script).
	OtherMethod, // line #769:   ///
	OtherMethod, // line #770:   /// The available report kinds are:
	OtherMethod, // line #771:   ///
	OtherMethod, // line #772:   /// report kind | meaning
	OtherMethod, // line #773:   /// ----------- | -------
	OtherMethod, // line #774:   /// Coverage | Provide code coverage information
	OtherMethod, // line #775:   /// PossibleBreakpoints | Provide a list of token positions which correspond
	OtherMethod, // line #776:   /// to possible breakpoints.
	OtherMethod, // line #777:   ///
	OtherMethod, // line #778:   /// The `scriptId` parameter is used to restrict the report to a particular
	OtherMethod, // line #779:   /// script. When analyzing a particular script, either or both of the
	OtherMethod, // line #780:   /// `tokenPos` and `endTokenPos` parameters may be provided to restrict the
	OtherMethod, // line #781:   /// analysis to a subrange of a script (for example, these can be used to
	OtherMethod, // line #782:   /// restrict the report to the range of a particular class or function).
	OtherMethod, // line #783:   ///
	OtherMethod, // line #784:   /// If the `scriptId` parameter is not provided then the reports are generated
	OtherMethod, // line #785:   /// for all loaded scripts and the `tokenPos` and `endTokenPos` parameters are
	OtherMethod, // line #786:   /// disallowed.
	OtherMethod, // line #787:   ///
	OtherMethod, // line #788:   /// The `forceCompilation` parameter can be used to force compilation of all
	OtherMethod, // line #789:   /// functions in the range of the report. Forcing compilation can cause a
	OtherMethod, // line #790:   /// compilation error, which could terminate the running Dart program. If this
	OtherMethod, // line #791:   /// parameter is not provided, it is considered to have the value `false`.
	OtherMethod, // line #792:   ///
	OtherMethod, // line #793:   /// If `isolateId` refers to an isolate which has exited, then the `Collected`
	OtherMethod, // line #794:   /// [Sentinel] is returned.
	OtherMethod, // line #795:   ///
	OtherMethod, // line #796:   /// See [SourceReport].
	OtherMethod, // line #797:   ///
	OtherMethod, // line #798:   /// This method will throw a [SentinelException] in the case a [Sentinel] is
	OtherMethod, // line #799:   /// returned.
	OtherMethod, // line #800:   Future<SourceReport> getSourceReport(
	OtherMethod, // line #801:     String isolateId,
	OtherMethod, // line #802:     /*List<SourceReportKind>*/
	OtherMethod, // line #803:     List<String> reports, {
	OtherMethod, // line #804:     String scriptId,
	OtherMethod, // line #805:     int tokenPos,
	OtherMethod, // line #806:     int endTokenPos,
	OtherMethod, // line #807:     bool forceCompile,
	OtherMethod, // line #808:   });
	BlankLine,   // line #809:
	OtherMethod, // line #810:   /// The `getVersion` RPC is used to determine what version of the Service
	OtherMethod, // line #811:   /// Protocol is served by a VM.
	OtherMethod, // line #812:   ///
	OtherMethod, // line #813:   /// See [Version].
	OtherMethod, // line #814:   Future<Version> getVersion();
	BlankLine,   // line #815:
	OtherMethod, // line #816:   /// The `getVM` RPC returns global information about a Dart virtual machine.
	OtherMethod, // line #817:   ///
	OtherMethod, // line #818:   /// See [VM].
	OtherMethod, // line #819:   Future<VM> getVM();
	BlankLine,   // line #820:
	OtherMethod, // line #821:   /// The `getVMTimeline` RPC is used to retrieve an object which contains VM
	OtherMethod, // line #822:   /// timeline events.
	OtherMethod, // line #823:   ///
	OtherMethod, // line #824:   /// The `timeOriginMicros` parameter is the beginning of the time range used
	OtherMethod, // line #825:   /// to filter timeline events. It uses the same monotonic clock as
	OtherMethod, // line #826:   /// dart:developer's `Timeline.now` and the VM embedding API's
	OtherMethod, // line #827:   /// `Dart_TimelineGetMicros`. See [getVMTimelineMicros] for access to this
	OtherMethod, // line #828:   /// clock through the service protocol.
	OtherMethod, // line #829:   ///
	OtherMethod, // line #830:   /// The `timeExtentMicros` parameter specifies how large the time range used
	OtherMethod, // line #831:   /// to filter timeline events should be.
	OtherMethod, // line #832:   ///
	OtherMethod, // line #833:   /// For example, given `timeOriginMicros` and `timeExtentMicros`, only
	OtherMethod, // line #834:   /// timeline events from the following time range will be returned:
	OtherMethod, // line #835:   /// `(timeOriginMicros, timeOriginMicros + timeExtentMicros)`.
	OtherMethod, // line #836:   ///
	OtherMethod, // line #837:   /// If `getVMTimeline` is invoked while the current recorder is one of Fuchsia
	OtherMethod, // line #838:   /// or Macos or Systrace, an [RPC error] with error code `114`, `invalid
	OtherMethod, // line #839:   /// timeline request`, will be returned as timeline events are handled by the
	OtherMethod, // line #840:   /// OS in these modes.
	OtherMethod, // line #841:   Future<Timeline> getVMTimeline({int timeOriginMicros, int timeExtentMicros});
	BlankLine,   // line #842:
	OtherMethod, // line #843:   /// The `getVMTimelineFlags` RPC returns information about the current VM
	OtherMethod, // line #844:   /// timeline configuration.
	OtherMethod, // line #845:   ///
	OtherMethod, // line #846:   /// To change which timeline streams are currently enabled, see
	OtherMethod, // line #847:   /// [setVMTimelineFlags].
	OtherMethod, // line #848:   ///
	OtherMethod, // line #849:   /// See [TimelineFlags].
	OtherMethod, // line #850:   Future<TimelineFlags> getVMTimelineFlags();
	BlankLine,   // line #851:
	OtherMethod, // line #852:   /// The `getVMTimelineMicros` RPC returns the current time stamp from the
	OtherMethod, // line #853:   /// clock used by the timeline, similar to `Timeline.now` in `dart:developer`
	OtherMethod, // line #854:   /// and `Dart_TimelineGetMicros` in the VM embedding API.
	OtherMethod, // line #855:   ///
	OtherMethod, // line #856:   /// See [Timestamp] and [getVMTimeline].
	OtherMethod, // line #857:   Future<Timestamp> getVMTimelineMicros();
	BlankLine,   // line #858:
	OtherMethod, // line #859:   /// The `pause` RPC is used to interrupt a running isolate. The RPC enqueues
	OtherMethod, // line #860:   /// the interrupt request and potentially returns before the isolate is
	OtherMethod, // line #861:   /// paused.
	OtherMethod, // line #862:   ///
	OtherMethod, // line #863:   /// When the isolate is paused an event will be sent on the `Debug` stream.
	OtherMethod, // line #864:   ///
	OtherMethod, // line #865:   /// If `isolateId` refers to an isolate which has exited, then the `Collected`
	OtherMethod, // line #866:   /// [Sentinel] is returned.
	OtherMethod, // line #867:   ///
	OtherMethod, // line #868:   /// See [Success].
	OtherMethod, // line #869:   ///
	OtherMethod, // line #870:   /// This method will throw a [SentinelException] in the case a [Sentinel] is
	OtherMethod, // line #871:   /// returned.
	OtherMethod, // line #872:   Future<Success> pause(String isolateId);
	BlankLine,   // line #873:
	OtherMethod, // line #874:   /// The `kill` RPC is used to kill an isolate as if by dart:isolate's
	OtherMethod, // line #875:   /// `Isolate.kill(IMMEDIATE)`.
	OtherMethod, // line #876:   ///
	OtherMethod, // line #877:   /// The isolate is killed regardless of whether it is paused or running.
	OtherMethod, // line #878:   ///
	OtherMethod, // line #879:   /// If `isolateId` refers to an isolate which has exited, then the `Collected`
	OtherMethod, // line #880:   /// [Sentinel] is returned.
	OtherMethod, // line #881:   ///
	OtherMethod, // line #882:   /// See [Success].
	OtherMethod, // line #883:   ///
	OtherMethod, // line #884:   /// This method will throw a [SentinelException] in the case a [Sentinel] is
	OtherMethod, // line #885:   /// returned.
	OtherMethod, // line #886:   Future<Success> kill(String isolateId);
	BlankLine,   // line #887:
	OtherMethod, // line #888:   /// Registers a service that can be invoked by other VM service clients, where
	OtherMethod, // line #889:   /// `service` is the name of the service to advertise and `alias` is an
	OtherMethod, // line #890:   /// alternative name for the registered service.
	OtherMethod, // line #891:   ///
	OtherMethod, // line #892:   /// Requests made to the new service will be forwarded to the client which
	OtherMethod, // line #893:   /// originally registered the service.
	OtherMethod, // line #894:   ///
	OtherMethod, // line #895:   /// See [Success].
	OtherMethod, // line #896:   Future<Success> registerService(String service, String alias);
	BlankLine,   // line #897:
	OtherMethod, // line #898:   /// The `reloadSources` RPC is used to perform a hot reload of an Isolate's
	OtherMethod, // line #899:   /// sources.
	OtherMethod, // line #900:   ///
	OtherMethod, // line #901:   /// if the `force` parameter is provided, it indicates that all of the
	OtherMethod, // line #902:   /// Isolate's sources should be reloaded regardless of modification time.
	OtherMethod, // line #903:   ///
	OtherMethod, // line #904:   /// if the `pause` parameter is provided, the isolate will pause immediately
	OtherMethod, // line #905:   /// after the reload.
	OtherMethod, // line #906:   ///
	OtherMethod, // line #907:   /// if the `rootLibUri` parameter is provided, it indicates the new uri to the
	OtherMethod, // line #908:   /// Isolate's root library.
	OtherMethod, // line #909:   ///
	OtherMethod, // line #910:   /// if the `packagesUri` parameter is provided, it indicates the new uri to
	OtherMethod, // line #911:   /// the Isolate's package map (.packages) file.
	OtherMethod, // line #912:   ///
	OtherMethod, // line #913:   /// If `isolateId` refers to an isolate which has exited, then the `Collected`
	OtherMethod, // line #914:   /// [Sentinel] is returned.
	OtherMethod, // line #915:   ///
	OtherMethod, // line #916:   /// This method will throw a [SentinelException] in the case a [Sentinel] is
	OtherMethod, // line #917:   /// returned.
	OtherMethod, // line #918:   Future<ReloadReport> reloadSources(
	OtherMethod, // line #919:     String isolateId, {
	OtherMethod, // line #920:     bool force,
	OtherMethod, // line #921:     bool pause,
	OtherMethod, // line #922:     String rootLibUri,
	OtherMethod, // line #923:     String packagesUri,
	OtherMethod, // line #924:   });
	BlankLine,   // line #925:
	OtherMethod, // line #926:   /// The `removeBreakpoint` RPC is used to remove a breakpoint by its `id`.
	OtherMethod, // line #927:   ///
	OtherMethod, // line #928:   /// Note that breakpoints are added and removed on a per-isolate basis.
	OtherMethod, // line #929:   ///
	OtherMethod, // line #930:   /// If `isolateId` refers to an isolate which has exited, then the `Collected`
	OtherMethod, // line #931:   /// [Sentinel] is returned.
	OtherMethod, // line #932:   ///
	OtherMethod, // line #933:   /// See [Success].
	OtherMethod, // line #934:   ///
	OtherMethod, // line #935:   /// This method will throw a [SentinelException] in the case a [Sentinel] is
	OtherMethod, // line #936:   /// returned.
	OtherMethod, // line #937:   Future<Success> removeBreakpoint(String isolateId, String breakpointId);
	BlankLine,   // line #938:
	OtherMethod, // line #939:   /// Requests a dump of the Dart heap of the given isolate.
	OtherMethod, // line #940:   ///
	OtherMethod, // line #941:   /// This method immediately returns success. The VM will then begin delivering
	OtherMethod, // line #942:   /// binary events on the `HeapSnapshot` event stream. The binary data in these
	OtherMethod, // line #943:   /// events, when concatenated together, conforms to the [SnapshotGraph] type.
	OtherMethod, // line #944:   /// The splitting of the SnapshotGraph into events can happen at any byte
	OtherMethod, // line #945:   /// offset.
	OtherMethod, // line #946:   ///
	OtherMethod, // line #947:   /// If `isolateId` refers to an isolate which has exited, then the `Collected`
	OtherMethod, // line #948:   /// [Sentinel] is returned.
	OtherMethod, // line #949:   ///
	OtherMethod, // line #950:   /// This method will throw a [SentinelException] in the case a [Sentinel] is
	OtherMethod, // line #951:   /// returned.
	OtherMethod, // line #952:   Future<Success> requestHeapSnapshot(String isolateId);
	BlankLine,   // line #953:
	OtherMethod, // line #954:   /// The `resume` RPC is used to resume execution of a paused isolate.
	OtherMethod, // line #955:   ///
	OtherMethod, // line #956:   /// If the `step` parameter is not provided, the program will resume regular
	OtherMethod, // line #957:   /// execution.
	OtherMethod, // line #958:   ///
	OtherMethod, // line #959:   /// If the `step` parameter is provided, it indicates what form of
	OtherMethod, // line #960:   /// single-stepping to use.
	OtherMethod, // line #961:   ///
	OtherMethod, // line #962:   /// step | meaning
	OtherMethod, // line #963:   /// ---- | -------
	OtherMethod, // line #964:   /// Into | Single step, entering function calls
	OtherMethod, // line #965:   /// Over | Single step, skipping over function calls
	OtherMethod, // line #966:   /// Out | Single step until the current function exits
	OtherMethod, // line #967:   /// Rewind | Immediately exit the top frame(s) without executing any code.
	OtherMethod, // line #968:   /// Isolate will be paused at the call of the last exited function.
	OtherMethod, // line #969:   ///
	OtherMethod, // line #970:   /// The `frameIndex` parameter is only used when the `step` parameter is
	OtherMethod, // line #971:   /// Rewind. It specifies the stack frame to rewind to. Stack frame 0 is the
	OtherMethod, // line #972:   /// currently executing function, so `frameIndex` must be at least 1.
	OtherMethod, // line #973:   ///
	OtherMethod, // line #974:   /// If the `frameIndex` parameter is not provided, it defaults to 1.
	OtherMethod, // line #975:   ///
	OtherMethod, // line #976:   /// If `isolateId` refers to an isolate which has exited, then the `Collected`
	OtherMethod, // line #977:   /// [Sentinel] is returned.
	OtherMethod, // line #978:   ///
	OtherMethod, // line #979:   /// See [Success], [StepOption].
	OtherMethod, // line #980:   ///
	OtherMethod, // line #981:   /// This method will throw a [SentinelException] in the case a [Sentinel] is
	OtherMethod, // line #982:   /// returned.
	OtherMethod, // line #983:   Future<Success> resume(String isolateId,
	OtherMethod, // line #984:       {/*StepOption*/ String step, int frameIndex});
	BlankLine,   // line #985:
	OtherMethod, // line #986:   /// The `setExceptionPauseMode` RPC is used to control if an isolate pauses
	OtherMethod, // line #987:   /// when an exception is thrown.
	OtherMethod, // line #988:   ///
	OtherMethod, // line #989:   /// mode | meaning
	OtherMethod, // line #990:   /// ---- | -------
	OtherMethod, // line #991:   /// None | Do not pause isolate on thrown exceptions
	OtherMethod, // line #992:   /// Unhandled | Pause isolate on unhandled exceptions
	OtherMethod, // line #993:   /// All  | Pause isolate on all thrown exceptions
	OtherMethod, // line #994:   ///
	OtherMethod, // line #995:   /// If `isolateId` refers to an isolate which has exited, then the `Collected`
	OtherMethod, // line #996:   /// [Sentinel] is returned.
	OtherMethod, // line #997:   ///
	OtherMethod, // line #998:   /// This method will throw a [SentinelException] in the case a [Sentinel] is
	OtherMethod, // line #999:   /// returned.
	OtherMethod, // line #1000:   Future<Success> setExceptionPauseMode(
	OtherMethod, // line #1001:       String isolateId, /*ExceptionPauseMode*/ String mode);
	BlankLine,   // line #1002:
	OtherMethod, // line #1003:   /// The `setFlag` RPC is used to set a VM flag at runtime. Returns an error if
	OtherMethod, // line #1004:   /// the named flag does not exist, the flag may not be set at runtime, or the
	OtherMethod, // line #1005:   /// value is of the wrong type for the flag.
	OtherMethod, // line #1006:   ///
	OtherMethod, // line #1007:   /// The following flags may be set at runtime:
	OtherMethod, // line #1008:   ///
	OtherMethod, // line #1009:   /// - pause_isolates_on_start
	OtherMethod, // line #1010:   /// - pause_isolates_on_exit
	OtherMethod, // line #1011:   /// - pause_isolates_on_unhandled_exceptions
	OtherMethod, // line #1012:   /// - profile_period
	OtherMethod, // line #1013:   /// - profiler
	OtherMethod, // line #1014:   ///
	OtherMethod, // line #1015:   /// Notes:
	OtherMethod, // line #1016:   ///
	OtherMethod, // line #1017:   /// - `profile_period` can be set to a minimum value of 50. Attempting to set
	OtherMethod, // line #1018:   /// `profile_period` to a lower value will result in a value of 50 being set.
	OtherMethod, // line #1019:   /// - Setting `profiler` will enable or disable the profiler depending on the
	OtherMethod, // line #1020:   /// provided value. If set to false when the profiler is already running, the
	OtherMethod, // line #1021:   /// profiler will be stopped but may not free its sample buffer depending on
	OtherMethod, // line #1022:   /// platform limitations.
	OtherMethod, // line #1023:   ///
	OtherMethod, // line #1024:   /// See [Success].
	OtherMethod, // line #1025:   ///
	OtherMethod, // line #1026:   /// The return value can be one of [Success] or [Error].
	OtherMethod, // line #1027:   Future<Response> setFlag(String name, String value);
	BlankLine,   // line #1028:
	OtherMethod, // line #1029:   /// The `setLibraryDebuggable` RPC is used to enable or disable whether
	OtherMethod, // line #1030:   /// breakpoints and stepping work for a given library.
	OtherMethod, // line #1031:   ///
	OtherMethod, // line #1032:   /// If `isolateId` refers to an isolate which has exited, then the `Collected`
	OtherMethod, // line #1033:   /// [Sentinel] is returned.
	OtherMethod, // line #1034:   ///
	OtherMethod, // line #1035:   /// See [Success].
	OtherMethod, // line #1036:   ///
	OtherMethod, // line #1037:   /// This method will throw a [SentinelException] in the case a [Sentinel] is
	OtherMethod, // line #1038:   /// returned.
	OtherMethod, // line #1039:   Future<Success> setLibraryDebuggable(
	OtherMethod, // line #1040:       String isolateId, String libraryId, bool isDebuggable);
	BlankLine,   // line #1041:
	OtherMethod, // line #1042:   /// The `setName` RPC is used to change the debugging name for an isolate.
	OtherMethod, // line #1043:   ///
	OtherMethod, // line #1044:   /// If `isolateId` refers to an isolate which has exited, then the `Collected`
	OtherMethod, // line #1045:   /// [Sentinel] is returned.
	OtherMethod, // line #1046:   ///
	OtherMethod, // line #1047:   /// See [Success].
	OtherMethod, // line #1048:   ///
	OtherMethod, // line #1049:   /// This method will throw a [SentinelException] in the case a [Sentinel] is
	OtherMethod, // line #1050:   /// returned.
	OtherMethod, // line #1051:   Future<Success> setName(String isolateId, String name);
	BlankLine,   // line #1052:
	OtherMethod, // line #1053:   /// The `setVMName` RPC is used to change the debugging name for the vm.
	OtherMethod, // line #1054:   ///
	OtherMethod, // line #1055:   /// See [Success].
	OtherMethod, // line #1056:   Future<Success> setVMName(String name);
	BlankLine,   // line #1057:
	OtherMethod, // line #1058:   /// The `setVMTimelineFlags` RPC is used to set which timeline streams are
	OtherMethod, // line #1059:   /// enabled.
	OtherMethod, // line #1060:   ///
	OtherMethod, // line #1061:   /// The `recordedStreams` parameter is the list of all timeline streams which
	OtherMethod, // line #1062:   /// are to be enabled. Streams not explicitly specified will be disabled.
	OtherMethod, // line #1063:   /// Invalid stream names are ignored.
	OtherMethod, // line #1064:   ///
	OtherMethod, // line #1065:   /// A `TimelineStreamSubscriptionsUpdate` event is sent on the `Timeline`
	OtherMethod, // line #1066:   /// stream as a result of invoking this RPC.
	OtherMethod, // line #1067:   ///
	OtherMethod, // line #1068:   /// To get the list of currently enabled timeline streams, see
	OtherMethod, // line #1069:   /// [getVMTimelineFlags].
	OtherMethod, // line #1070:   ///
	OtherMethod, // line #1071:   /// See [Success].
	OtherMethod, // line #1072:   Future<Success> setVMTimelineFlags(List<String> recordedStreams);
	BlankLine,   // line #1073:
	OtherMethod, // line #1074:   /// The `streamCancel` RPC cancels a stream subscription in the VM.
	OtherMethod, // line #1075:   ///
	OtherMethod, // line #1076:   /// If the client is not subscribed to the stream, the `104` (Stream not
	OtherMethod, // line #1077:   /// subscribed) [RPC error] code is returned.
	OtherMethod, // line #1078:   ///
	OtherMethod, // line #1079:   /// See [Success].
	OtherMethod, // line #1080:   Future<Success> streamCancel(String streamId);
	BlankLine,   // line #1081:
	OtherMethod, // line #1082:   /// The `streamListen` RPC subscribes to a stream in the VM. Once subscribed,
	OtherMethod, // line #1083:   /// the client will begin receiving events from the stream.
	OtherMethod, // line #1084:   ///
	OtherMethod, // line #1085:   /// If the client is already subscribed to the stream, the `103` (Stream
	OtherMethod, // line #1086:   /// already subscribed) [RPC error] code is returned.
	OtherMethod, // line #1087:   ///
	OtherMethod, // line #1088:   /// The `streamId` parameter may have the following published values:
	OtherMethod, // line #1089:   ///
	OtherMethod, // line #1090:   /// streamId | event types provided
	OtherMethod, // line #1091:   /// -------- | -----------
	OtherMethod, // line #1092:   /// VM | VMUpdate, VMFlagUpdate
	OtherMethod, // line #1093:   /// Isolate | IsolateStart, IsolateRunnable, IsolateExit, IsolateUpdate,
	OtherMethod, // line #1094:   /// IsolateReload, ServiceExtensionAdded
	OtherMethod, // line #1095:   /// Debug | PauseStart, PauseExit, PauseBreakpoint, PauseInterrupted,
	OtherMethod, // line #1096:   /// PauseException, PausePostRequest, Resume, BreakpointAdded,
	OtherMethod, // line #1097:   /// BreakpointResolved, BreakpointRemoved, Inspect, None
	OtherMethod, // line #1098:   /// GC | GC
	OtherMethod, // line #1099:   /// Extension | Extension
	OtherMethod, // line #1100:   /// Timeline | TimelineEvents, TimelineStreamsSubscriptionUpdate
	OtherMethod, // line #1101:   /// Logging | Logging
	OtherMethod, // line #1102:   /// Service | ServiceRegistered, ServiceUnregistered
	OtherMethod, // line #1103:   /// HeapSnapshot | HeapSnapshot
	OtherMethod, // line #1104:   ///
	OtherMethod, // line #1105:   /// Additionally, some embedders provide the `Stdout` and `Stderr` streams.
	OtherMethod, // line #1106:   /// These streams allow the client to subscribe to writes to stdout and
	OtherMethod, // line #1107:   /// stderr.
	OtherMethod, // line #1108:   ///
	OtherMethod, // line #1109:   /// streamId | event types provided
	OtherMethod, // line #1110:   /// -------- | -----------
	OtherMethod, // line #1111:   /// Stdout | WriteEvent
	OtherMethod, // line #1112:   /// Stderr | WriteEvent
	OtherMethod, // line #1113:   ///
	OtherMethod, // line #1114:   /// It is considered a `backwards compatible` change to add a new type of
	OtherMethod, // line #1115:   /// event to an existing stream. Clients should be written to handle this
	OtherMethod, // line #1116:   /// gracefully, perhaps by warning and ignoring.
	OtherMethod, // line #1117:   ///
	OtherMethod, // line #1118:   /// See [Success].
	OtherMethod, // line #1119:   Future<Success> streamListen(String streamId);
	BlankLine,   // line #1120:
}

var wantVMServiceClass02 = []EntityType{
	Unknown,                 // line #1122: {
	OtherMethod,             // line #1123:   Future<Map<String, Object>> get future => _completer.future;
	PrivateInstanceVariable, // line #1124:   final _completer = Completer<Map<String, Object>>();
	BlankLine,               // line #1125:
	InstanceVariable,        // line #1126:   final dynamic originalId;
	BlankLine,               // line #1127:
	MainConstructor,         // line #1128:   _PendingServiceRequest(this.originalId);
	BlankLine,               // line #1129:
	OtherMethod,             // line #1130:   void complete(Map<String, Object> response) {
	OtherMethod,             // line #1131:     response['id'] = originalId;
	OtherMethod,             // line #1132:     _completer.complete(response);
	OtherMethod,             // line #1133:   }
	BlankLine,               // line #1134:
}

var wantVMServiceClass03 = []EntityType{
	Unknown,                 // line #1142: {
	PrivateInstanceVariable, // line #1143:   final Stream<Map<String, Object>> _requestStream;
	PrivateInstanceVariable, // line #1144:   final StreamSink<Map<String, Object>> _responseSink;
	PrivateInstanceVariable, // line #1145:   final ServiceExtensionRegistry _serviceExtensionRegistry;
	PrivateInstanceVariable, // line #1146:   final VmServiceInterface _serviceImplementation;
	BlankLine,               // line #1147:
	PrivateInstanceVariable, // line #1148:   /// Used to create unique ids when acting as a proxy between clients.
	PrivateInstanceVariable, // line #1149:   int _nextServiceRequestId = 0;
	BlankLine,               // line #1150:
	PrivateInstanceVariable, // line #1151:   /// Manages streams for `streamListen` and `streamCancel` requests.
	PrivateInstanceVariable, // line #1152:   final _streamSubscriptions = <String, StreamSubscription>{};
	BlankLine,               // line #1153:
	OtherMethod,             // line #1154:   /// Completes when [_requestStream] is done.
	OtherMethod,             // line #1155:   Future get done => _doneCompleter.future;
	PrivateInstanceVariable, // line #1156:   final _doneCompleter = Completer<Null>();
	BlankLine,               // line #1157:
	PrivateInstanceVariable, // line #1158:   /// Pending service extension requests to this client by id.
	PrivateInstanceVariable, // line #1159:   final _pendingServiceExtensionRequests = <dynamic, _PendingServiceRequest>{};
	BlankLine,               // line #1160:
	MainConstructor,         // line #1161:   VmServerConnection(this._requestStream, this._responseSink,
	MainConstructor,         // line #1162:       this._serviceExtensionRegistry, this._serviceImplementation) {
	MainConstructor,         // line #1163:     _requestStream.listen(_delegateRequest, onDone: _doneCompleter.complete);
	MainConstructor,         // line #1164:     done.then(
	MainConstructor,         // line #1165:         (_) => _streamSubscriptions.values.forEach((sub) => sub.cancel()));
	MainConstructor,         // line #1166:   }
	BlankLine,               // line #1167:
	OtherMethod,             // line #1168:   /// Invoked when the current client has registered some extension, and
	OtherMethod,             // line #1169:   /// another client sends an RPC request for that extension.
	OtherMethod,             // line #1170:   ///
	OtherMethod,             // line #1171:   /// We don't attempt to do any serialization or deserialization of the
	OtherMethod,             // line #1172:   /// request or response in this case
	OtherMethod,             // line #1173:   Future<Map<String, Object>> _forwardServiceExtensionRequest(
	OtherMethod,             // line #1174:       Map<String, Object> request) {
	OtherMethod,             // line #1175:     var originalId = request['id'];
	OtherMethod,             // line #1176:     request = Map.of(request);
	OtherMethod,             // line #1177:     // Modify the request ID to ensure we don't have conflicts between
	OtherMethod,             // line #1178:     // multiple clients ids.
	OtherMethod,             // line #1179:     var newId = '${_nextServiceRequestId++}:$originalId';
	OtherMethod,             // line #1180:     request['id'] = newId;
	OtherMethod,             // line #1181:     var pendingRequest = _PendingServiceRequest(originalId);
	OtherMethod,             // line #1182:     _pendingServiceExtensionRequests[newId] = pendingRequest;
	OtherMethod,             // line #1183:     _responseSink.add(request);
	OtherMethod,             // line #1184:     return pendingRequest.future;
	OtherMethod,             // line #1185:   }
	BlankLine,               // line #1186:
	OtherMethod,             // line #1187:   void _delegateRequest(Map<String, Object> request) async {
	OtherMethod,             // line #1188:     try {
	OtherMethod,             // line #1189:       var id = request['id'];
	OtherMethod,             // line #1190:       // Check if this is actually a response to a pending request.
	OtherMethod,             // line #1191:       if (_pendingServiceExtensionRequests.containsKey(id)) {
	OtherMethod,             // line #1192:         final pending = _pendingServiceExtensionRequests[id];
	OtherMethod,             // line #1193:         pending.complete(Map.of(request));
	OtherMethod,             // line #1194:         return;
	OtherMethod,             // line #1195:       }
	OtherMethod,             // line #1196:       var method = request['method'] as String;
	OtherMethod,             // line #1197:       if (method == null) {
	OtherMethod,             // line #1198:         throw RPCError(
	OtherMethod,             // line #1199:             null, RPCError.kInvalidRequest, 'Invalid Request', request);
	OtherMethod,             // line #1200:       }
	OtherMethod,             // line #1201:       var params = request['params'] as Map;
	OtherMethod,             // line #1202:       Response response;
	OtherMethod,             // line #1203:
	OtherMethod,             // line #1204:       switch (method) {
	OtherMethod,             // line #1205:         case 'registerService':
	OtherMethod,             // line #1206:           _serviceExtensionRegistry.registerExtension(params['service'], this);
	OtherMethod,             // line #1207:           response = Success();
	OtherMethod,             // line #1208:           break;
	OtherMethod,             // line #1209:         case 'addBreakpoint':
	OtherMethod,             // line #1210:           response = await _serviceImplementation.addBreakpoint(
	OtherMethod,             // line #1211:             params['isolateId'],
	OtherMethod,             // line #1212:             params['scriptId'],
	OtherMethod,             // line #1213:             params['line'],
	OtherMethod,             // line #1214:             column: params['column'],
	OtherMethod,             // line #1215:           );
	OtherMethod,             // line #1216:           break;
	OtherMethod,             // line #1217:         case 'addBreakpointWithScriptUri':
	OtherMethod,             // line #1218:           response = await _serviceImplementation.addBreakpointWithScriptUri(
	OtherMethod,             // line #1219:             params['isolateId'],
	OtherMethod,             // line #1220:             params['scriptUri'],
	OtherMethod,             // line #1221:             params['line'],
	OtherMethod,             // line #1222:             column: params['column'],
	OtherMethod,             // line #1223:           );
	OtherMethod,             // line #1224:           break;
	OtherMethod,             // line #1225:         case 'addBreakpointAtEntry':
	OtherMethod,             // line #1226:           response = await _serviceImplementation.addBreakpointAtEntry(
	OtherMethod,             // line #1227:             params['isolateId'],
	OtherMethod,             // line #1228:             params['functionId'],
	OtherMethod,             // line #1229:           );
	OtherMethod,             // line #1230:           break;
	OtherMethod,             // line #1231:         case 'clearCpuSamples':
	OtherMethod,             // line #1232:           response = await _serviceImplementation.clearCpuSamples(
	OtherMethod,             // line #1233:             params['isolateId'],
	OtherMethod,             // line #1234:           );
	OtherMethod,             // line #1235:           break;
	OtherMethod,             // line #1236:         case 'clearVMTimeline':
	OtherMethod,             // line #1237:           response = await _serviceImplementation.clearVMTimeline();
	OtherMethod,             // line #1238:           break;
	OtherMethod,             // line #1239:         case 'invoke':
	OtherMethod,             // line #1240:           response = await _serviceImplementation.invoke(
	OtherMethod,             // line #1241:             params['isolateId'],
	OtherMethod,             // line #1242:             params['targetId'],
	OtherMethod,             // line #1243:             params['selector'],
	OtherMethod,             // line #1244:             List<String>.from(params['argumentIds'] ?? []),
	OtherMethod,             // line #1245:             disableBreakpoints: params['disableBreakpoints'],
	OtherMethod,             // line #1246:           );
	OtherMethod,             // line #1247:           break;
	OtherMethod,             // line #1248:         case 'evaluate':
	OtherMethod,             // line #1249:           response = await _serviceImplementation.evaluate(
	OtherMethod,             // line #1250:             params['isolateId'],
	OtherMethod,             // line #1251:             params['targetId'],
	OtherMethod,             // line #1252:             params['expression'],
	OtherMethod,             // line #1253:             scope: params['scope']?.cast<String, String>(),
	OtherMethod,             // line #1254:             disableBreakpoints: params['disableBreakpoints'],
	OtherMethod,             // line #1255:           );
	OtherMethod,             // line #1256:           break;
	OtherMethod,             // line #1257:         case 'evaluateInFrame':
	OtherMethod,             // line #1258:           response = await _serviceImplementation.evaluateInFrame(
	OtherMethod,             // line #1259:             params['isolateId'],
	OtherMethod,             // line #1260:             params['frameIndex'],
	OtherMethod,             // line #1261:             params['expression'],
	OtherMethod,             // line #1262:             scope: params['scope']?.cast<String, String>(),
	OtherMethod,             // line #1263:             disableBreakpoints: params['disableBreakpoints'],
	OtherMethod,             // line #1264:           );
	OtherMethod,             // line #1265:           break;
	OtherMethod,             // line #1266:         case 'getAllocationProfile':
	OtherMethod,             // line #1267:           response = await _serviceImplementation.getAllocationProfile(
	OtherMethod,             // line #1268:             params['isolateId'],
	OtherMethod,             // line #1269:             reset: params['reset'],
	OtherMethod,             // line #1270:             gc: params['gc'],
	OtherMethod,             // line #1271:           );
	OtherMethod,             // line #1272:           break;
	OtherMethod,             // line #1273:         case 'getClassList':
	OtherMethod,             // line #1274:           response = await _serviceImplementation.getClassList(
	OtherMethod,             // line #1275:             params['isolateId'],
	OtherMethod,             // line #1276:           );
	OtherMethod,             // line #1277:           break;
	OtherMethod,             // line #1278:         case 'getCpuSamples':
	OtherMethod,             // line #1279:           response = await _serviceImplementation.getCpuSamples(
	OtherMethod,             // line #1280:             params['isolateId'],
	OtherMethod,             // line #1281:             params['timeOriginMicros'],
	OtherMethod,             // line #1282:             params['timeExtentMicros'],
	OtherMethod,             // line #1283:           );
	OtherMethod,             // line #1284:           break;
	OtherMethod,             // line #1285:         case 'getFlagList':
	OtherMethod,             // line #1286:           response = await _serviceImplementation.getFlagList();
	OtherMethod,             // line #1287:           break;
	OtherMethod,             // line #1288:         case 'getInboundReferences':
	OtherMethod,             // line #1289:           response = await _serviceImplementation.getInboundReferences(
	OtherMethod,             // line #1290:             params['isolateId'],
	OtherMethod,             // line #1291:             params['targetId'],
	OtherMethod,             // line #1292:             params['limit'],
	OtherMethod,             // line #1293:           );
	OtherMethod,             // line #1294:           break;
	OtherMethod,             // line #1295:         case 'getInstances':
	OtherMethod,             // line #1296:           response = await _serviceImplementation.getInstances(
	OtherMethod,             // line #1297:             params['isolateId'],
	OtherMethod,             // line #1298:             params['objectId'],
	OtherMethod,             // line #1299:             params['limit'],
	OtherMethod,             // line #1300:           );
	OtherMethod,             // line #1301:           break;
	OtherMethod,             // line #1302:         case 'getIsolate':
	OtherMethod,             // line #1303:           response = await _serviceImplementation.getIsolate(
	OtherMethod,             // line #1304:             params['isolateId'],
	OtherMethod,             // line #1305:           );
	OtherMethod,             // line #1306:           break;
	OtherMethod,             // line #1307:         case 'getIsolateGroup':
	OtherMethod,             // line #1308:           response = await _serviceImplementation.getIsolateGroup(
	OtherMethod,             // line #1309:             params['isolateGroupId'],
	OtherMethod,             // line #1310:           );
	OtherMethod,             // line #1311:           break;
	OtherMethod,             // line #1312:         case 'getMemoryUsage':
	OtherMethod,             // line #1313:           response = await _serviceImplementation.getMemoryUsage(
	OtherMethod,             // line #1314:             params['isolateId'],
	OtherMethod,             // line #1315:           );
	OtherMethod,             // line #1316:           break;
	OtherMethod,             // line #1317:         case 'getIsolateGroupMemoryUsage':
	OtherMethod,             // line #1318:           response = await _serviceImplementation.getIsolateGroupMemoryUsage(
	OtherMethod,             // line #1319:             params['isolateGroupId'],
	OtherMethod,             // line #1320:           );
	OtherMethod,             // line #1321:           break;
	OtherMethod,             // line #1322:         case 'getScripts':
	OtherMethod,             // line #1323:           response = await _serviceImplementation.getScripts(
	OtherMethod,             // line #1324:             params['isolateId'],
	OtherMethod,             // line #1325:           );
	OtherMethod,             // line #1326:           break;
	OtherMethod,             // line #1327:         case 'getObject':
	OtherMethod,             // line #1328:           response = await _serviceImplementation.getObject(
	OtherMethod,             // line #1329:             params['isolateId'],
	OtherMethod,             // line #1330:             params['objectId'],
	OtherMethod,             // line #1331:             offset: params['offset'],
	OtherMethod,             // line #1332:             count: params['count'],
	OtherMethod,             // line #1333:           );
	OtherMethod,             // line #1334:           break;
	OtherMethod,             // line #1335:         case 'getPorts':
	OtherMethod,             // line #1336:           response = await _serviceImplementation.getPorts(
	OtherMethod,             // line #1337:             params['isolateId'],
	OtherMethod,             // line #1338:           );
	OtherMethod,             // line #1339:           break;
	OtherMethod,             // line #1340:         case 'getRetainingPath':
	OtherMethod,             // line #1341:           response = await _serviceImplementation.getRetainingPath(
	OtherMethod,             // line #1342:             params['isolateId'],
	OtherMethod,             // line #1343:             params['targetId'],
	OtherMethod,             // line #1344:             params['limit'],
	OtherMethod,             // line #1345:           );
	OtherMethod,             // line #1346:           break;
	OtherMethod,             // line #1347:         case 'getProcessMemoryUsage':
	OtherMethod,             // line #1348:           response = await _serviceImplementation.getProcessMemoryUsage();
	OtherMethod,             // line #1349:           break;
	OtherMethod,             // line #1350:         case 'getStack':
	OtherMethod,             // line #1351:           response = await _serviceImplementation.getStack(
	OtherMethod,             // line #1352:             params['isolateId'],
	OtherMethod,             // line #1353:             limit: params['limit'],
	OtherMethod,             // line #1354:           );
	OtherMethod,             // line #1355:           break;
	OtherMethod,             // line #1356:         case 'getSupportedProtocols':
	OtherMethod,             // line #1357:           response = await _serviceImplementation.getSupportedProtocols();
	OtherMethod,             // line #1358:           break;
	OtherMethod,             // line #1359:         case 'getSourceReport':
	OtherMethod,             // line #1360:           response = await _serviceImplementation.getSourceReport(
	OtherMethod,             // line #1361:             params['isolateId'],
	OtherMethod,             // line #1362:             List<String>.from(params['reports'] ?? []),
	OtherMethod,             // line #1363:             scriptId: params['scriptId'],
	OtherMethod,             // line #1364:             tokenPos: params['tokenPos'],
	OtherMethod,             // line #1365:             endTokenPos: params['endTokenPos'],
	OtherMethod,             // line #1366:             forceCompile: params['forceCompile'],
	OtherMethod,             // line #1367:           );
	OtherMethod,             // line #1368:           break;
	OtherMethod,             // line #1369:         case 'getVersion':
	OtherMethod,             // line #1370:           response = await _serviceImplementation.getVersion();
	OtherMethod,             // line #1371:           break;
	OtherMethod,             // line #1372:         case 'getVM':
	OtherMethod,             // line #1373:           response = await _serviceImplementation.getVM();
	OtherMethod,             // line #1374:           break;
	OtherMethod,             // line #1375:         case 'getVMTimeline':
	OtherMethod,             // line #1376:           response = await _serviceImplementation.getVMTimeline(
	OtherMethod,             // line #1377:             timeOriginMicros: params['timeOriginMicros'],
	OtherMethod,             // line #1378:             timeExtentMicros: params['timeExtentMicros'],
	OtherMethod,             // line #1379:           );
	OtherMethod,             // line #1380:           break;
	OtherMethod,             // line #1381:         case 'getVMTimelineFlags':
	OtherMethod,             // line #1382:           response = await _serviceImplementation.getVMTimelineFlags();
	OtherMethod,             // line #1383:           break;
	OtherMethod,             // line #1384:         case 'getVMTimelineMicros':
	OtherMethod,             // line #1385:           response = await _serviceImplementation.getVMTimelineMicros();
	OtherMethod,             // line #1386:           break;
	OtherMethod,             // line #1387:         case 'pause':
	OtherMethod,             // line #1388:           response = await _serviceImplementation.pause(
	OtherMethod,             // line #1389:             params['isolateId'],
	OtherMethod,             // line #1390:           );
	OtherMethod,             // line #1391:           break;
	OtherMethod,             // line #1392:         case 'kill':
	OtherMethod,             // line #1393:           response = await _serviceImplementation.kill(
	OtherMethod,             // line #1394:             params['isolateId'],
	OtherMethod,             // line #1395:           );
	OtherMethod,             // line #1396:           break;
	OtherMethod,             // line #1397:         case 'reloadSources':
	OtherMethod,             // line #1398:           response = await _serviceImplementation.reloadSources(
	OtherMethod,             // line #1399:             params['isolateId'],
	OtherMethod,             // line #1400:             force: params['force'],
	OtherMethod,             // line #1401:             pause: params['pause'],
	OtherMethod,             // line #1402:             rootLibUri: params['rootLibUri'],
	OtherMethod,             // line #1403:             packagesUri: params['packagesUri'],
	OtherMethod,             // line #1404:           );
	OtherMethod,             // line #1405:           break;
	OtherMethod,             // line #1406:         case 'removeBreakpoint':
	OtherMethod,             // line #1407:           response = await _serviceImplementation.removeBreakpoint(
	OtherMethod,             // line #1408:             params['isolateId'],
	OtherMethod,             // line #1409:             params['breakpointId'],
	OtherMethod,             // line #1410:           );
	OtherMethod,             // line #1411:           break;
	OtherMethod,             // line #1412:         case 'requestHeapSnapshot':
	OtherMethod,             // line #1413:           response = await _serviceImplementation.requestHeapSnapshot(
	OtherMethod,             // line #1414:             params['isolateId'],
	OtherMethod,             // line #1415:           );
	OtherMethod,             // line #1416:           break;
	OtherMethod,             // line #1417:         case 'resume':
	OtherMethod,             // line #1418:           response = await _serviceImplementation.resume(
	OtherMethod,             // line #1419:             params['isolateId'],
	OtherMethod,             // line #1420:             step: params['step'],
	OtherMethod,             // line #1421:             frameIndex: params['frameIndex'],
	OtherMethod,             // line #1422:           );
	OtherMethod,             // line #1423:           break;
	OtherMethod,             // line #1424:         case 'setExceptionPauseMode':
	OtherMethod,             // line #1425:           response = await _serviceImplementation.setExceptionPauseMode(
	OtherMethod,             // line #1426:             params['isolateId'],
	OtherMethod,             // line #1427:             params['mode'],
	OtherMethod,             // line #1428:           );
	OtherMethod,             // line #1429:           break;
	OtherMethod,             // line #1430:         case 'setFlag':
	OtherMethod,             // line #1431:           response = await _serviceImplementation.setFlag(
	OtherMethod,             // line #1432:             params['name'],
	OtherMethod,             // line #1433:             params['value'],
	OtherMethod,             // line #1434:           );
	OtherMethod,             // line #1435:           break;
	OtherMethod,             // line #1436:         case 'setLibraryDebuggable':
	OtherMethod,             // line #1437:           response = await _serviceImplementation.setLibraryDebuggable(
	OtherMethod,             // line #1438:             params['isolateId'],
	OtherMethod,             // line #1439:             params['libraryId'],
	OtherMethod,             // line #1440:             params['isDebuggable'],
	OtherMethod,             // line #1441:           );
	OtherMethod,             // line #1442:           break;
	OtherMethod,             // line #1443:         case 'setName':
	OtherMethod,             // line #1444:           response = await _serviceImplementation.setName(
	OtherMethod,             // line #1445:             params['isolateId'],
	OtherMethod,             // line #1446:             params['name'],
	OtherMethod,             // line #1447:           );
	OtherMethod,             // line #1448:           break;
	OtherMethod,             // line #1449:         case 'setVMName':
	OtherMethod,             // line #1450:           response = await _serviceImplementation.setVMName(
	OtherMethod,             // line #1451:             params['name'],
	OtherMethod,             // line #1452:           );
	OtherMethod,             // line #1453:           break;
	OtherMethod,             // line #1454:         case 'setVMTimelineFlags':
	OtherMethod,             // line #1455:           response = await _serviceImplementation.setVMTimelineFlags(
	OtherMethod,             // line #1456:             List<String>.from(params['recordedStreams'] ?? []),
	OtherMethod,             // line #1457:           );
	OtherMethod,             // line #1458:           break;
	OtherMethod,             // line #1459:         case 'streamCancel':
	OtherMethod,             // line #1460:           var id = params['streamId'];
	OtherMethod,             // line #1461:           var existing = _streamSubscriptions.remove(id);
	OtherMethod,             // line #1462:           if (existing == null) {
	OtherMethod,             // line #1463:             throw RPCError.withDetails(
	OtherMethod,             // line #1464:               'streamCancel',
	OtherMethod,             // line #1465:               104,
	OtherMethod,             // line #1466:               'Stream not subscribed',
	OtherMethod,             // line #1467:               details: "The stream '$id' is not subscribed",
	OtherMethod,             // line #1468:             );
	OtherMethod,             // line #1469:           }
	OtherMethod,             // line #1470:           await existing.cancel();
	OtherMethod,             // line #1471:           response = Success();
	OtherMethod,             // line #1472:           break;
	OtherMethod,             // line #1473:         case 'streamListen':
	OtherMethod,             // line #1474:           var id = params['streamId'];
	OtherMethod,             // line #1475:           if (_streamSubscriptions.containsKey(id)) {
	OtherMethod,             // line #1476:             throw RPCError.withDetails(
	OtherMethod,             // line #1477:               'streamListen',
	OtherMethod,             // line #1478:               103,
	OtherMethod,             // line #1479:               'Stream already subscribed',
	OtherMethod,             // line #1480:               details: "The stream '$id' is already subscribed",
	OtherMethod,             // line #1481:             );
	OtherMethod,             // line #1482:           }
	OtherMethod,             // line #1483:
	OtherMethod,             // line #1484:           var stream = id == 'Service'
	OtherMethod,             // line #1485:               ? _serviceExtensionRegistry.onExtensionEvent
	OtherMethod,             // line #1486:               : _serviceImplementation.onEvent(id);
	OtherMethod,             // line #1487:           _streamSubscriptions[id] = stream.listen((e) {
	OtherMethod,             // line #1488:             _responseSink.add({
	OtherMethod,             // line #1489:               'jsonrpc': '2.0',
	OtherMethod,             // line #1490:               'method': 'streamNotify',
	OtherMethod,             // line #1491:               'params': {
	OtherMethod,             // line #1492:                 'streamId': id,
	OtherMethod,             // line #1493:                 'event': e.toJson(),
	OtherMethod,             // line #1494:               },
	OtherMethod,             // line #1495:             });
	OtherMethod,             // line #1496:           });
	OtherMethod,             // line #1497:           response = Success();
	OtherMethod,             // line #1498:           break;
	OtherMethod,             // line #1499:         default:
	OtherMethod,             // line #1500:           var registeredClient = _serviceExtensionRegistry.clientFor(method);
	OtherMethod,             // line #1501:           if (registeredClient != null) {
	OtherMethod,             // line #1502:             // Check for any client which has registered this extension, if we
	OtherMethod,             // line #1503:             // have one then delegate the request to that client.
	OtherMethod,             // line #1504:             _responseSink.add(await registeredClient
	OtherMethod,             // line #1505:                 ._forwardServiceExtensionRequest(request));
	OtherMethod,             // line #1506:             // Bail out early in this case, we are just acting as a proxy and
	OtherMethod,             // line #1507:             // never get a `Response` instance.
	OtherMethod,             // line #1508:             return;
	OtherMethod,             // line #1509:           } else if (method.startsWith('ext.')) {
	OtherMethod,             // line #1510:             // Remaining methods with `ext.` are assumed to be registered via
	OtherMethod,             // line #1511:             // dart:developer, which the service implementation handles.
	OtherMethod,             // line #1512:             var args = params == null ? null : Map.of(params);
	OtherMethod,             // line #1513:             var isolateId = args?.remove('isolateId');
	OtherMethod,             // line #1514:             response = await _serviceImplementation.callServiceExtension(method,
	OtherMethod,             // line #1515:                 isolateId: isolateId, args: args);
	OtherMethod,             // line #1516:           } else {
	OtherMethod,             // line #1517:             throw RPCError(
	OtherMethod,             // line #1518:                 method, RPCError.kMethodNotFound, 'Method not found', request);
	OtherMethod,             // line #1519:           }
	OtherMethod,             // line #1520:       }
	OtherMethod,             // line #1521:       if (response == null) {
	OtherMethod,             // line #1522:         throw StateError('Invalid null response from service');
	OtherMethod,             // line #1523:       }
	OtherMethod,             // line #1524:       _responseSink.add({
	OtherMethod,             // line #1525:         'jsonrpc': '2.0',
	OtherMethod,             // line #1526:         'id': id,
	OtherMethod,             // line #1527:         'result': response.toJson(),
	OtherMethod,             // line #1528:       });
	OtherMethod,             // line #1529:     } catch (e, st) {
	OtherMethod,             // line #1530:       var error = e is RPCError
	OtherMethod,             // line #1531:           ? e.toMap()
	OtherMethod,             // line #1532:           : {
	OtherMethod,             // line #1533:               'code': RPCError.kInternalError,
	OtherMethod,             // line #1534:               'message': '${request['method']}: $e',
	OtherMethod,             // line #1535:               'data': {'details': '$st'},
	OtherMethod,             // line #1536:             };
	OtherMethod,             // line #1537:       _responseSink.add({
	OtherMethod,             // line #1538:         'jsonrpc': '2.0',
	OtherMethod,             // line #1539:         'id': request['id'],
	OtherMethod,             // line #1540:         'error': error,
	OtherMethod,             // line #1541:       });
	OtherMethod,             // line #1542:     }
	OtherMethod,             // line #1543:   }
	BlankLine,               // line #1544:
}

var wantVMServiceClass04 = []EntityType{
	Unknown,                 // line #1546: {
	PrivateInstanceVariable, // line #1547:   StreamSubscription _streamSub;
	PrivateInstanceVariable, // line #1548:   Function _writeMessage;
	PrivateInstanceVariable, // line #1549:   int _id = 0;
	PrivateInstanceVariable, // line #1550:   Map<String, Completer> _completers = {};
	PrivateInstanceVariable, // line #1551:   Map<String, String> _methodCalls = {};
	PrivateInstanceVariable, // line #1552:   Map<String, ServiceCallback> _services = {};
	PrivateInstanceVariable, // line #1553:   Log _log;
	BlankLine,               // line #1554:
	PrivateInstanceVariable, // line #1555:   StreamController<String> _onSend = StreamController.broadcast(sync: true);
	PrivateInstanceVariable, // line #1556:   StreamController<String> _onReceive = StreamController.broadcast(sync: true);
	BlankLine,               // line #1557:
	PrivateInstanceVariable, // line #1558:   final Completer _onDoneCompleter = Completer();
	BlankLine,               // line #1559:
	PrivateInstanceVariable, // line #1560:   Map<String, StreamController<Event>> _eventControllers = {};
	BlankLine,               // line #1561:
	OtherMethod,             // line #1562:   StreamController<Event> _getEventController(String eventName) {
	OtherMethod,             // line #1563:     StreamController<Event> controller = _eventControllers[eventName];
	OtherMethod,             // line #1564:     if (controller == null) {
	OtherMethod,             // line #1565:       controller = StreamController.broadcast();
	OtherMethod,             // line #1566:       _eventControllers[eventName] = controller;
	OtherMethod,             // line #1567:     }
	OtherMethod,             // line #1568:     return controller;
	OtherMethod,             // line #1569:   }
	BlankLine,               // line #1570:
	PrivateInstanceVariable, // line #1571:   DisposeHandler _disposeHandler;
	BlankLine,               // line #1572:
	MainConstructor,         // line #1573:   VmService(
	MainConstructor,         // line #1574:     Stream<dynamic> /*String|List<int>*/ inStream,
	MainConstructor,         // line #1575:     void writeMessage(String message), {
	MainConstructor,         // line #1576:     Log log,
	MainConstructor,         // line #1577:     DisposeHandler disposeHandler,
	MainConstructor,         // line #1578:     Future streamClosed,
	MainConstructor,         // line #1579:   }) {
	MainConstructor,         // line #1580:     _streamSub = inStream.listen(_processMessage,
	MainConstructor,         // line #1581:         onDone: () => _onDoneCompleter.complete());
	MainConstructor,         // line #1582:     _writeMessage = writeMessage;
	MainConstructor,         // line #1583:     _log = log == null ? _NullLog() : log;
	MainConstructor,         // line #1584:     _disposeHandler = disposeHandler;
	MainConstructor,         // line #1585:     streamClosed?.then((_) {
	MainConstructor,         // line #1586:       if (!_onDoneCompleter.isCompleted) {
	MainConstructor,         // line #1587:         _onDoneCompleter.complete();
	MainConstructor,         // line #1588:       }
	MainConstructor,         // line #1589:     });
	MainConstructor,         // line #1590:   }
	BlankLine,               // line #1591:
	OverrideMethod,          // line #1592:   @override
	OverrideMethod,          // line #1593:   Stream<Event> onEvent(String streamId) =>
	OverrideMethod,          // line #1594:       _getEventController(streamId).stream;
	BlankLine,               // line #1595:
	OtherMethod,             // line #1596:   // VMUpdate, VMFlagUpdate
	OtherMethod,             // line #1597:   Stream<Event> get onVMEvent => _getEventController('VM').stream;
	BlankLine,               // line #1598:
	OtherMethod,             // line #1599:   // IsolateStart, IsolateRunnable, IsolateExit, IsolateUpdate, IsolateReload, ServiceExtensionAdded
	OtherMethod,             // line #1600:   Stream<Event> get onIsolateEvent => _getEventController('Isolate').stream;
	BlankLine,               // line #1601:
	OtherMethod,             // line #1602:   // PauseStart, PauseExit, PauseBreakpoint, PauseInterrupted, PauseException, PausePostRequest, Resume, BreakpointAdded, BreakpointResolved, BreakpointRemoved, Inspect, None
	OtherMethod,             // line #1603:   Stream<Event> get onDebugEvent => _getEventController('Debug').stream;
	BlankLine,               // line #1604:
	OtherMethod,             // line #1605:   // GC
	OtherMethod,             // line #1606:   Stream<Event> get onGCEvent => _getEventController('GC').stream;
	BlankLine,               // line #1607:
	OtherMethod,             // line #1608:   // Extension
	OtherMethod,             // line #1609:   Stream<Event> get onExtensionEvent => _getEventController('Extension').stream;
	BlankLine,               // line #1610:
	OtherMethod,             // line #1611:   // TimelineEvents, TimelineStreamsSubscriptionUpdate
	OtherMethod,             // line #1612:   Stream<Event> get onTimelineEvent => _getEventController('Timeline').stream;
	BlankLine,               // line #1613:
	OtherMethod,             // line #1614:   // Logging
	OtherMethod,             // line #1615:   Stream<Event> get onLoggingEvent => _getEventController('Logging').stream;
	BlankLine,               // line #1616:
	OtherMethod,             // line #1617:   // ServiceRegistered, ServiceUnregistered
	OtherMethod,             // line #1618:   Stream<Event> get onServiceEvent => _getEventController('Service').stream;
	BlankLine,               // line #1619:
	OtherMethod,             // line #1620:   // HeapSnapshot
	OtherMethod,             // line #1621:   Stream<Event> get onHeapSnapshotEvent =>
	OtherMethod,             // line #1622:       _getEventController('HeapSnapshot').stream;
	BlankLine,               // line #1623:
	OtherMethod,             // line #1624:   // WriteEvent
	OtherMethod,             // line #1625:   Stream<Event> get onStdoutEvent => _getEventController('Stdout').stream;
	BlankLine,               // line #1626:
	OtherMethod,             // line #1627:   // WriteEvent
	OtherMethod,             // line #1628:   Stream<Event> get onStderrEvent => _getEventController('Stderr').stream;
	BlankLine,               // line #1629:
	OverrideMethod,          // line #1630:   @override
	OverrideMethod,          // line #1631:   Future<Breakpoint> addBreakpoint(
	OverrideMethod,          // line #1632:     String isolateId,
	OverrideMethod,          // line #1633:     String scriptId,
	OverrideMethod,          // line #1634:     int line, {
	OverrideMethod,          // line #1635:     int column,
	OverrideMethod,          // line #1636:   }) =>
	OverrideMethod,          // line #1637:       _call('addBreakpoint', {
	OverrideMethod,          // line #1638:         'isolateId': isolateId,
	OverrideMethod,          // line #1639:         'scriptId': scriptId,
	OverrideMethod,          // line #1640:         'line': line,
	OverrideMethod,          // line #1641:         if (column != null) 'column': column,
	OverrideMethod,          // line #1642:       });
	BlankLine,               // line #1643:
	OverrideMethod,          // line #1644:   @override
	OverrideMethod,          // line #1645:   Future<Breakpoint> addBreakpointWithScriptUri(
	OverrideMethod,          // line #1646:     String isolateId,
	OverrideMethod,          // line #1647:     String scriptUri,
	OverrideMethod,          // line #1648:     int line, {
	OverrideMethod,          // line #1649:     int column,
	OverrideMethod,          // line #1650:   }) =>
	OverrideMethod,          // line #1651:       _call('addBreakpointWithScriptUri', {
	OverrideMethod,          // line #1652:         'isolateId': isolateId,
	OverrideMethod,          // line #1653:         'scriptUri': scriptUri,
	OverrideMethod,          // line #1654:         'line': line,
	OverrideMethod,          // line #1655:         if (column != null) 'column': column,
	OverrideMethod,          // line #1656:       });
	BlankLine,               // line #1657:
	OverrideMethod,          // line #1658:   @override
	OverrideMethod,          // line #1659:   Future<Breakpoint> addBreakpointAtEntry(
	OverrideMethod,          // line #1660:           String isolateId, String functionId) =>
	OverrideMethod,          // line #1661:       _call('addBreakpointAtEntry',
	OverrideMethod,          // line #1662:           {'isolateId': isolateId, 'functionId': functionId});
	BlankLine,               // line #1663:
	OverrideMethod,          // line #1664:   @override
	OverrideMethod,          // line #1665:   Future<Success> clearCpuSamples(String isolateId) =>
	OverrideMethod,          // line #1666:       _call('clearCpuSamples', {'isolateId': isolateId});
	BlankLine,               // line #1667:
	OverrideMethod,          // line #1668:   @override
	OverrideMethod,          // line #1669:   Future<Success> clearVMTimeline() => _call('clearVMTimeline');
	BlankLine,               // line #1670:
	OverrideMethod,          // line #1671:   @override
	OverrideMethod,          // line #1672:   Future<Response> invoke(
	OverrideMethod,          // line #1673:     String isolateId,
	OverrideMethod,          // line #1674:     String targetId,
	OverrideMethod,          // line #1675:     String selector,
	OverrideMethod,          // line #1676:     List<String> argumentIds, {
	OverrideMethod,          // line #1677:     bool disableBreakpoints,
	OverrideMethod,          // line #1678:   }) =>
	OverrideMethod,          // line #1679:       _call('invoke', {
	OverrideMethod,          // line #1680:         'isolateId': isolateId,
	OverrideMethod,          // line #1681:         'targetId': targetId,
	OverrideMethod,          // line #1682:         'selector': selector,
	OverrideMethod,          // line #1683:         'argumentIds': argumentIds,
	OverrideMethod,          // line #1684:         if (disableBreakpoints != null)
	OverrideMethod,          // line #1685:           'disableBreakpoints': disableBreakpoints,
	OverrideMethod,          // line #1686:       });
	BlankLine,               // line #1687:
	OverrideMethod,          // line #1688:   @override
	OverrideMethod,          // line #1689:   Future<Response> evaluate(
	OverrideMethod,          // line #1690:     String isolateId,
	OverrideMethod,          // line #1691:     String targetId,
	OverrideMethod,          // line #1692:     String expression, {
	OverrideMethod,          // line #1693:     Map<String, String> scope,
	OverrideMethod,          // line #1694:     bool disableBreakpoints,
	OverrideMethod,          // line #1695:   }) =>
	OverrideMethod,          // line #1696:       _call('evaluate', {
	OverrideMethod,          // line #1697:         'isolateId': isolateId,
	OverrideMethod,          // line #1698:         'targetId': targetId,
	OverrideMethod,          // line #1699:         'expression': expression,
	OverrideMethod,          // line #1700:         if (scope != null) 'scope': scope,
	OverrideMethod,          // line #1701:         if (disableBreakpoints != null)
	OverrideMethod,          // line #1702:           'disableBreakpoints': disableBreakpoints,
	OverrideMethod,          // line #1703:       });
	BlankLine,               // line #1704:
	OverrideMethod,          // line #1705:   @override
	OverrideMethod,          // line #1706:   Future<Response> evaluateInFrame(
	OverrideMethod,          // line #1707:     String isolateId,
	OverrideMethod,          // line #1708:     int frameIndex,
	OverrideMethod,          // line #1709:     String expression, {
	OverrideMethod,          // line #1710:     Map<String, String> scope,
	OverrideMethod,          // line #1711:     bool disableBreakpoints,
	OverrideMethod,          // line #1712:   }) =>
	OverrideMethod,          // line #1713:       _call('evaluateInFrame', {
	OverrideMethod,          // line #1714:         'isolateId': isolateId,
	OverrideMethod,          // line #1715:         'frameIndex': frameIndex,
	OverrideMethod,          // line #1716:         'expression': expression,
	OverrideMethod,          // line #1717:         if (scope != null) 'scope': scope,
	OverrideMethod,          // line #1718:         if (disableBreakpoints != null)
	OverrideMethod,          // line #1719:           'disableBreakpoints': disableBreakpoints,
	OverrideMethod,          // line #1720:       });
	BlankLine,               // line #1721:
	OverrideMethod,          // line #1722:   @override
	OverrideMethod,          // line #1723:   Future<AllocationProfile> getAllocationProfile(String isolateId,
	OverrideMethod,          // line #1724:           {bool reset, bool gc}) =>
	OverrideMethod,          // line #1725:       _call('getAllocationProfile', {
	OverrideMethod,          // line #1726:         'isolateId': isolateId,
	OverrideMethod,          // line #1727:         if (reset != null && reset) 'reset': reset,
	OverrideMethod,          // line #1728:         if (gc != null && gc) 'gc': gc,
	OverrideMethod,          // line #1729:       });
	BlankLine,               // line #1730:
	OverrideMethod,          // line #1731:   @override
	OverrideMethod,          // line #1732:   Future<ClassList> getClassList(String isolateId) =>
	OverrideMethod,          // line #1733:       _call('getClassList', {'isolateId': isolateId});
	BlankLine,               // line #1734:
	OverrideMethod,          // line #1735:   @override
	OverrideMethod,          // line #1736:   Future<CpuSamples> getCpuSamples(
	OverrideMethod,          // line #1737:           String isolateId, int timeOriginMicros, int timeExtentMicros) =>
	OverrideMethod,          // line #1738:       _call('getCpuSamples', {
	OverrideMethod,          // line #1739:         'isolateId': isolateId,
	OverrideMethod,          // line #1740:         'timeOriginMicros': timeOriginMicros,
	OverrideMethod,          // line #1741:         'timeExtentMicros': timeExtentMicros
	OverrideMethod,          // line #1742:       });
	BlankLine,               // line #1743:
	OverrideMethod,          // line #1744:   @override
	OverrideMethod,          // line #1745:   Future<FlagList> getFlagList() => _call('getFlagList');
	BlankLine,               // line #1746:
	OverrideMethod,          // line #1747:   @override
	OverrideMethod,          // line #1748:   Future<InboundReferences> getInboundReferences(
	OverrideMethod,          // line #1749:           String isolateId, String targetId, int limit) =>
	OverrideMethod,          // line #1750:       _call('getInboundReferences',
	OverrideMethod,          // line #1751:           {'isolateId': isolateId, 'targetId': targetId, 'limit': limit});
	BlankLine,               // line #1752:
	OverrideMethod,          // line #1753:   @override
	OverrideMethod,          // line #1754:   Future<InstanceSet> getInstances(
	OverrideMethod,          // line #1755:           String isolateId, String objectId, int limit) =>
	OverrideMethod,          // line #1756:       _call('getInstances',
	OverrideMethod,          // line #1757:           {'isolateId': isolateId, 'objectId': objectId, 'limit': limit});
	BlankLine,               // line #1758:
	OverrideMethod,          // line #1759:   @override
	OverrideMethod,          // line #1760:   Future<Isolate> getIsolate(String isolateId) =>
	OverrideMethod,          // line #1761:       _call('getIsolate', {'isolateId': isolateId});
	BlankLine,               // line #1762:
	OverrideMethod,          // line #1763:   @override
	OverrideMethod,          // line #1764:   Future<IsolateGroup> getIsolateGroup(String isolateGroupId) =>
	OverrideMethod,          // line #1765:       _call('getIsolateGroup', {'isolateGroupId': isolateGroupId});
	BlankLine,               // line #1766:
	OverrideMethod,          // line #1767:   @override
	OverrideMethod,          // line #1768:   Future<MemoryUsage> getMemoryUsage(String isolateId) =>
	OverrideMethod,          // line #1769:       _call('getMemoryUsage', {'isolateId': isolateId});
	BlankLine,               // line #1770:
	OverrideMethod,          // line #1771:   @override
	OverrideMethod,          // line #1772:   Future<MemoryUsage> getIsolateGroupMemoryUsage(String isolateGroupId) =>
	OverrideMethod,          // line #1773:       _call('getIsolateGroupMemoryUsage', {'isolateGroupId': isolateGroupId});
	BlankLine,               // line #1774:
	OverrideMethod,          // line #1775:   @override
	OverrideMethod,          // line #1776:   Future<ScriptList> getScripts(String isolateId) =>
	OverrideMethod,          // line #1777:       _call('getScripts', {'isolateId': isolateId});
	BlankLine,               // line #1778:
	OverrideMethod,          // line #1779:   @override
	OverrideMethod,          // line #1780:   Future<Obj> getObject(
	OverrideMethod,          // line #1781:     String isolateId,
	OverrideMethod,          // line #1782:     String objectId, {
	OverrideMethod,          // line #1783:     int offset,
	OverrideMethod,          // line #1784:     int count,
	OverrideMethod,          // line #1785:   }) =>
	OverrideMethod,          // line #1786:       _call('getObject', {
	OverrideMethod,          // line #1787:         'isolateId': isolateId,
	OverrideMethod,          // line #1788:         'objectId': objectId,
	OverrideMethod,          // line #1789:         if (offset != null) 'offset': offset,
	OverrideMethod,          // line #1790:         if (count != null) 'count': count,
	OverrideMethod,          // line #1791:       });
	BlankLine,               // line #1792:
	OverrideMethod,          // line #1793:   @override
	OverrideMethod,          // line #1794:   Future<PortList> getPorts(String isolateId) =>
	OverrideMethod,          // line #1795:       _call('getPorts', {'isolateId': isolateId});
	BlankLine,               // line #1796:
	OverrideMethod,          // line #1797:   @override
	OverrideMethod,          // line #1798:   Future<RetainingPath> getRetainingPath(
	OverrideMethod,          // line #1799:           String isolateId, String targetId, int limit) =>
	OverrideMethod,          // line #1800:       _call('getRetainingPath',
	OverrideMethod,          // line #1801:           {'isolateId': isolateId, 'targetId': targetId, 'limit': limit});
	BlankLine,               // line #1802:
	OverrideMethod,          // line #1803:   @override
	OverrideMethod,          // line #1804:   Future<ProcessMemoryUsage> getProcessMemoryUsage() =>
	OverrideMethod,          // line #1805:       _call('getProcessMemoryUsage');
	BlankLine,               // line #1806:
	OverrideMethod,          // line #1807:   @override
	OverrideMethod,          // line #1808:   Future<Stack> getStack(String isolateId, {int limit}) => _call('getStack', {
	OverrideMethod,          // line #1809:         'isolateId': isolateId,
	OverrideMethod,          // line #1810:         if (limit != null) 'limit': limit,
	OverrideMethod,          // line #1811:       });
	BlankLine,               // line #1812:
	OverrideMethod,          // line #1813:   @override
	OverrideMethod,          // line #1814:   Future<ProtocolList> getSupportedProtocols() =>
	OverrideMethod,          // line #1815:       _call('getSupportedProtocols');
	BlankLine,               // line #1816:
	OverrideMethod,          // line #1817:   @override
	OverrideMethod,          // line #1818:   Future<SourceReport> getSourceReport(
	OverrideMethod,          // line #1819:     String isolateId,
	OverrideMethod,          // line #1820:     /*List<SourceReportKind>*/
	OverrideMethod,          // line #1821:     List<String> reports, {
	OverrideMethod,          // line #1822:     String scriptId,
	OverrideMethod,          // line #1823:     int tokenPos,
	OverrideMethod,          // line #1824:     int endTokenPos,
	OverrideMethod,          // line #1825:     bool forceCompile,
	OverrideMethod,          // line #1826:   }) =>
	OverrideMethod,          // line #1827:       _call('getSourceReport', {
	OverrideMethod,          // line #1828:         'isolateId': isolateId,
	OverrideMethod,          // line #1829:         'reports': reports,
	OverrideMethod,          // line #1830:         if (scriptId != null) 'scriptId': scriptId,
	OverrideMethod,          // line #1831:         if (tokenPos != null) 'tokenPos': tokenPos,
	OverrideMethod,          // line #1832:         if (endTokenPos != null) 'endTokenPos': endTokenPos,
	OverrideMethod,          // line #1833:         if (forceCompile != null) 'forceCompile': forceCompile,
	OverrideMethod,          // line #1834:       });
	BlankLine,               // line #1835:
	OverrideMethod,          // line #1836:   @override
	OverrideMethod,          // line #1837:   Future<Version> getVersion() => _call('getVersion');
	BlankLine,               // line #1838:
	OverrideMethod,          // line #1839:   @override
	OverrideMethod,          // line #1840:   Future<VM> getVM() => _call('getVM');
	BlankLine,               // line #1841:
	OverrideMethod,          // line #1842:   @override
	OverrideMethod,          // line #1843:   Future<Timeline> getVMTimeline(
	OverrideMethod,          // line #1844:           {int timeOriginMicros, int timeExtentMicros}) =>
	OverrideMethod,          // line #1845:       _call('getVMTimeline', {
	OverrideMethod,          // line #1846:         if (timeOriginMicros != null) 'timeOriginMicros': timeOriginMicros,
	OverrideMethod,          // line #1847:         if (timeExtentMicros != null) 'timeExtentMicros': timeExtentMicros,
	OverrideMethod,          // line #1848:       });
	BlankLine,               // line #1849:
	OverrideMethod,          // line #1850:   @override
	OverrideMethod,          // line #1851:   Future<TimelineFlags> getVMTimelineFlags() => _call('getVMTimelineFlags');
	BlankLine,               // line #1852:
	OverrideMethod,          // line #1853:   @override
	OverrideMethod,          // line #1854:   Future<Timestamp> getVMTimelineMicros() => _call('getVMTimelineMicros');
	BlankLine,               // line #1855:
	OverrideMethod,          // line #1856:   @override
	OverrideMethod,          // line #1857:   Future<Success> pause(String isolateId) =>
	OverrideMethod,          // line #1858:       _call('pause', {'isolateId': isolateId});
	BlankLine,               // line #1859:
	OverrideMethod,          // line #1860:   @override
	OverrideMethod,          // line #1861:   Future<Success> kill(String isolateId) =>
	OverrideMethod,          // line #1862:       _call('kill', {'isolateId': isolateId});
	BlankLine,               // line #1863:
	OverrideMethod,          // line #1864:   @override
	OverrideMethod,          // line #1865:   Future<Success> registerService(String service, String alias) =>
	OverrideMethod,          // line #1866:       _call('registerService', {'service': service, 'alias': alias});
	BlankLine,               // line #1867:
	OverrideMethod,          // line #1868:   @override
	OverrideMethod,          // line #1869:   Future<ReloadReport> reloadSources(
	OverrideMethod,          // line #1870:     String isolateId, {
	OverrideMethod,          // line #1871:     bool force,
	OverrideMethod,          // line #1872:     bool pause,
	OverrideMethod,          // line #1873:     String rootLibUri,
	OverrideMethod,          // line #1874:     String packagesUri,
	OverrideMethod,          // line #1875:   }) =>
	OverrideMethod,          // line #1876:       _call('reloadSources', {
	OverrideMethod,          // line #1877:         'isolateId': isolateId,
	OverrideMethod,          // line #1878:         if (force != null) 'force': force,
	OverrideMethod,          // line #1879:         if (pause != null) 'pause': pause,
	OverrideMethod,          // line #1880:         if (rootLibUri != null) 'rootLibUri': rootLibUri,
	OverrideMethod,          // line #1881:         if (packagesUri != null) 'packagesUri': packagesUri,
	OverrideMethod,          // line #1882:       });
	BlankLine,               // line #1883:
	OverrideMethod,          // line #1884:   @override
	OverrideMethod,          // line #1885:   Future<Success> removeBreakpoint(String isolateId, String breakpointId) =>
	OverrideMethod,          // line #1886:       _call('removeBreakpoint',
	OverrideMethod,          // line #1887:           {'isolateId': isolateId, 'breakpointId': breakpointId});
	BlankLine,               // line #1888:
	OverrideMethod,          // line #1889:   @override
	OverrideMethod,          // line #1890:   Future<Success> requestHeapSnapshot(String isolateId) =>
	OverrideMethod,          // line #1891:       _call('requestHeapSnapshot', {'isolateId': isolateId});
	BlankLine,               // line #1892:
	OverrideMethod,          // line #1893:   @override
	OverrideMethod,          // line #1894:   Future<Success> resume(String isolateId,
	OverrideMethod,          // line #1895:           {/*StepOption*/ String step, int frameIndex}) =>
	OverrideMethod,          // line #1896:       _call('resume', {
	OverrideMethod,          // line #1897:         'isolateId': isolateId,
	OverrideMethod,          // line #1898:         if (step != null) 'step': step,
	OverrideMethod,          // line #1899:         if (frameIndex != null) 'frameIndex': frameIndex,
	OverrideMethod,          // line #1900:       });
	BlankLine,               // line #1901:
	OverrideMethod,          // line #1902:   @override
	OverrideMethod,          // line #1903:   Future<Success> setExceptionPauseMode(
	OverrideMethod,          // line #1904:           String isolateId, /*ExceptionPauseMode*/ String mode) =>
	OverrideMethod,          // line #1905:       _call('setExceptionPauseMode', {'isolateId': isolateId, 'mode': mode});
	BlankLine,               // line #1906:
	OverrideMethod,          // line #1907:   @override
	OverrideMethod,          // line #1908:   Future<Response> setFlag(String name, String value) =>
	OverrideMethod,          // line #1909:       _call('setFlag', {'name': name, 'value': value});
	BlankLine,               // line #1910:
	OverrideMethod,          // line #1911:   @override
	OverrideMethod,          // line #1912:   Future<Success> setLibraryDebuggable(
	OverrideMethod,          // line #1913:           String isolateId, String libraryId, bool isDebuggable) =>
	OverrideMethod,          // line #1914:       _call('setLibraryDebuggable', {
	OverrideMethod,          // line #1915:         'isolateId': isolateId,
	OverrideMethod,          // line #1916:         'libraryId': libraryId,
	OverrideMethod,          // line #1917:         'isDebuggable': isDebuggable
	OverrideMethod,          // line #1918:       });
	BlankLine,               // line #1919:
	OverrideMethod,          // line #1920:   @override
	OverrideMethod,          // line #1921:   Future<Success> setName(String isolateId, String name) =>
	OverrideMethod,          // line #1922:       _call('setName', {'isolateId': isolateId, 'name': name});
	BlankLine,               // line #1923:
	OverrideMethod,          // line #1924:   @override
	OverrideMethod,          // line #1925:   Future<Success> setVMName(String name) => _call('setVMName', {'name': name});
	BlankLine,               // line #1926:
	OverrideMethod,          // line #1927:   @override
	OverrideMethod,          // line #1928:   Future<Success> setVMTimelineFlags(List<String> recordedStreams) =>
	OverrideMethod,          // line #1929:       _call('setVMTimelineFlags', {'recordedStreams': recordedStreams});
	BlankLine,               // line #1930:
	OverrideMethod,          // line #1931:   @override
	OverrideMethod,          // line #1932:   Future<Success> streamCancel(String streamId) =>
	OverrideMethod,          // line #1933:       _call('streamCancel', {'streamId': streamId});
	BlankLine,               // line #1934:
	OverrideMethod,          // line #1935:   @override
	OverrideMethod,          // line #1936:   Future<Success> streamListen(String streamId) =>
	OverrideMethod,          // line #1937:       _call('streamListen', {'streamId': streamId});
	BlankLine,               // line #1938:
	OtherMethod,             // line #1939:   /// Call an arbitrary service protocol method. This allows clients to call
	OtherMethod,             // line #1940:   /// methods not explicitly exposed by this library.
	OtherMethod,             // line #1941:   Future<Response> callMethod(String method, {String isolateId, Map args}) {
	OtherMethod,             // line #1942:     return callServiceExtension(method, isolateId: isolateId, args: args);
	OtherMethod,             // line #1943:   }
	BlankLine,               // line #1944:
	OverrideMethod,          // line #1945:   /// Invoke a specific service protocol extension method.
	OverrideMethod,          // line #1946:   ///
	OverrideMethod,          // line #1947:   /// See https://api.dart.dev/stable/dart-developer/dart-developer-library.html.
	OverrideMethod,          // line #1948:   @override
	OverrideMethod,          // line #1949:   Future<Response> callServiceExtension(String method,
	OverrideMethod,          // line #1950:       {String isolateId, Map args}) {
	OverrideMethod,          // line #1951:     if (args == null && isolateId == null) {
	OverrideMethod,          // line #1952:       return _call(method);
	OverrideMethod,          // line #1953:     } else if (args == null) {
	OverrideMethod,          // line #1954:       return _call(method, {'isolateId': isolateId});
	OverrideMethod,          // line #1955:     } else {
	OverrideMethod,          // line #1956:       args = Map.from(args);
	OverrideMethod,          // line #1957:       if (isolateId != null) {
	OverrideMethod,          // line #1958:         args['isolateId'] = isolateId;
	OverrideMethod,          // line #1959:       }
	OverrideMethod,          // line #1960:       return _call(method, args);
	OverrideMethod,          // line #1961:     }
	OverrideMethod,          // line #1962:   }
	BlankLine,               // line #1963:
	OtherMethod,             // line #1964:   Stream<String> get onSend => _onSend.stream;
	BlankLine,               // line #1965:
	OtherMethod,             // line #1966:   Stream<String> get onReceive => _onReceive.stream;
	BlankLine,               // line #1967:
	OtherMethod,             // line #1968:   void dispose() {
	OtherMethod,             // line #1969:     _streamSub.cancel();
	OtherMethod,             // line #1970:     _completers.forEach((id, c) {
	OtherMethod,             // line #1971:       final method = _methodCalls[id];
	OtherMethod,             // line #1972:       return c.completeError(RPCError(
	OtherMethod,             // line #1973:           method, RPCError.kServerError, 'Service connection disposed'));
	OtherMethod,             // line #1974:     });
	OtherMethod,             // line #1975:     _completers.clear();
	OtherMethod,             // line #1976:     if (_disposeHandler != null) {
	OtherMethod,             // line #1977:       _disposeHandler();
	OtherMethod,             // line #1978:     }
	OtherMethod,             // line #1979:     if (!_onDoneCompleter.isCompleted) {
	OtherMethod,             // line #1980:       _onDoneCompleter.complete();
	OtherMethod,             // line #1981:     }
	OtherMethod,             // line #1982:   }
	BlankLine,               // line #1983:
	OtherMethod,             // line #1984:   Future get onDone => _onDoneCompleter.future;
	BlankLine,               // line #1985:
	OtherMethod,             // line #1986:   Future<T> _call<T>(String method, [Map args = const {}]) {
	OtherMethod,             // line #1987:     String id = '${++_id}';
	OtherMethod,             // line #1988:     Completer<T> completer = Completer<T>();
	OtherMethod,             // line #1989:     _completers[id] = completer;
	OtherMethod,             // line #1990:     _methodCalls[id] = method;
	OtherMethod,             // line #1991:     Map m = {
	OtherMethod,             // line #1992:       'jsonrpc': '2.0',
	OtherMethod,             // line #1993:       'id': id,
	OtherMethod,             // line #1994:       'method': method,
	OtherMethod,             // line #1995:       'params': args,
	OtherMethod,             // line #1996:     };
	OtherMethod,             // line #1997:     String message = jsonEncode(m);
	OtherMethod,             // line #1998:     _onSend.add(message);
	OtherMethod,             // line #1999:     _writeMessage(message);
	OtherMethod,             // line #2000:     return completer.future;
	OtherMethod,             // line #2001:   }
	BlankLine,               // line #2002:
	OtherMethod,             // line #2003:   /// Register a service for invocation.
	OtherMethod,             // line #2004:   void registerServiceCallback(String service, ServiceCallback cb) {
	OtherMethod,             // line #2005:     if (_services.containsKey(service)) {
	OtherMethod,             // line #2006:       throw Exception('Service \'${service}\' already registered');
	OtherMethod,             // line #2007:     }
	OtherMethod,             // line #2008:     _services[service] = cb;
	OtherMethod,             // line #2009:   }
	BlankLine,               // line #2010:
	OtherMethod,             // line #2011:   void _processMessage(dynamic message) {
	OtherMethod,             // line #2012:     // Expect a String, an int[], or a ByteData.
	OtherMethod,             // line #2013:
	OtherMethod,             // line #2014:     if (message is String) {
	OtherMethod,             // line #2015:       _processMessageStr(message);
	OtherMethod,             // line #2016:     } else if (message is List<int>) {
	OtherMethod,             // line #2017:       Uint8List list = Uint8List.fromList(message);
	OtherMethod,             // line #2018:       _processMessageByteData(ByteData.view(list.buffer));
	OtherMethod,             // line #2019:     } else if (message is ByteData) {
	OtherMethod,             // line #2020:       _processMessageByteData(message);
	OtherMethod,             // line #2021:     } else {
	OtherMethod,             // line #2022:       _log.warning('unknown message type: ${message.runtimeType}');
	OtherMethod,             // line #2023:     }
	OtherMethod,             // line #2024:   }
	BlankLine,               // line #2025:
	OtherMethod,             // line #2026:   void _processMessageByteData(ByteData bytes) {
	OtherMethod,             // line #2027:     final int metaOffset = 4;
	OtherMethod,             // line #2028:     final int dataOffset = bytes.getUint32(0, Endian.little);
	OtherMethod,             // line #2029:     final metaLength = dataOffset - metaOffset;
	OtherMethod,             // line #2030:     final dataLength = bytes.lengthInBytes - dataOffset;
	OtherMethod,             // line #2031:     final meta = utf8.decode(Uint8List.view(
	OtherMethod,             // line #2032:         bytes.buffer, bytes.offsetInBytes + metaOffset, metaLength));
	OtherMethod,             // line #2033:     final data = ByteData.view(
	OtherMethod,             // line #2034:         bytes.buffer, bytes.offsetInBytes + dataOffset, dataLength);
	OtherMethod,             // line #2035:     dynamic map = jsonDecode(meta);
	OtherMethod,             // line #2036:     if (map != null && map['method'] == 'streamNotify') {
	OtherMethod,             // line #2037:       String streamId = map['params']['streamId'];
	OtherMethod,             // line #2038:       Map event = map['params']['event'];
	OtherMethod,             // line #2039:       event['data'] = data;
	OtherMethod,             // line #2040:       _getEventController(streamId)
	OtherMethod,             // line #2041:           .add(createServiceObject(event, const ['Event']));
	OtherMethod,             // line #2042:     }
	OtherMethod,             // line #2043:   }
	BlankLine,               // line #2044:
	OtherMethod,             // line #2045:   void _processMessageStr(String message) {
	OtherMethod,             // line #2046:     var json;
	OtherMethod,             // line #2047:     try {
	OtherMethod,             // line #2048:       _onReceive.add(message);
	OtherMethod,             // line #2049:
	OtherMethod,             // line #2050:       json = jsonDecode(message);
	OtherMethod,             // line #2051:     } catch (e, s) {
	OtherMethod,             // line #2052:       _log.severe('unable to decode message: ${message}, ${e}\n${s}');
	OtherMethod,             // line #2053:       return;
	OtherMethod,             // line #2054:     }
	OtherMethod,             // line #2055:
	OtherMethod,             // line #2056:     if (json.containsKey('method')) {
	OtherMethod,             // line #2057:       if (json.containsKey('id')) {
	OtherMethod,             // line #2058:         _processRequest(json);
	OtherMethod,             // line #2059:       } else {
	OtherMethod,             // line #2060:         _processNotification(json);
	OtherMethod,             // line #2061:       }
	OtherMethod,             // line #2062:     } else if (json.containsKey('id') &&
	OtherMethod,             // line #2063:         (json.containsKey('result') || json.containsKey('error'))) {
	OtherMethod,             // line #2064:       _processResponse(json);
	OtherMethod,             // line #2065:     } else {
	OtherMethod,             // line #2066:       _log.severe('unknown message type: ${message}');
	OtherMethod,             // line #2067:     }
	OtherMethod,             // line #2068:   }
	BlankLine,               // line #2069:
	OtherMethod,             // line #2070:   void _processResponse(Map<String, dynamic> json) {
	OtherMethod,             // line #2071:     Completer completer = _completers.remove(json['id']);
	OtherMethod,             // line #2072:     String methodName = _methodCalls.remove(json['id']);
	OtherMethod,             // line #2073:     List<String> returnTypes = _methodReturnTypes[methodName];
	OtherMethod,             // line #2074:     if (completer == null) {
	OtherMethod,             // line #2075:       _log.severe('unmatched request response: ${jsonEncode(json)}');
	OtherMethod,             // line #2076:     } else if (json['error'] != null) {
	OtherMethod,             // line #2077:       completer.completeError(RPCError.parse(methodName, json['error']));
	OtherMethod,             // line #2078:     } else {
	OtherMethod,             // line #2079:       Map<String, dynamic> result = json['result'] as Map<String, dynamic>;
	OtherMethod,             // line #2080:       String type = result['type'];
	OtherMethod,             // line #2081:       if (type == 'Sentinel') {
	OtherMethod,             // line #2082:         completer.completeError(SentinelException.parse(methodName, result));
	OtherMethod,             // line #2083:       } else if (_typeFactories[type] == null) {
	OtherMethod,             // line #2084:         completer.complete(Response.parse(result));
	OtherMethod,             // line #2085:       } else {
	OtherMethod,             // line #2086:         completer.complete(createServiceObject(result, returnTypes));
	OtherMethod,             // line #2087:       }
	OtherMethod,             // line #2088:     }
	OtherMethod,             // line #2089:   }
	BlankLine,               // line #2090:
	OtherMethod,             // line #2091:   Future _processRequest(Map<String, dynamic> json) async {
	OtherMethod,             // line #2092:     final Map m = await _routeRequest(
	OtherMethod,             // line #2093:         json['method'], json['params'] ?? <String, dynamic>{});
	OtherMethod,             // line #2094:     m['id'] = json['id'];
	OtherMethod,             // line #2095:     m['jsonrpc'] = '2.0';
	OtherMethod,             // line #2096:     String message = jsonEncode(m);
	OtherMethod,             // line #2097:     _onSend.add(message);
	OtherMethod,             // line #2098:     _writeMessage(message);
	OtherMethod,             // line #2099:   }
	BlankLine,               // line #2100:
	OtherMethod,             // line #2101:   Future _processNotification(Map<String, dynamic> json) async {
	OtherMethod,             // line #2102:     final String method = json['method'];
	OtherMethod,             // line #2103:     final Map params = json['params'] ?? <String, dynamic>{};
	OtherMethod,             // line #2104:     if (method == 'streamNotify') {
	OtherMethod,             // line #2105:       String streamId = params['streamId'];
	OtherMethod,             // line #2106:       _getEventController(streamId)
	OtherMethod,             // line #2107:           .add(createServiceObject(params['event'], const ['Event']));
	OtherMethod,             // line #2108:     } else {
	OtherMethod,             // line #2109:       await _routeRequest(method, params);
	OtherMethod,             // line #2110:     }
	OtherMethod,             // line #2111:   }
	BlankLine,               // line #2112:
	OtherMethod,             // line #2113:   Future<Map> _routeRequest(String method, Map<String, dynamic> params) async {
	OtherMethod,             // line #2114:     if (!_services.containsKey(method)) {
	OtherMethod,             // line #2115:       RPCError error = RPCError(
	OtherMethod,             // line #2116:           method, RPCError.kMethodNotFound, 'method not found \'$method\'');
	OtherMethod,             // line #2117:       return {'error': error.toMap()};
	OtherMethod,             // line #2118:     }
	OtherMethod,             // line #2119:
	OtherMethod,             // line #2120:     try {
	OtherMethod,             // line #2121:       return await _services[method](params);
	OtherMethod,             // line #2122:     } catch (e, st) {
	OtherMethod,             // line #2123:       RPCError error = RPCError.withDetails(
	OtherMethod,             // line #2124:         method,
	OtherMethod,             // line #2125:         RPCError.kServerError,
	OtherMethod,             // line #2126:         '$e',
	OtherMethod,             // line #2127:         details: '$st',
	OtherMethod,             // line #2128:       );
	OtherMethod,             // line #2129:       return {'error': error.toMap()};
	OtherMethod,             // line #2130:     }
	OtherMethod,             // line #2131:   }
	BlankLine,               // line #2132:
}

var wantVMServiceClass05 = []EntityType{
	Unknown,          // line #2136: {
	StaticVariable,   // line #2137:   /// Application specific error codes.
	StaticVariable,   // line #2138:   static const int kServerError = -32000;
	BlankLine,        // line #2139:
	StaticVariable,   // line #2140:   /// The JSON sent is not a valid Request object.
	StaticVariable,   // line #2141:   static const int kInvalidRequest = -32600;
	BlankLine,        // line #2142:
	StaticVariable,   // line #2143:   /// The method does not exist or is not available.
	StaticVariable,   // line #2144:   static const int kMethodNotFound = -32601;
	BlankLine,        // line #2145:
	StaticVariable,   // line #2146:   /// Invalid method parameter(s), such as a mismatched type.
	StaticVariable,   // line #2147:   static const int kInvalidParams = -32602;
	BlankLine,        // line #2148:
	StaticVariable,   // line #2149:   /// Internal JSON-RPC error.
	StaticVariable,   // line #2150:   static const int kInternalError = -32603;
	BlankLine,        // line #2151:
	OtherMethod,      // line #2152:   static RPCError parse(String callingMethod, dynamic json) {
	OtherMethod,      // line #2153:     return RPCError(callingMethod, json['code'], json['message'], json['data']);
	OtherMethod,      // line #2154:   }
	BlankLine,        // line #2155:
	InstanceVariable, // line #2156:   final String callingMethod;
	InstanceVariable, // line #2157:   final int code;
	InstanceVariable, // line #2158:   final String message;
	InstanceVariable, // line #2159:   final Map data;
	BlankLine,        // line #2160:
	MainConstructor,  // line #2161:   RPCError(this.callingMethod, this.code, this.message, [this.data]);
	BlankLine,        // line #2162:
	NamedConstructor, // line #2163:   RPCError.withDetails(this.callingMethod, this.code, this.message,
	NamedConstructor, // line #2164:       {Object details})
	NamedConstructor, // line #2165:       : data = details == null ? null : <String, dynamic>{} {
	NamedConstructor, // line #2166:     if (details != null) {
	NamedConstructor, // line #2167:       data['details'] = details;
	NamedConstructor, // line #2168:     }
	NamedConstructor, // line #2169:   }
	BlankLine,        // line #2170:
	OtherMethod,      // line #2171:   String get details => data == null ? null : data['details'];
	BlankLine,        // line #2172:
	OtherMethod,      // line #2173:   /// Return a map representation of this error suitable for converstion to
	OtherMethod,      // line #2174:   /// json.
	OtherMethod,      // line #2175:   Map<String, dynamic> toMap() {
	OtherMethod,      // line #2176:     Map<String, dynamic> map = {
	OtherMethod,      // line #2177:       'code': code,
	OtherMethod,      // line #2178:       'message': message,
	OtherMethod,      // line #2179:     };
	OtherMethod,      // line #2180:     if (data != null) {
	OtherMethod,      // line #2181:       map['data'] = data;
	OtherMethod,      // line #2182:     }
	OtherMethod,      // line #2183:     return map;
	OtherMethod,      // line #2184:   }
	BlankLine,        // line #2185:
	OtherMethod,      // line #2186:   String toString() {
	OtherMethod,      // line #2187:     if (details == null) {
	OtherMethod,      // line #2188:       return '$callingMethod: ($code) $message';
	OtherMethod,      // line #2189:     } else {
	OtherMethod,      // line #2190:       return '$callingMethod: ($code) $message\n$details';
	OtherMethod,      // line #2191:     }
	OtherMethod,      // line #2192:   }
	BlankLine,        // line #2193:
}

var wantVMServiceClass06 = []EntityType{
	Unknown,          // line #2196: {
	InstanceVariable, // line #2197:   final String callingMethod;
	InstanceVariable, // line #2198:   final Sentinel sentinel;
	BlankLine,        // line #2199:
	NamedConstructor, // line #2200:   SentinelException.parse(this.callingMethod, Map<String, dynamic> data)
	NamedConstructor, // line #2201:       : sentinel = Sentinel.parse(data);
	BlankLine,        // line #2202:
	OtherMethod,      // line #2203:   String toString() => '$sentinel from ${callingMethod}()';
	BlankLine,        // line #2204:
}

var wantVMServiceClass07 = []EntityType{
	Unknown,          // line #2207: {
	OtherMethod,      // line #2208:   static ExtensionData parse(Map json) =>
	OtherMethod,      // line #2209:       json == null ? null : ExtensionData._fromJson(json);
	BlankLine,        // line #2210:
	InstanceVariable, // line #2211:   final Map data;
	BlankLine,        // line #2212:
	MainConstructor,  // line #2213:   ExtensionData() : data = {};
	BlankLine,        // line #2214:
	NamedConstructor, // line #2215:   ExtensionData._fromJson(this.data);
	BlankLine,        // line #2216:
	OtherMethod,      // line #2217:   String toString() => '[ExtensionData ${data}]';
	BlankLine,        // line #2218:
}

var wantVMServiceClass08 = []EntityType{
	Unknown,     // line #2222: {
	OtherMethod, // line #2223:   /// Log a warning level message.
	OtherMethod, // line #2224:   void warning(String message);
	BlankLine,   // line #2225:
	OtherMethod, // line #2226:   /// Log an error level message.
	OtherMethod, // line #2227:   void severe(String message);
	BlankLine,   // line #2228:
}

var wantVMServiceClass09 = []EntityType{
	Unknown,     // line #2230: {
	OtherMethod, // line #2231:   void warning(String message) {}
	OtherMethod, // line #2232:   void severe(String message) {}
	BlankLine,   // line #2233:
}

var wantVMServiceClass10 = []EntityType{
	Unknown,          // line #2236: {
	NamedConstructor, // line #2237:   CodeKind._();
	BlankLine,        // line #2238:
	StaticVariable,   // line #2239:   static const String kDart = 'Dart';
	StaticVariable,   // line #2240:   static const String kNative = 'Native';
	StaticVariable,   // line #2241:   static const String kStub = 'Stub';
	StaticVariable,   // line #2242:   static const String kTag = 'Tag';
	StaticVariable,   // line #2243:   static const String kCollected = 'Collected';
	BlankLine,        // line #2244:
}

var wantVMServiceClass11 = []EntityType{
	Unknown,          // line #2246: {
	NamedConstructor, // line #2247:   ErrorKind._();
	BlankLine,        // line #2248:
	StaticVariable,   // line #2249:   /// The isolate has encountered an unhandled Dart exception.
	StaticVariable,   // line #2250:   static const String kUnhandledException = 'UnhandledException';
	BlankLine,        // line #2251:
	StaticVariable,   // line #2252:   /// The isolate has encountered a Dart language error in the program.
	StaticVariable,   // line #2253:   static const String kLanguageError = 'LanguageError';
	BlankLine,        // line #2254:
	StaticVariable,   // line #2255:   /// The isolate has encountered an internal error. These errors should be
	StaticVariable,   // line #2256:   /// reported as bugs.
	StaticVariable,   // line #2257:   static const String kInternalError = 'InternalError';
	BlankLine,        // line #2258:
	StaticVariable,   // line #2259:   /// The isolate has been terminated by an external source.
	StaticVariable,   // line #2260:   static const String kTerminationError = 'TerminationError';
	BlankLine,        // line #2261:
}

var wantVMServiceClass12 = []EntityType{
	Unknown,          // line #2264: {
	NamedConstructor, // line #2265:   EventStreams._();
	BlankLine,        // line #2266:
	StaticVariable,   // line #2267:   static const String kVM = 'VM';
	StaticVariable,   // line #2268:   static const String kIsolate = 'Isolate';
	StaticVariable,   // line #2269:   static const String kDebug = 'Debug';
	StaticVariable,   // line #2270:   static const String kGC = 'GC';
	StaticVariable,   // line #2271:   static const String kExtension = 'Extension';
	StaticVariable,   // line #2272:   static const String kTimeline = 'Timeline';
	StaticVariable,   // line #2273:   static const String kLogging = 'Logging';
	StaticVariable,   // line #2274:   static const String kService = 'Service';
	StaticVariable,   // line #2275:   static const String kHeapSnapshot = 'HeapSnapshot';
	StaticVariable,   // line #2276:   static const String kStdout = 'Stdout';
	StaticVariable,   // line #2277:   static const String kStderr = 'Stderr';
	BlankLine,        // line #2278:
}

var wantVMServiceClass13 = []EntityType{
	Unknown,          // line #2282: {
	NamedConstructor, // line #2283:   EventKind._();
	BlankLine,        // line #2284:
	StaticVariable,   // line #2285:   /// Notification that VM identifying information has changed. Currently used
	StaticVariable,   // line #2286:   /// to notify of changes to the VM debugging name via setVMName.
	StaticVariable,   // line #2287:   static const String kVMUpdate = 'VMUpdate';
	BlankLine,        // line #2288:
	StaticVariable,   // line #2289:   /// Notification that a VM flag has been changed via the service protocol.
	StaticVariable,   // line #2290:   static const String kVMFlagUpdate = 'VMFlagUpdate';
	BlankLine,        // line #2291:
	StaticVariable,   // line #2292:   /// Notification that a new isolate has started.
	StaticVariable,   // line #2293:   static const String kIsolateStart = 'IsolateStart';
	BlankLine,        // line #2294:
	StaticVariable,   // line #2295:   /// Notification that an isolate is ready to run.
	StaticVariable,   // line #2296:   static const String kIsolateRunnable = 'IsolateRunnable';
	BlankLine,        // line #2297:
	StaticVariable,   // line #2298:   /// Notification that an isolate has exited.
	StaticVariable,   // line #2299:   static const String kIsolateExit = 'IsolateExit';
	BlankLine,        // line #2300:
	StaticVariable,   // line #2301:   /// Notification that isolate identifying information has changed. Currently
	StaticVariable,   // line #2302:   /// used to notify of changes to the isolate debugging name via setName.
	StaticVariable,   // line #2303:   static const String kIsolateUpdate = 'IsolateUpdate';
	BlankLine,        // line #2304:
	StaticVariable,   // line #2305:   /// Notification that an isolate has been reloaded.
	StaticVariable,   // line #2306:   static const String kIsolateReload = 'IsolateReload';
	BlankLine,        // line #2307:
	StaticVariable,   // line #2308:   /// Notification that an extension RPC was registered on an isolate.
	StaticVariable,   // line #2309:   static const String kServiceExtensionAdded = 'ServiceExtensionAdded';
	BlankLine,        // line #2310:
	StaticVariable,   // line #2311:   /// An isolate has paused at start, before executing code.
	StaticVariable,   // line #2312:   static const String kPauseStart = 'PauseStart';
	BlankLine,        // line #2313:
	StaticVariable,   // line #2314:   /// An isolate has paused at exit, before terminating.
	StaticVariable,   // line #2315:   static const String kPauseExit = 'PauseExit';
	BlankLine,        // line #2316:
	StaticVariable,   // line #2317:   /// An isolate has paused at a breakpoint or due to stepping.
	StaticVariable,   // line #2318:   static const String kPauseBreakpoint = 'PauseBreakpoint';
	BlankLine,        // line #2319:
	StaticVariable,   // line #2320:   /// An isolate has paused due to interruption via pause.
	StaticVariable,   // line #2321:   static const String kPauseInterrupted = 'PauseInterrupted';
	BlankLine,        // line #2322:
	StaticVariable,   // line #2323:   /// An isolate has paused due to an exception.
	StaticVariable,   // line #2324:   static const String kPauseException = 'PauseException';
	BlankLine,        // line #2325:
	StaticVariable,   // line #2326:   /// An isolate has paused after a service request.
	StaticVariable,   // line #2327:   static const String kPausePostRequest = 'PausePostRequest';
	BlankLine,        // line #2328:
	StaticVariable,   // line #2329:   /// An isolate has started or resumed execution.
	StaticVariable,   // line #2330:   static const String kResume = 'Resume';
	BlankLine,        // line #2331:
	StaticVariable,   // line #2332:   /// Indicates an isolate is not yet runnable. Only appears in an Isolate's
	StaticVariable,   // line #2333:   /// pauseEvent. Never sent over a stream.
	StaticVariable,   // line #2334:   static const String kNone = 'None';
	BlankLine,        // line #2335:
	StaticVariable,   // line #2336:   /// A breakpoint has been added for an isolate.
	StaticVariable,   // line #2337:   static const String kBreakpointAdded = 'BreakpointAdded';
	BlankLine,        // line #2338:
	StaticVariable,   // line #2339:   /// An unresolved breakpoint has been resolved for an isolate.
	StaticVariable,   // line #2340:   static const String kBreakpointResolved = 'BreakpointResolved';
	BlankLine,        // line #2341:
	StaticVariable,   // line #2342:   /// A breakpoint has been removed.
	StaticVariable,   // line #2343:   static const String kBreakpointRemoved = 'BreakpointRemoved';
	BlankLine,        // line #2344:
	StaticVariable,   // line #2345:   /// A garbage collection event.
	StaticVariable,   // line #2346:   static const String kGC = 'GC';
	BlankLine,        // line #2347:
	StaticVariable,   // line #2348:   /// Notification of bytes written, for example, to stdout/stderr.
	StaticVariable,   // line #2349:   static const String kWriteEvent = 'WriteEvent';
	BlankLine,        // line #2350:
	StaticVariable,   // line #2351:   /// Notification from dart:developer.inspect.
	StaticVariable,   // line #2352:   static const String kInspect = 'Inspect';
	BlankLine,        // line #2353:
	StaticVariable,   // line #2354:   /// Event from dart:developer.postEvent.
	StaticVariable,   // line #2355:   static const String kExtension = 'Extension';
	BlankLine,        // line #2356:
	StaticVariable,   // line #2357:   /// Event from dart:developer.log.
	StaticVariable,   // line #2358:   static const String kLogging = 'Logging';
	BlankLine,        // line #2359:
	StaticVariable,   // line #2360:   /// A block of timeline events has been completed.
	StaticVariable,   // line #2361:   ///
	StaticVariable,   // line #2362:   /// This service event is not sent for individual timeline events. It is
	StaticVariable,   // line #2363:   /// subject to buffering, so the most recent timeline events may never be
	StaticVariable,   // line #2364:   /// included in any TimelineEvents event if no timeline events occur later to
	StaticVariable,   // line #2365:   /// complete the block.
	StaticVariable,   // line #2366:   static const String kTimelineEvents = 'TimelineEvents';
	BlankLine,        // line #2367:
	StaticVariable,   // line #2368:   /// The set of active timeline streams was changed via `setVMTimelineFlags`.
	StaticVariable,   // line #2369:   static const String kTimelineStreamSubscriptionsUpdate =
	StaticVariable,   // line #2370:       'TimelineStreamSubscriptionsUpdate';
	BlankLine,        // line #2371:
	StaticVariable,   // line #2372:   /// Notification that a Service has been registered into the Service Protocol
	StaticVariable,   // line #2373:   /// from another client.
	StaticVariable,   // line #2374:   static const String kServiceRegistered = 'ServiceRegistered';
	BlankLine,        // line #2375:
	StaticVariable,   // line #2376:   /// Notification that a Service has been removed from the Service Protocol
	StaticVariable,   // line #2377:   /// from another client.
	StaticVariable,   // line #2378:   static const String kServiceUnregistered = 'ServiceUnregistered';
	BlankLine,        // line #2379:
}

var wantVMServiceClass14 = []EntityType{
	Unknown,          // line #2383: {
	NamedConstructor, // line #2384:   InstanceKind._();
	BlankLine,        // line #2385:
	StaticVariable,   // line #2386:   /// A general instance of the Dart class Object.
	StaticVariable,   // line #2387:   static const String kPlainInstance = 'PlainInstance';
	BlankLine,        // line #2388:
	StaticVariable,   // line #2389:   /// null instance.
	StaticVariable,   // line #2390:   static const String kNull = 'Null';
	BlankLine,        // line #2391:
	StaticVariable,   // line #2392:   /// true or false.
	StaticVariable,   // line #2393:   static const String kBool = 'Bool';
	BlankLine,        // line #2394:
	StaticVariable,   // line #2395:   /// An instance of the Dart class double.
	StaticVariable,   // line #2396:   static const String kDouble = 'Double';
	BlankLine,        // line #2397:
	StaticVariable,   // line #2398:   /// An instance of the Dart class int.
	StaticVariable,   // line #2399:   static const String kInt = 'Int';
	BlankLine,        // line #2400:
	StaticVariable,   // line #2401:   /// An instance of the Dart class String.
	StaticVariable,   // line #2402:   static const String kString = 'String';
	BlankLine,        // line #2403:
	StaticVariable,   // line #2404:   /// An instance of the built-in VM List implementation. User-defined Lists
	StaticVariable,   // line #2405:   /// will be PlainInstance.
	StaticVariable,   // line #2406:   static const String kList = 'List';
	BlankLine,        // line #2407:
	StaticVariable,   // line #2408:   /// An instance of the built-in VM Map implementation. User-defined Maps will
	StaticVariable,   // line #2409:   /// be PlainInstance.
	StaticVariable,   // line #2410:   static const String kMap = 'Map';
	BlankLine,        // line #2411:
	StaticVariable,   // line #2412:   /// Vector instance kinds.
	StaticVariable,   // line #2413:   static const String kFloat32x4 = 'Float32x4';
	StaticVariable,   // line #2414:   static const String kFloat64x2 = 'Float64x2';
	StaticVariable,   // line #2415:   static const String kInt32x4 = 'Int32x4';
	BlankLine,        // line #2416:
	StaticVariable,   // line #2417:   /// An instance of the built-in VM TypedData implementations. User-defined
	StaticVariable,   // line #2418:   /// TypedDatas will be PlainInstance.
	StaticVariable,   // line #2419:   static const String kUint8ClampedList = 'Uint8ClampedList';
	StaticVariable,   // line #2420:   static const String kUint8List = 'Uint8List';
	StaticVariable,   // line #2421:   static const String kUint16List = 'Uint16List';
	StaticVariable,   // line #2422:   static const String kUint32List = 'Uint32List';
	StaticVariable,   // line #2423:   static const String kUint64List = 'Uint64List';
	StaticVariable,   // line #2424:   static const String kInt8List = 'Int8List';
	StaticVariable,   // line #2425:   static const String kInt16List = 'Int16List';
	StaticVariable,   // line #2426:   static const String kInt32List = 'Int32List';
	StaticVariable,   // line #2427:   static const String kInt64List = 'Int64List';
	StaticVariable,   // line #2428:   static const String kFloat32List = 'Float32List';
	StaticVariable,   // line #2429:   static const String kFloat64List = 'Float64List';
	StaticVariable,   // line #2430:   static const String kInt32x4List = 'Int32x4List';
	StaticVariable,   // line #2431:   static const String kFloat32x4List = 'Float32x4List';
	StaticVariable,   // line #2432:   static const String kFloat64x2List = 'Float64x2List';
	BlankLine,        // line #2433:
	StaticVariable,   // line #2434:   /// An instance of the Dart class StackTrace.
	StaticVariable,   // line #2435:   static const String kStackTrace = 'StackTrace';
	BlankLine,        // line #2436:
	StaticVariable,   // line #2437:   /// An instance of the built-in VM Closure implementation. User-defined
	StaticVariable,   // line #2438:   /// Closures will be PlainInstance.
	StaticVariable,   // line #2439:   static const String kClosure = 'Closure';
	BlankLine,        // line #2440:
	StaticVariable,   // line #2441:   /// An instance of the Dart class MirrorReference.
	StaticVariable,   // line #2442:   static const String kMirrorReference = 'MirrorReference';
	BlankLine,        // line #2443:
	StaticVariable,   // line #2444:   /// An instance of the Dart class RegExp.
	StaticVariable,   // line #2445:   static const String kRegExp = 'RegExp';
	BlankLine,        // line #2446:
	StaticVariable,   // line #2447:   /// An instance of the Dart class WeakProperty.
	StaticVariable,   // line #2448:   static const String kWeakProperty = 'WeakProperty';
	BlankLine,        // line #2449:
	StaticVariable,   // line #2450:   /// An instance of the Dart class Type.
	StaticVariable,   // line #2451:   static const String kType = 'Type';
	BlankLine,        // line #2452:
	StaticVariable,   // line #2453:   /// An instance of the Dart class TypeParameter.
	StaticVariable,   // line #2454:   static const String kTypeParameter = 'TypeParameter';
	BlankLine,        // line #2455:
	StaticVariable,   // line #2456:   /// An instance of the Dart class TypeRef.
	StaticVariable,   // line #2457:   static const String kTypeRef = 'TypeRef';
	BlankLine,        // line #2458:
	StaticVariable,   // line #2459:   /// An instance of the Dart class BoundedType.
	StaticVariable,   // line #2460:   static const String kBoundedType = 'BoundedType';
	BlankLine,        // line #2461:
	StaticVariable,   // line #2462:   /// An instance of the Dart class ReceivePort.
	StaticVariable,   // line #2463:   static const String kReceivePort = 'ReceivePort';
	BlankLine,        // line #2464:
}

var wantVMServiceClass15 = []EntityType{
	Unknown,          // line #2471: {
	NamedConstructor, // line #2472:   SentinelKind._();
	BlankLine,        // line #2473:
	StaticVariable,   // line #2474:   /// Indicates that the object referred to has been collected by the GC.
	StaticVariable,   // line #2475:   static const String kCollected = 'Collected';
	BlankLine,        // line #2476:
	StaticVariable,   // line #2477:   /// Indicates that an object id has expired.
	StaticVariable,   // line #2478:   static const String kExpired = 'Expired';
	BlankLine,        // line #2479:
	StaticVariable,   // line #2480:   /// Indicates that a variable or field has not been initialized.
	StaticVariable,   // line #2481:   static const String kNotInitialized = 'NotInitialized';
	BlankLine,        // line #2482:
	StaticVariable,   // line #2483:   /// Indicates that a variable or field is in the process of being initialized.
	StaticVariable,   // line #2484:   static const String kBeingInitialized = 'BeingInitialized';
	BlankLine,        // line #2485:
	StaticVariable,   // line #2486:   /// Indicates that a variable has been eliminated by the optimizing compiler.
	StaticVariable,   // line #2487:   static const String kOptimizedOut = 'OptimizedOut';
	BlankLine,        // line #2488:
	StaticVariable,   // line #2489:   /// Reserved for future use.
	StaticVariable,   // line #2490:   static const String kFree = 'Free';
	BlankLine,        // line #2491:
}

var wantVMServiceClass16 = []EntityType{
	Unknown,          // line #2494: {
	NamedConstructor, // line #2495:   FrameKind._();
	BlankLine,        // line #2496:
	StaticVariable,   // line #2497:   static const String kRegular = 'Regular';
	StaticVariable,   // line #2498:   static const String kAsyncCausal = 'AsyncCausal';
	StaticVariable,   // line #2499:   static const String kAsyncSuspensionMarker = 'AsyncSuspensionMarker';
	StaticVariable,   // line #2500:   static const String kAsyncActivation = 'AsyncActivation';
	BlankLine,        // line #2501:
}

var wantVMServiceClass17 = []EntityType{
	Unknown,          // line #2503: {
	NamedConstructor, // line #2504:   SourceReportKind._();
	BlankLine,        // line #2505:
	StaticVariable,   // line #2506:   /// Used to request a code coverage information.
	StaticVariable,   // line #2507:   static const String kCoverage = 'Coverage';
	BlankLine,        // line #2508:
	StaticVariable,   // line #2509:   /// Used to request a list of token positions of possible breakpoints.
	StaticVariable,   // line #2510:   static const String kPossibleBreakpoints = 'PossibleBreakpoints';
	BlankLine,        // line #2511:
}

var wantVMServiceClass18 = []EntityType{
	Unknown,          // line #2515: {
	NamedConstructor, // line #2516:   ExceptionPauseMode._();
	BlankLine,        // line #2517:
	StaticVariable,   // line #2518:   static const String kNone = 'None';
	StaticVariable,   // line #2519:   static const String kUnhandled = 'Unhandled';
	StaticVariable,   // line #2520:   static const String kAll = 'All';
	BlankLine,        // line #2521:
}

var wantVMServiceClass19 = []EntityType{
	Unknown,          // line #2525: {
	NamedConstructor, // line #2526:   StepOption._();
	BlankLine,        // line #2527:
	StaticVariable,   // line #2528:   static const String kInto = 'Into';
	StaticVariable,   // line #2529:   static const String kOver = 'Over';
	StaticVariable,   // line #2530:   static const String kOverAsyncSuspension = 'OverAsyncSuspension';
	StaticVariable,   // line #2531:   static const String kOut = 'Out';
	StaticVariable,   // line #2532:   static const String kRewind = 'Rewind';
	BlankLine,        // line #2533:
}

var wantVMServiceClass20 = []EntityType{
	Unknown,          // line #2537: {
	OtherMethod,      // line #2538:   static AllocationProfile parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #2539:       json == null ? null : AllocationProfile._fromJson(json);
	BlankLine,        // line #2540:
	InstanceVariable, // line #2541:   /// Allocation information for all class types.
	InstanceVariable, // line #2542:   List<ClassHeapStats> members;
	BlankLine,        // line #2543:
	InstanceVariable, // line #2544:   /// Information about memory usage for the isolate.
	InstanceVariable, // line #2545:   MemoryUsage memoryUsage;
	BlankLine,        // line #2546:
	InstanceVariable, // line #2547:   /// The timestamp of the last accumulator reset.
	InstanceVariable, // line #2548:   ///
	InstanceVariable, // line #2549:   /// If the accumulators have not been reset, this field is not present.
	InstanceVariable, // line #2550:   @optional
	InstanceVariable, // line #2551:   int dateLastAccumulatorReset;
	BlankLine,        // line #2552:
	InstanceVariable, // line #2553:   /// The timestamp of the last manually triggered GC.
	InstanceVariable, // line #2554:   ///
	InstanceVariable, // line #2555:   /// If a GC has not been triggered manually, this field is not present.
	InstanceVariable, // line #2556:   @optional
	InstanceVariable, // line #2557:   int dateLastServiceGC;
	BlankLine,        // line #2558:
	MainConstructor,  // line #2559:   AllocationProfile({
	MainConstructor,  // line #2560:     @required this.members,
	MainConstructor,  // line #2561:     @required this.memoryUsage,
	MainConstructor,  // line #2562:     this.dateLastAccumulatorReset,
	MainConstructor,  // line #2563:     this.dateLastServiceGC,
	MainConstructor,  // line #2564:   });
	BlankLine,        // line #2565:
	NamedConstructor, // line #2566:   AllocationProfile._fromJson(Map<String, dynamic> json)
	NamedConstructor, // line #2567:       : super._fromJson(json) {
	NamedConstructor, // line #2568:     members = List<ClassHeapStats>.from(
	NamedConstructor, // line #2569:         createServiceObject(json['members'], const ['ClassHeapStats']) ?? []);
	NamedConstructor, // line #2570:     memoryUsage =
	NamedConstructor, // line #2571:         createServiceObject(json['memoryUsage'], const ['MemoryUsage']);
	NamedConstructor, // line #2572:     dateLastAccumulatorReset = json['dateLastAccumulatorReset'] is String
	NamedConstructor, // line #2573:         ? int.parse(json['dateLastAccumulatorReset'])
	NamedConstructor, // line #2574:         : json['dateLastAccumulatorReset'];
	NamedConstructor, // line #2575:     dateLastServiceGC = json['dateLastServiceGC'] is String
	NamedConstructor, // line #2576:         ? int.parse(json['dateLastServiceGC'])
	NamedConstructor, // line #2577:         : json['dateLastServiceGC'];
	NamedConstructor, // line #2578:   }
	BlankLine,        // line #2579:
	OverrideMethod,   // line #2580:   @override
	OverrideMethod,   // line #2581:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #2582:     var json = <String, dynamic>{};
	OverrideMethod,   // line #2583:     json['type'] = 'AllocationProfile';
	OverrideMethod,   // line #2584:     json.addAll({
	OverrideMethod,   // line #2585:       'members': members.map((f) => f.toJson()).toList(),
	OverrideMethod,   // line #2586:       'memoryUsage': memoryUsage.toJson(),
	OverrideMethod,   // line #2587:     });
	OverrideMethod,   // line #2588:     _setIfNotNull(json, 'dateLastAccumulatorReset', dateLastAccumulatorReset);
	OverrideMethod,   // line #2589:     _setIfNotNull(json, 'dateLastServiceGC', dateLastServiceGC);
	OverrideMethod,   // line #2590:     return json;
	OverrideMethod,   // line #2591:   }
	BlankLine,        // line #2592:
	OtherMethod,      // line #2593:   String toString() => '[AllocationProfile ' //
	OtherMethod,      // line #2594:       'type: ${type}, members: ${members}, memoryUsage: ${memoryUsage}]';
	BlankLine,        // line #2595:
}

var wantVMServiceClass21 = []EntityType{
	Unknown,          // line #2605: {
	OtherMethod,      // line #2606:   static BoundField parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #2607:       json == null ? null : BoundField._fromJson(json);
	BlankLine,        // line #2608:
	InstanceVariable, // line #2609:   FieldRef decl;
	BlankLine,        // line #2610:
	InstanceVariable, // line #2611:   /// [value] can be one of [InstanceRef] or [Sentinel].
	InstanceVariable, // line #2612:   dynamic value;
	BlankLine,        // line #2613:
	MainConstructor,  // line #2614:   BoundField({
	MainConstructor,  // line #2615:     @required this.decl,
	MainConstructor,  // line #2616:     @required this.value,
	MainConstructor,  // line #2617:   });
	BlankLine,        // line #2618:
	NamedConstructor, // line #2619:   BoundField._fromJson(Map<String, dynamic> json) {
	NamedConstructor, // line #2620:     decl = createServiceObject(json['decl'], const ['FieldRef']);
	NamedConstructor, // line #2621:     value =
	NamedConstructor, // line #2622:         createServiceObject(json['value'], const ['InstanceRef', 'Sentinel']);
	NamedConstructor, // line #2623:   }
	BlankLine,        // line #2624:
	OtherMethod,      // line #2625:   Map<String, dynamic> toJson() {
	OtherMethod,      // line #2626:     var json = <String, dynamic>{};
	OtherMethod,      // line #2627:     json.addAll({
	OtherMethod,      // line #2628:       'decl': decl.toJson(),
	OtherMethod,      // line #2629:       'value': value.toJson(),
	OtherMethod,      // line #2630:     });
	OtherMethod,      // line #2631:     return json;
	OtherMethod,      // line #2632:   }
	BlankLine,        // line #2633:
	OtherMethod,      // line #2634:   String toString() => '[BoundField decl: ${decl}, value: ${value}]';
	BlankLine,        // line #2635:
}

var wantVMServiceClass22 = []EntityType{
	Unknown,          // line #2648: {
	OtherMethod,      // line #2649:   static BoundVariable parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #2650:       json == null ? null : BoundVariable._fromJson(json);
	BlankLine,        // line #2651:
	InstanceVariable, // line #2652:   String name;
	BlankLine,        // line #2653:
	InstanceVariable, // line #2654:   /// [value] can be one of [InstanceRef], [TypeArgumentsRef] or [Sentinel].
	InstanceVariable, // line #2655:   dynamic value;
	BlankLine,        // line #2656:
	InstanceVariable, // line #2657:   /// The token position where this variable was declared.
	InstanceVariable, // line #2658:   int declarationTokenPos;
	BlankLine,        // line #2659:
	InstanceVariable, // line #2660:   /// The first token position where this variable is visible to the scope.
	InstanceVariable, // line #2661:   int scopeStartTokenPos;
	BlankLine,        // line #2662:
	InstanceVariable, // line #2663:   /// The last token position where this variable is visible to the scope.
	InstanceVariable, // line #2664:   int scopeEndTokenPos;
	BlankLine,        // line #2665:
	MainConstructor,  // line #2666:   BoundVariable({
	MainConstructor,  // line #2667:     @required this.name,
	MainConstructor,  // line #2668:     @required this.value,
	MainConstructor,  // line #2669:     @required this.declarationTokenPos,
	MainConstructor,  // line #2670:     @required this.scopeStartTokenPos,
	MainConstructor,  // line #2671:     @required this.scopeEndTokenPos,
	MainConstructor,  // line #2672:   });
	BlankLine,        // line #2673:
	NamedConstructor, // line #2674:   BoundVariable._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #2675:     name = json['name'];
	NamedConstructor, // line #2676:     value = createServiceObject(
	NamedConstructor, // line #2677:         json['value'], const ['InstanceRef', 'TypeArgumentsRef', 'Sentinel']);
	NamedConstructor, // line #2678:     declarationTokenPos = json['declarationTokenPos'];
	NamedConstructor, // line #2679:     scopeStartTokenPos = json['scopeStartTokenPos'];
	NamedConstructor, // line #2680:     scopeEndTokenPos = json['scopeEndTokenPos'];
	NamedConstructor, // line #2681:   }
	BlankLine,        // line #2682:
	OverrideMethod,   // line #2683:   @override
	OverrideMethod,   // line #2684:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #2685:     var json = <String, dynamic>{};
	OverrideMethod,   // line #2686:     json['type'] = 'BoundVariable';
	OverrideMethod,   // line #2687:     json.addAll({
	OverrideMethod,   // line #2688:       'name': name,
	OverrideMethod,   // line #2689:       'value': value.toJson(),
	OverrideMethod,   // line #2690:       'declarationTokenPos': declarationTokenPos,
	OverrideMethod,   // line #2691:       'scopeStartTokenPos': scopeStartTokenPos,
	OverrideMethod,   // line #2692:       'scopeEndTokenPos': scopeEndTokenPos,
	OverrideMethod,   // line #2693:     });
	OverrideMethod,   // line #2694:     return json;
	OverrideMethod,   // line #2695:   }
	BlankLine,        // line #2696:
	OtherMethod,      // line #2697:   String toString() => '[BoundVariable ' //
	OtherMethod,      // line #2698:       'type: ${type}, name: ${name}, value: ${value}, declarationTokenPos: ${declarationTokenPos}, ' //
	OtherMethod,      // line #2699:       'scopeStartTokenPos: ${scopeStartTokenPos}, scopeEndTokenPos: ${scopeEndTokenPos}]';
	BlankLine,        // line #2700:
}

var wantVMServiceClass23 = []EntityType{
	Unknown,          // line #2708: {
	OtherMethod,      // line #2709:   static Breakpoint parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #2710:       json == null ? null : Breakpoint._fromJson(json);
	BlankLine,        // line #2711:
	InstanceVariable, // line #2712:   /// A number identifying this breakpoint to the user.
	InstanceVariable, // line #2713:   int breakpointNumber;
	BlankLine,        // line #2714:
	InstanceVariable, // line #2715:   /// Has this breakpoint been assigned to a specific program location?
	InstanceVariable, // line #2716:   bool resolved;
	BlankLine,        // line #2717:
	InstanceVariable, // line #2718:   /// Is this a breakpoint that was added synthetically as part of a step
	InstanceVariable, // line #2719:   /// OverAsyncSuspension resume command?
	InstanceVariable, // line #2720:   @optional
	InstanceVariable, // line #2721:   bool isSyntheticAsyncContinuation;
	BlankLine,        // line #2722:
	InstanceVariable, // line #2723:   /// SourceLocation when breakpoint is resolved, UnresolvedSourceLocation when
	InstanceVariable, // line #2724:   /// a breakpoint is not resolved.
	InstanceVariable, // line #2725:   ///
	InstanceVariable, // line #2726:   /// [location] can be one of [SourceLocation] or [UnresolvedSourceLocation].
	InstanceVariable, // line #2727:   dynamic location;
	BlankLine,        // line #2728:
	MainConstructor,  // line #2729:   Breakpoint({
	MainConstructor,  // line #2730:     @required this.breakpointNumber,
	MainConstructor,  // line #2731:     @required this.resolved,
	MainConstructor,  // line #2732:     @required this.location,
	MainConstructor,  // line #2733:     @required String id,
	MainConstructor,  // line #2734:     this.isSyntheticAsyncContinuation,
	MainConstructor,  // line #2735:   }) : super(id: id);
	BlankLine,        // line #2736:
	NamedConstructor, // line #2737:   Breakpoint._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #2738:     breakpointNumber = json['breakpointNumber'];
	NamedConstructor, // line #2739:     resolved = json['resolved'];
	NamedConstructor, // line #2740:     isSyntheticAsyncContinuation = json['isSyntheticAsyncContinuation'];
	NamedConstructor, // line #2741:     location = createServiceObject(
	NamedConstructor, // line #2742:         json['location'], const ['SourceLocation', 'UnresolvedSourceLocation']);
	NamedConstructor, // line #2743:   }
	BlankLine,        // line #2744:
	OverrideMethod,   // line #2745:   @override
	OverrideMethod,   // line #2746:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #2747:     var json = super.toJson();
	OverrideMethod,   // line #2748:     json['type'] = 'Breakpoint';
	OverrideMethod,   // line #2749:     json.addAll({
	OverrideMethod,   // line #2750:       'breakpointNumber': breakpointNumber,
	OverrideMethod,   // line #2751:       'resolved': resolved,
	OverrideMethod,   // line #2752:       'location': location.toJson(),
	OverrideMethod,   // line #2753:     });
	OverrideMethod,   // line #2754:     _setIfNotNull(
	OverrideMethod,   // line #2755:         json, 'isSyntheticAsyncContinuation', isSyntheticAsyncContinuation);
	OverrideMethod,   // line #2756:     return json;
	OverrideMethod,   // line #2757:   }
	BlankLine,        // line #2758:
	OtherMethod,      // line #2759:   int get hashCode => id.hashCode;
	BlankLine,        // line #2760:
	OtherMethod,      // line #2761:   operator ==(other) => other is Breakpoint && id == other.id;
	BlankLine,        // line #2762:
	OtherMethod,      // line #2763:   String toString() => '[Breakpoint ' //
	OtherMethod,      // line #2764:       'type: ${type}, id: ${id}, breakpointNumber: ${breakpointNumber}, ' //
	OtherMethod,      // line #2765:       'resolved: ${resolved}, location: ${location}]';
	BlankLine,        // line #2766:
}

var wantVMServiceClass24 = []EntityType{
	Unknown,          // line #2769: {
	OtherMethod,      // line #2770:   static ClassRef parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #2771:       json == null ? null : ClassRef._fromJson(json);
	BlankLine,        // line #2772:
	InstanceVariable, // line #2773:   /// The name of this class.
	InstanceVariable, // line #2774:   String name;
	BlankLine,        // line #2775:
	MainConstructor,  // line #2776:   ClassRef({
	MainConstructor,  // line #2777:     @required this.name,
	MainConstructor,  // line #2778:     @required String id,
	MainConstructor,  // line #2779:   }) : super(id: id);
	BlankLine,        // line #2780:
	NamedConstructor, // line #2781:   ClassRef._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #2782:     name = json['name'];
	NamedConstructor, // line #2783:   }
	BlankLine,        // line #2784:
	OverrideMethod,   // line #2785:   @override
	OverrideMethod,   // line #2786:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #2787:     var json = super.toJson();
	OverrideMethod,   // line #2788:     json['type'] = '@Class';
	OverrideMethod,   // line #2789:     json.addAll({
	OverrideMethod,   // line #2790:       'name': name,
	OverrideMethod,   // line #2791:     });
	OverrideMethod,   // line #2792:     return json;
	OverrideMethod,   // line #2793:   }
	BlankLine,        // line #2794:
	OtherMethod,      // line #2795:   int get hashCode => id.hashCode;
	BlankLine,        // line #2796:
	OtherMethod,      // line #2797:   operator ==(other) => other is ClassRef && id == other.id;
	BlankLine,        // line #2798:
	OtherMethod,      // line #2799:   String toString() => '[ClassRef type: ${type}, id: ${id}, name: ${name}]';
	BlankLine,        // line #2800:
}

var wantVMServiceClass25 = []EntityType{
	Unknown,          // line #2803: {
	OtherMethod,      // line #2804:   static Class parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #2805:       json == null ? null : Class._fromJson(json);
	BlankLine,        // line #2806:
	InstanceVariable, // line #2807:   /// The name of this class.
	InstanceVariable, // line #2808:   String name;
	BlankLine,        // line #2809:
	InstanceVariable, // line #2810:   /// The error which occurred during class finalization, if it exists.
	InstanceVariable, // line #2811:   @optional
	InstanceVariable, // line #2812:   ErrorRef error;
	BlankLine,        // line #2813:
	InstanceVariable, // line #2814:   /// Is this an abstract class?
	InstanceVariable, // line #2815:   bool isAbstract;
	BlankLine,        // line #2816:
	InstanceVariable, // line #2817:   /// Is this a const class?
	InstanceVariable, // line #2818:   bool isConst;
	BlankLine,        // line #2819:
	InstanceVariable, // line #2820:   /// The library which contains this class.
	InstanceVariable, // line #2821:   LibraryRef library;
	BlankLine,        // line #2822:
	InstanceVariable, // line #2823:   /// The location of this class in the source code.
	InstanceVariable, // line #2824:   @optional
	InstanceVariable, // line #2825:   SourceLocation location;
	BlankLine,        // line #2826:
	InstanceVariable, // line #2827:   /// The superclass of this class, if any.
	InstanceVariable, // line #2828:   @optional
	InstanceVariable, // line #2829:   ClassRef superClass;
	BlankLine,        // line #2830:
	InstanceVariable, // line #2831:   /// The supertype for this class, if any.
	InstanceVariable, // line #2832:   ///
	InstanceVariable, // line #2833:   /// The value will be of the kind: Type.
	InstanceVariable, // line #2834:   @optional
	InstanceVariable, // line #2835:   InstanceRef superType;
	BlankLine,        // line #2836:
	InstanceVariable, // line #2837:   /// A list of interface types for this class.
	InstanceVariable, // line #2838:   ///
	InstanceVariable, // line #2839:   /// The values will be of the kind: Type.
	InstanceVariable, // line #2840:   List<InstanceRef> interfaces;
	BlankLine,        // line #2841:
	InstanceVariable, // line #2842:   /// The mixin type for this class, if any.
	InstanceVariable, // line #2843:   ///
	InstanceVariable, // line #2844:   /// The value will be of the kind: Type.
	InstanceVariable, // line #2845:   @optional
	InstanceVariable, // line #2846:   InstanceRef mixin;
	BlankLine,        // line #2847:
	InstanceVariable, // line #2848:   /// A list of fields in this class. Does not include fields from superclasses.
	InstanceVariable, // line #2849:   List<FieldRef> fields;
	BlankLine,        // line #2850:
	InstanceVariable, // line #2851:   /// A list of functions in this class. Does not include functions from
	InstanceVariable, // line #2852:   /// superclasses.
	InstanceVariable, // line #2853:   List<FuncRef> functions;
	BlankLine,        // line #2854:
	InstanceVariable, // line #2855:   /// A list of subclasses of this class.
	InstanceVariable, // line #2856:   List<ClassRef> subclasses;
	BlankLine,        // line #2857:
	MainConstructor,  // line #2858:   Class({
	MainConstructor,  // line #2859:     @required this.name,
	MainConstructor,  // line #2860:     @required this.isAbstract,
	MainConstructor,  // line #2861:     @required this.isConst,
	MainConstructor,  // line #2862:     @required this.library,
	MainConstructor,  // line #2863:     @required this.interfaces,
	MainConstructor,  // line #2864:     @required this.fields,
	MainConstructor,  // line #2865:     @required this.functions,
	MainConstructor,  // line #2866:     @required this.subclasses,
	MainConstructor,  // line #2867:     @required String id,
	MainConstructor,  // line #2868:     this.error,
	MainConstructor,  // line #2869:     this.location,
	MainConstructor,  // line #2870:     this.superClass,
	MainConstructor,  // line #2871:     this.superType,
	MainConstructor,  // line #2872:     this.mixin,
	MainConstructor,  // line #2873:   }) : super(id: id);
	BlankLine,        // line #2874:
	NamedConstructor, // line #2875:   Class._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #2876:     name = json['name'];
	NamedConstructor, // line #2877:     error = createServiceObject(json['error'], const ['ErrorRef']);
	NamedConstructor, // line #2878:     isAbstract = json['abstract'];
	NamedConstructor, // line #2879:     isConst = json['const'];
	NamedConstructor, // line #2880:     library = createServiceObject(json['library'], const ['LibraryRef']);
	NamedConstructor, // line #2881:     location = createServiceObject(json['location'], const ['SourceLocation']);
	NamedConstructor, // line #2882:     superClass = createServiceObject(json['super'], const ['ClassRef']);
	NamedConstructor, // line #2883:     superType = createServiceObject(json['superType'], const ['InstanceRef']);
	NamedConstructor, // line #2884:     interfaces = List<InstanceRef>.from(
	NamedConstructor, // line #2885:         createServiceObject(json['interfaces'], const ['InstanceRef']) ?? []);
	NamedConstructor, // line #2886:     mixin = createServiceObject(json['mixin'], const ['InstanceRef']);
	NamedConstructor, // line #2887:     fields = List<FieldRef>.from(
	NamedConstructor, // line #2888:         createServiceObject(json['fields'], const ['FieldRef']) ?? []);
	NamedConstructor, // line #2889:     functions = List<FuncRef>.from(
	NamedConstructor, // line #2890:         createServiceObject(json['functions'], const ['FuncRef']) ?? []);
	NamedConstructor, // line #2891:     subclasses = List<ClassRef>.from(
	NamedConstructor, // line #2892:         createServiceObject(json['subclasses'], const ['ClassRef']) ?? []);
	NamedConstructor, // line #2893:   }
	BlankLine,        // line #2894:
	OverrideMethod,   // line #2895:   @override
	OverrideMethod,   // line #2896:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #2897:     var json = super.toJson();
	OverrideMethod,   // line #2898:     json['type'] = 'Class';
	OverrideMethod,   // line #2899:     json.addAll({
	OverrideMethod,   // line #2900:       'name': name,
	OverrideMethod,   // line #2901:       'abstract': isAbstract,
	OverrideMethod,   // line #2902:       'const': isConst,
	OverrideMethod,   // line #2903:       'library': library.toJson(),
	OverrideMethod,   // line #2904:       'interfaces': interfaces.map((f) => f.toJson()).toList(),
	OverrideMethod,   // line #2905:       'fields': fields.map((f) => f.toJson()).toList(),
	OverrideMethod,   // line #2906:       'functions': functions.map((f) => f.toJson()).toList(),
	OverrideMethod,   // line #2907:       'subclasses': subclasses.map((f) => f.toJson()).toList(),
	OverrideMethod,   // line #2908:     });
	OverrideMethod,   // line #2909:     _setIfNotNull(json, 'error', error?.toJson());
	OverrideMethod,   // line #2910:     _setIfNotNull(json, 'location', location?.toJson());
	OverrideMethod,   // line #2911:     _setIfNotNull(json, 'super', superClass?.toJson());
	OverrideMethod,   // line #2912:     _setIfNotNull(json, 'superType', superType?.toJson());
	OverrideMethod,   // line #2913:     _setIfNotNull(json, 'mixin', mixin?.toJson());
	OverrideMethod,   // line #2914:     return json;
	OverrideMethod,   // line #2915:   }
	BlankLine,        // line #2916:
	OtherMethod,      // line #2917:   int get hashCode => id.hashCode;
	BlankLine,        // line #2918:
	OtherMethod,      // line #2919:   operator ==(other) => other is Class && id == other.id;
	BlankLine,        // line #2920:
	OtherMethod,      // line #2921:   String toString() => '[Class]';
	BlankLine,        // line #2922:
}

var wantVMServiceClass26 = []EntityType{
	Unknown,          // line #2924: {
	OtherMethod,      // line #2925:   static ClassHeapStats parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #2926:       json == null ? null : ClassHeapStats._fromJson(json);
	BlankLine,        // line #2927:
	InstanceVariable, // line #2928:   /// The class for which this memory information is associated.
	InstanceVariable, // line #2929:   ClassRef classRef;
	BlankLine,        // line #2930:
	InstanceVariable, // line #2931:   /// The number of bytes allocated for instances of class since the accumulator
	InstanceVariable, // line #2932:   /// was last reset.
	InstanceVariable, // line #2933:   int accumulatedSize;
	BlankLine,        // line #2934:
	InstanceVariable, // line #2935:   /// The number of bytes currently allocated for instances of class.
	InstanceVariable, // line #2936:   int bytesCurrent;
	BlankLine,        // line #2937:
	InstanceVariable, // line #2938:   /// The number of instances of class which have been allocated since the
	InstanceVariable, // line #2939:   /// accumulator was last reset.
	InstanceVariable, // line #2940:   int instancesAccumulated;
	BlankLine,        // line #2941:
	InstanceVariable, // line #2942:   /// The number of instances of class which are currently alive.
	InstanceVariable, // line #2943:   int instancesCurrent;
	BlankLine,        // line #2944:
	MainConstructor,  // line #2945:   ClassHeapStats({
	MainConstructor,  // line #2946:     @required this.classRef,
	MainConstructor,  // line #2947:     @required this.accumulatedSize,
	MainConstructor,  // line #2948:     @required this.bytesCurrent,
	MainConstructor,  // line #2949:     @required this.instancesAccumulated,
	MainConstructor,  // line #2950:     @required this.instancesCurrent,
	MainConstructor,  // line #2951:   });
	BlankLine,        // line #2952:
	NamedConstructor, // line #2953:   ClassHeapStats._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #2954:     classRef = createServiceObject(json['class'], const ['ClassRef']);
	NamedConstructor, // line #2955:     accumulatedSize = json['accumulatedSize'];
	NamedConstructor, // line #2956:     bytesCurrent = json['bytesCurrent'];
	NamedConstructor, // line #2957:     instancesAccumulated = json['instancesAccumulated'];
	NamedConstructor, // line #2958:     instancesCurrent = json['instancesCurrent'];
	NamedConstructor, // line #2959:   }
	BlankLine,        // line #2960:
	OverrideMethod,   // line #2961:   @override
	OverrideMethod,   // line #2962:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #2963:     var json = <String, dynamic>{};
	OverrideMethod,   // line #2964:     json['type'] = 'ClassHeapStats';
	OverrideMethod,   // line #2965:     json.addAll({
	OverrideMethod,   // line #2966:       'class': classRef.toJson(),
	OverrideMethod,   // line #2967:       'accumulatedSize': accumulatedSize,
	OverrideMethod,   // line #2968:       'bytesCurrent': bytesCurrent,
	OverrideMethod,   // line #2969:       'instancesAccumulated': instancesAccumulated,
	OverrideMethod,   // line #2970:       'instancesCurrent': instancesCurrent,
	OverrideMethod,   // line #2971:     });
	OverrideMethod,   // line #2972:     return json;
	OverrideMethod,   // line #2973:   }
	BlankLine,        // line #2974:
	OtherMethod,      // line #2975:   String toString() => '[ClassHeapStats ' //
	OtherMethod,      // line #2976:       'type: ${type}, classRef: ${classRef}, accumulatedSize: ${accumulatedSize}, ' //
	OtherMethod,      // line #2977:       'bytesCurrent: ${bytesCurrent}, instancesAccumulated: ${instancesAccumulated}, instancesCurrent: ${instancesCurrent}]';
	BlankLine,        // line #2978:
}

var wantVMServiceClass27 = []EntityType{
	Unknown,          // line #2980: {
	OtherMethod,      // line #2981:   static ClassList parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #2982:       json == null ? null : ClassList._fromJson(json);
	BlankLine,        // line #2983:
	InstanceVariable, // line #2984:   List<ClassRef> classes;
	BlankLine,        // line #2985:
	MainConstructor,  // line #2986:   ClassList({
	MainConstructor,  // line #2987:     @required this.classes,
	MainConstructor,  // line #2988:   });
	BlankLine,        // line #2989:
	NamedConstructor, // line #2990:   ClassList._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #2991:     classes = List<ClassRef>.from(
	NamedConstructor, // line #2992:         createServiceObject(json['classes'], const ['ClassRef']) ?? []);
	NamedConstructor, // line #2993:   }
	BlankLine,        // line #2994:
	OverrideMethod,   // line #2995:   @override
	OverrideMethod,   // line #2996:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #2997:     var json = <String, dynamic>{};
	OverrideMethod,   // line #2998:     json['type'] = 'ClassList';
	OverrideMethod,   // line #2999:     json.addAll({
	OverrideMethod,   // line #3000:       'classes': classes.map((f) => f.toJson()).toList(),
	OverrideMethod,   // line #3001:     });
	OverrideMethod,   // line #3002:     return json;
	OverrideMethod,   // line #3003:   }
	BlankLine,        // line #3004:
	OtherMethod,      // line #3005:   String toString() => '[ClassList type: ${type}, classes: ${classes}]';
	BlankLine,        // line #3006:
}

var wantVMServiceClass28 = []EntityType{
	Unknown,           // line #3009: {
	OtherMethod,       // line #3010:   static CodeRef parse(Map<String, dynamic> json) =>
	OtherMethod,       // line #3011:       json == null ? null : CodeRef._fromJson(json);
	BlankLine,         // line #3012:
	InstanceVariable,  // line #3013:   /// A name for this code object.
	InstanceVariable,  // line #3014:   String name;
	BlankLine,         // line #3015:
	SingleLineComment, // line #3016:   /// What kind of code object is this?
	MultiLineComment,  // line #3017:   /*CodeKind*/ String kind;
	BlankLine,         // line #3018:
	MainConstructor,   // line #3019:   CodeRef({
	MainConstructor,   // line #3020:     @required this.name,
	MainConstructor,   // line #3021:     @required this.kind,
	MainConstructor,   // line #3022:     @required String id,
	MainConstructor,   // line #3023:   }) : super(id: id);
	BlankLine,         // line #3024:
	NamedConstructor,  // line #3025:   CodeRef._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor,  // line #3026:     name = json['name'];
	NamedConstructor,  // line #3027:     kind = json['kind'];
	NamedConstructor,  // line #3028:   }
	BlankLine,         // line #3029:
	OverrideMethod,    // line #3030:   @override
	OverrideMethod,    // line #3031:   Map<String, dynamic> toJson() {
	OverrideMethod,    // line #3032:     var json = super.toJson();
	OverrideMethod,    // line #3033:     json['type'] = '@Code';
	OverrideMethod,    // line #3034:     json.addAll({
	OverrideMethod,    // line #3035:       'name': name,
	OverrideMethod,    // line #3036:       'kind': kind,
	OverrideMethod,    // line #3037:     });
	OverrideMethod,    // line #3038:     return json;
	OverrideMethod,    // line #3039:   }
	BlankLine,         // line #3040:
	OtherMethod,       // line #3041:   int get hashCode => id.hashCode;
	BlankLine,         // line #3042:
	OtherMethod,       // line #3043:   operator ==(other) => other is CodeRef && id == other.id;
	BlankLine,         // line #3044:
	OtherMethod,       // line #3045:   String toString() =>
	OtherMethod,       // line #3046:       '[CodeRef type: ${type}, id: ${id}, name: ${name}, kind: ${kind}]';
	BlankLine,         // line #3047:
}

var wantVMServiceClass29 = []EntityType{
	Unknown,           // line #3050: {
	OtherMethod,       // line #3051:   static Code parse(Map<String, dynamic> json) =>
	OtherMethod,       // line #3052:       json == null ? null : Code._fromJson(json);
	BlankLine,         // line #3053:
	InstanceVariable,  // line #3054:   /// A name for this code object.
	InstanceVariable,  // line #3055:   String name;
	BlankLine,         // line #3056:
	SingleLineComment, // line #3057:   /// What kind of code object is this?
	MultiLineComment,  // line #3058:   /*CodeKind*/ String kind;
	BlankLine,         // line #3059:
	MainConstructor,   // line #3060:   Code({
	MainConstructor,   // line #3061:     @required this.name,
	MainConstructor,   // line #3062:     @required this.kind,
	MainConstructor,   // line #3063:     @required String id,
	MainConstructor,   // line #3064:   }) : super(id: id);
	BlankLine,         // line #3065:
	NamedConstructor,  // line #3066:   Code._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor,  // line #3067:     name = json['name'];
	NamedConstructor,  // line #3068:     kind = json['kind'];
	NamedConstructor,  // line #3069:   }
	BlankLine,         // line #3070:
	OverrideMethod,    // line #3071:   @override
	OverrideMethod,    // line #3072:   Map<String, dynamic> toJson() {
	OverrideMethod,    // line #3073:     var json = super.toJson();
	OverrideMethod,    // line #3074:     json['type'] = 'Code';
	OverrideMethod,    // line #3075:     json.addAll({
	OverrideMethod,    // line #3076:       'name': name,
	OverrideMethod,    // line #3077:       'kind': kind,
	OverrideMethod,    // line #3078:     });
	OverrideMethod,    // line #3079:     return json;
	OverrideMethod,    // line #3080:   }
	BlankLine,         // line #3081:
	OtherMethod,       // line #3082:   int get hashCode => id.hashCode;
	BlankLine,         // line #3083:
	OtherMethod,       // line #3084:   operator ==(other) => other is Code && id == other.id;
	BlankLine,         // line #3085:
	OtherMethod,       // line #3086:   String toString() =>
	OtherMethod,       // line #3087:       '[Code type: ${type}, id: ${id}, name: ${name}, kind: ${kind}]';
	BlankLine,         // line #3088:
}

var wantVMServiceClass30 = []EntityType{
	Unknown,          // line #3090: {
	OtherMethod,      // line #3091:   static ContextRef parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #3092:       json == null ? null : ContextRef._fromJson(json);
	BlankLine,        // line #3093:
	InstanceVariable, // line #3094:   /// The number of variables in this context.
	InstanceVariable, // line #3095:   int length;
	BlankLine,        // line #3096:
	MainConstructor,  // line #3097:   ContextRef({
	MainConstructor,  // line #3098:     @required this.length,
	MainConstructor,  // line #3099:     @required String id,
	MainConstructor,  // line #3100:   }) : super(id: id);
	BlankLine,        // line #3101:
	NamedConstructor, // line #3102:   ContextRef._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #3103:     length = json['length'];
	NamedConstructor, // line #3104:   }
	BlankLine,        // line #3105:
	OverrideMethod,   // line #3106:   @override
	OverrideMethod,   // line #3107:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #3108:     var json = super.toJson();
	OverrideMethod,   // line #3109:     json['type'] = '@Context';
	OverrideMethod,   // line #3110:     json.addAll({
	OverrideMethod,   // line #3111:       'length': length,
	OverrideMethod,   // line #3112:     });
	OverrideMethod,   // line #3113:     return json;
	OverrideMethod,   // line #3114:   }
	BlankLine,        // line #3115:
	OtherMethod,      // line #3116:   int get hashCode => id.hashCode;
	BlankLine,        // line #3117:
	OtherMethod,      // line #3118:   operator ==(other) => other is ContextRef && id == other.id;
	BlankLine,        // line #3119:
	OtherMethod,      // line #3120:   String toString() =>
	OtherMethod,      // line #3121:       '[ContextRef type: ${type}, id: ${id}, length: ${length}]';
	BlankLine,        // line #3122:
}

var wantVMServiceClass31 = []EntityType{
	Unknown,          // line #3126: {
	OtherMethod,      // line #3127:   static Context parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #3128:       json == null ? null : Context._fromJson(json);
	BlankLine,        // line #3129:
	InstanceVariable, // line #3130:   /// The number of variables in this context.
	InstanceVariable, // line #3131:   int length;
	BlankLine,        // line #3132:
	InstanceVariable, // line #3133:   /// The enclosing context for this context.
	InstanceVariable, // line #3134:   @optional
	InstanceVariable, // line #3135:   Context parent;
	BlankLine,        // line #3136:
	InstanceVariable, // line #3137:   /// The variables in this context object.
	InstanceVariable, // line #3138:   List<ContextElement> variables;
	BlankLine,        // line #3139:
	MainConstructor,  // line #3140:   Context({
	MainConstructor,  // line #3141:     @required this.length,
	MainConstructor,  // line #3142:     @required this.variables,
	MainConstructor,  // line #3143:     @required String id,
	MainConstructor,  // line #3144:     this.parent,
	MainConstructor,  // line #3145:   }) : super(id: id);
	BlankLine,        // line #3146:
	NamedConstructor, // line #3147:   Context._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #3148:     length = json['length'];
	NamedConstructor, // line #3149:     parent = createServiceObject(json['parent'], const ['Context']);
	NamedConstructor, // line #3150:     variables = List<ContextElement>.from(
	NamedConstructor, // line #3151:         createServiceObject(json['variables'], const ['ContextElement']) ?? []);
	NamedConstructor, // line #3152:   }
	BlankLine,        // line #3153:
	OverrideMethod,   // line #3154:   @override
	OverrideMethod,   // line #3155:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #3156:     var json = super.toJson();
	OverrideMethod,   // line #3157:     json['type'] = 'Context';
	OverrideMethod,   // line #3158:     json.addAll({
	OverrideMethod,   // line #3159:       'length': length,
	OverrideMethod,   // line #3160:       'variables': variables.map((f) => f.toJson()).toList(),
	OverrideMethod,   // line #3161:     });
	OverrideMethod,   // line #3162:     _setIfNotNull(json, 'parent', parent?.toJson());
	OverrideMethod,   // line #3163:     return json;
	OverrideMethod,   // line #3164:   }
	BlankLine,        // line #3165:
	OtherMethod,      // line #3166:   int get hashCode => id.hashCode;
	BlankLine,        // line #3167:
	OtherMethod,      // line #3168:   operator ==(other) => other is Context && id == other.id;
	BlankLine,        // line #3169:
	OtherMethod,      // line #3170:   String toString() => '[Context ' //
	OtherMethod,      // line #3171:       'type: ${type}, id: ${id}, length: ${length}, variables: ${variables}]';
	BlankLine,        // line #3172:
}

var wantVMServiceClass32 = []EntityType{
	Unknown,          // line #3174: {
	OtherMethod,      // line #3175:   static ContextElement parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #3176:       json == null ? null : ContextElement._fromJson(json);
	BlankLine,        // line #3177:
	InstanceVariable, // line #3178:   /// [value] can be one of [InstanceRef] or [Sentinel].
	InstanceVariable, // line #3179:   dynamic value;
	BlankLine,        // line #3180:
	MainConstructor,  // line #3181:   ContextElement({
	MainConstructor,  // line #3182:     @required this.value,
	MainConstructor,  // line #3183:   });
	BlankLine,        // line #3184:
	NamedConstructor, // line #3185:   ContextElement._fromJson(Map<String, dynamic> json) {
	NamedConstructor, // line #3186:     value =
	NamedConstructor, // line #3187:         createServiceObject(json['value'], const ['InstanceRef', 'Sentinel']);
	NamedConstructor, // line #3188:   }
	BlankLine,        // line #3189:
	OtherMethod,      // line #3190:   Map<String, dynamic> toJson() {
	OtherMethod,      // line #3191:     var json = <String, dynamic>{};
	OtherMethod,      // line #3192:     json.addAll({
	OtherMethod,      // line #3193:       'value': value.toJson(),
	OtherMethod,      // line #3194:     });
	OtherMethod,      // line #3195:     return json;
	OtherMethod,      // line #3196:   }
	BlankLine,        // line #3197:
	OtherMethod,      // line #3198:   String toString() => '[ContextElement value: ${value}]';
	BlankLine,        // line #3199:
}

var wantVMServiceClass33 = []EntityType{
	Unknown,          // line #3202: {
	OtherMethod,      // line #3203:   static CpuSamples parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #3204:       json == null ? null : CpuSamples._fromJson(json);
	BlankLine,        // line #3205:
	InstanceVariable, // line #3206:   /// The sampling rate for the profiler in microseconds.
	InstanceVariable, // line #3207:   int samplePeriod;
	BlankLine,        // line #3208:
	InstanceVariable, // line #3209:   /// The maximum possible stack depth for samples.
	InstanceVariable, // line #3210:   int maxStackDepth;
	BlankLine,        // line #3211:
	InstanceVariable, // line #3212:   /// The number of samples returned.
	InstanceVariable, // line #3213:   int sampleCount;
	BlankLine,        // line #3214:
	InstanceVariable, // line #3215:   /// The timespan the set of returned samples covers, in microseconds.
	InstanceVariable, // line #3216:   int timeSpan;
	BlankLine,        // line #3217:
	InstanceVariable, // line #3218:   /// The start of the period of time in which the returned samples were
	InstanceVariable, // line #3219:   /// collected.
	InstanceVariable, // line #3220:   int timeOriginMicros;
	BlankLine,        // line #3221:
	InstanceVariable, // line #3222:   /// The duration of time covered by the returned samples.
	InstanceVariable, // line #3223:   int timeExtentMicros;
	BlankLine,        // line #3224:
	InstanceVariable, // line #3225:   /// The process ID for the VM.
	InstanceVariable, // line #3226:   int pid;
	BlankLine,        // line #3227:
	InstanceVariable, // line #3228:   /// A list of functions seen in the relevant samples. These references can be
	InstanceVariable, // line #3229:   /// looked up using the indicies provided in a `CpuSample` `stack` to
	InstanceVariable, // line #3230:   /// determine which function was on the stack.
	InstanceVariable, // line #3231:   List<ProfileFunction> functions;
	BlankLine,        // line #3232:
	InstanceVariable, // line #3233:   /// A list of samples collected in the range `[timeOriginMicros,
	InstanceVariable, // line #3234:   /// timeOriginMicros + timeExtentMicros]`
	InstanceVariable, // line #3235:   List<CpuSample> samples;
	BlankLine,        // line #3236:
	MainConstructor,  // line #3237:   CpuSamples({
	MainConstructor,  // line #3238:     @required this.samplePeriod,
	MainConstructor,  // line #3239:     @required this.maxStackDepth,
	MainConstructor,  // line #3240:     @required this.sampleCount,
	MainConstructor,  // line #3241:     @required this.timeSpan,
	MainConstructor,  // line #3242:     @required this.timeOriginMicros,
	MainConstructor,  // line #3243:     @required this.timeExtentMicros,
	MainConstructor,  // line #3244:     @required this.pid,
	MainConstructor,  // line #3245:     @required this.functions,
	MainConstructor,  // line #3246:     @required this.samples,
	MainConstructor,  // line #3247:   });
	BlankLine,        // line #3248:
	NamedConstructor, // line #3249:   CpuSamples._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #3250:     samplePeriod = json['samplePeriod'];
	NamedConstructor, // line #3251:     maxStackDepth = json['maxStackDepth'];
	NamedConstructor, // line #3252:     sampleCount = json['sampleCount'];
	NamedConstructor, // line #3253:     timeSpan = json['timeSpan'];
	NamedConstructor, // line #3254:     timeOriginMicros = json['timeOriginMicros'];
	NamedConstructor, // line #3255:     timeExtentMicros = json['timeExtentMicros'];
	NamedConstructor, // line #3256:     pid = json['pid'];
	NamedConstructor, // line #3257:     functions = List<ProfileFunction>.from(
	NamedConstructor, // line #3258:         createServiceObject(json['functions'], const ['ProfileFunction']) ??
	NamedConstructor, // line #3259:             []);
	NamedConstructor, // line #3260:     samples = List<CpuSample>.from(
	NamedConstructor, // line #3261:         createServiceObject(json['samples'], const ['CpuSample']) ?? []);
	NamedConstructor, // line #3262:   }
	BlankLine,        // line #3263:
	OverrideMethod,   // line #3264:   @override
	OverrideMethod,   // line #3265:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #3266:     var json = <String, dynamic>{};
	OverrideMethod,   // line #3267:     json['type'] = 'CpuSamples';
	OverrideMethod,   // line #3268:     json.addAll({
	OverrideMethod,   // line #3269:       'samplePeriod': samplePeriod,
	OverrideMethod,   // line #3270:       'maxStackDepth': maxStackDepth,
	OverrideMethod,   // line #3271:       'sampleCount': sampleCount,
	OverrideMethod,   // line #3272:       'timeSpan': timeSpan,
	OverrideMethod,   // line #3273:       'timeOriginMicros': timeOriginMicros,
	OverrideMethod,   // line #3274:       'timeExtentMicros': timeExtentMicros,
	OverrideMethod,   // line #3275:       'pid': pid,
	OverrideMethod,   // line #3276:       'functions': functions.map((f) => f.toJson()).toList(),
	OverrideMethod,   // line #3277:       'samples': samples.map((f) => f.toJson()).toList(),
	OverrideMethod,   // line #3278:     });
	OverrideMethod,   // line #3279:     return json;
	OverrideMethod,   // line #3280:   }
	BlankLine,        // line #3281:
	OtherMethod,      // line #3282:   String toString() => '[CpuSamples]';
	BlankLine,        // line #3283:
}

var wantVMServiceClass34 = []EntityType{
	Unknown,          // line #3286: {
	OtherMethod,      // line #3287:   static CpuSample parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #3288:       json == null ? null : CpuSample._fromJson(json);
	BlankLine,        // line #3289:
	InstanceVariable, // line #3290:   /// The thread ID representing the thread on which this sample was collected.
	InstanceVariable, // line #3291:   int tid;
	BlankLine,        // line #3292:
	InstanceVariable, // line #3293:   /// The time this sample was collected in microseconds.
	InstanceVariable, // line #3294:   int timestamp;
	BlankLine,        // line #3295:
	InstanceVariable, // line #3296:   /// The name of VM tag set when this sample was collected. Omitted if the VM
	InstanceVariable, // line #3297:   /// tag for the sample is not considered valid.
	InstanceVariable, // line #3298:   @optional
	InstanceVariable, // line #3299:   String vmTag;
	BlankLine,        // line #3300:
	InstanceVariable, // line #3301:   /// The name of the User tag set when this sample was collected. Omitted if no
	InstanceVariable, // line #3302:   /// User tag was set when this sample was collected.
	InstanceVariable, // line #3303:   @optional
	InstanceVariable, // line #3304:   String userTag;
	BlankLine,        // line #3305:
	InstanceVariable, // line #3306:   /// Provided and set to true if the sample's stack was truncated. This can
	InstanceVariable, // line #3307:   /// happen if the stack is deeper than the `stackDepth` in the `CpuSamples`
	InstanceVariable, // line #3308:   /// response.
	InstanceVariable, // line #3309:   @optional
	InstanceVariable, // line #3310:   bool truncated;
	BlankLine,        // line #3311:
	InstanceVariable, // line #3312:   /// The call stack at the time this sample was collected. The stack is to be
	InstanceVariable, // line #3313:   /// interpreted as top to bottom. Each element in this array is a key into the
	InstanceVariable, // line #3314:   /// `functions` array in `CpuSamples`.
	InstanceVariable, // line #3315:   ///
	InstanceVariable, // line #3316:   /// Example:
	InstanceVariable, // line #3317:   ///
	InstanceVariable, // line #3318:   /// `functions[stack[0]] = @Function(bar())` `functions[stack[1]] =
	InstanceVariable, // line #3319:   /// @Function(foo())` `functions[stack[2]] = @Function(main())`
	InstanceVariable, // line #3320:   List<int> stack;
	BlankLine,        // line #3321:
	MainConstructor,  // line #3322:   CpuSample({
	MainConstructor,  // line #3323:     @required this.tid,
	MainConstructor,  // line #3324:     @required this.timestamp,
	MainConstructor,  // line #3325:     @required this.stack,
	MainConstructor,  // line #3326:     this.vmTag,
	MainConstructor,  // line #3327:     this.userTag,
	MainConstructor,  // line #3328:     this.truncated,
	MainConstructor,  // line #3329:   });
	BlankLine,        // line #3330:
	NamedConstructor, // line #3331:   CpuSample._fromJson(Map<String, dynamic> json) {
	NamedConstructor, // line #3332:     tid = json['tid'];
	NamedConstructor, // line #3333:     timestamp = json['timestamp'];
	NamedConstructor, // line #3334:     vmTag = json['vmTag'];
	NamedConstructor, // line #3335:     userTag = json['userTag'];
	NamedConstructor, // line #3336:     truncated = json['truncated'];
	NamedConstructor, // line #3337:     stack = List<int>.from(json['stack']);
	NamedConstructor, // line #3338:   }
	BlankLine,        // line #3339:
	OtherMethod,      // line #3340:   Map<String, dynamic> toJson() {
	OtherMethod,      // line #3341:     var json = <String, dynamic>{};
	OtherMethod,      // line #3342:     json.addAll({
	OtherMethod,      // line #3343:       'tid': tid,
	OtherMethod,      // line #3344:       'timestamp': timestamp,
	OtherMethod,      // line #3345:       'stack': stack.map((f) => f).toList(),
	OtherMethod,      // line #3346:     });
	OtherMethod,      // line #3347:     _setIfNotNull(json, 'vmTag', vmTag);
	OtherMethod,      // line #3348:     _setIfNotNull(json, 'userTag', userTag);
	OtherMethod,      // line #3349:     _setIfNotNull(json, 'truncated', truncated);
	OtherMethod,      // line #3350:     return json;
	OtherMethod,      // line #3351:   }
	BlankLine,        // line #3352:
	OtherMethod,      // line #3353:   String toString() =>
	OtherMethod,      // line #3354:       '[CpuSample tid: ${tid}, timestamp: ${timestamp}, stack: ${stack}]';
	BlankLine,        // line #3355:
}

var wantVMServiceClass35 = []EntityType{
	Unknown,           // line #3358: {
	OtherMethod,       // line #3359:   static ErrorRef parse(Map<String, dynamic> json) =>
	OtherMethod,       // line #3360:       json == null ? null : ErrorRef._fromJson(json);
	BlankLine,         // line #3361:
	SingleLineComment, // line #3362:   /// What kind of error is this?
	MultiLineComment,  // line #3363:   /*ErrorKind*/ String kind;
	BlankLine,         // line #3364:
	InstanceVariable,  // line #3365:   /// A description of the error.
	InstanceVariable,  // line #3366:   String message;
	BlankLine,         // line #3367:
	MainConstructor,   // line #3368:   ErrorRef({
	MainConstructor,   // line #3369:     @required this.kind,
	MainConstructor,   // line #3370:     @required this.message,
	MainConstructor,   // line #3371:     @required String id,
	MainConstructor,   // line #3372:   }) : super(id: id);
	BlankLine,         // line #3373:
	NamedConstructor,  // line #3374:   ErrorRef._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor,  // line #3375:     kind = json['kind'];
	NamedConstructor,  // line #3376:     message = json['message'];
	NamedConstructor,  // line #3377:   }
	BlankLine,         // line #3378:
	OverrideMethod,    // line #3379:   @override
	OverrideMethod,    // line #3380:   Map<String, dynamic> toJson() {
	OverrideMethod,    // line #3381:     var json = super.toJson();
	OverrideMethod,    // line #3382:     json['type'] = '@Error';
	OverrideMethod,    // line #3383:     json.addAll({
	OverrideMethod,    // line #3384:       'kind': kind,
	OverrideMethod,    // line #3385:       'message': message,
	OverrideMethod,    // line #3386:     });
	OverrideMethod,    // line #3387:     return json;
	OverrideMethod,    // line #3388:   }
	BlankLine,         // line #3389:
	OtherMethod,       // line #3390:   int get hashCode => id.hashCode;
	BlankLine,         // line #3391:
	OtherMethod,       // line #3392:   operator ==(other) => other is ErrorRef && id == other.id;
	BlankLine,         // line #3393:
	OtherMethod,       // line #3394:   String toString() =>
	OtherMethod,       // line #3395:       '[ErrorRef type: ${type}, id: ${id}, kind: ${kind}, message: ${message}]';
	BlankLine,         // line #3396:
}

var wantVMServiceClass36 = []EntityType{
	Unknown,           // line #3400: {
	OtherMethod,       // line #3401:   static Error parse(Map<String, dynamic> json) =>
	OtherMethod,       // line #3402:       json == null ? null : Error._fromJson(json);
	BlankLine,         // line #3403:
	SingleLineComment, // line #3404:   /// What kind of error is this?
	MultiLineComment,  // line #3405:   /*ErrorKind*/ String kind;
	BlankLine,         // line #3406:
	InstanceVariable,  // line #3407:   /// A description of the error.
	InstanceVariable,  // line #3408:   String message;
	BlankLine,         // line #3409:
	InstanceVariable,  // line #3410:   /// If this error is due to an unhandled exception, this is the exception
	InstanceVariable,  // line #3411:   /// thrown.
	InstanceVariable,  // line #3412:   @optional
	InstanceVariable,  // line #3413:   InstanceRef exception;
	BlankLine,         // line #3414:
	InstanceVariable,  // line #3415:   /// If this error is due to an unhandled exception, this is the stacktrace
	InstanceVariable,  // line #3416:   /// object.
	InstanceVariable,  // line #3417:   @optional
	InstanceVariable,  // line #3418:   InstanceRef stacktrace;
	BlankLine,         // line #3419:
	MainConstructor,   // line #3420:   Error({
	MainConstructor,   // line #3421:     @required this.kind,
	MainConstructor,   // line #3422:     @required this.message,
	MainConstructor,   // line #3423:     @required String id,
	MainConstructor,   // line #3424:     this.exception,
	MainConstructor,   // line #3425:     this.stacktrace,
	MainConstructor,   // line #3426:   }) : super(id: id);
	BlankLine,         // line #3427:
	NamedConstructor,  // line #3428:   Error._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor,  // line #3429:     kind = json['kind'];
	NamedConstructor,  // line #3430:     message = json['message'];
	NamedConstructor,  // line #3431:     exception = createServiceObject(json['exception'], const ['InstanceRef']);
	NamedConstructor,  // line #3432:     stacktrace = createServiceObject(json['stacktrace'], const ['InstanceRef']);
	NamedConstructor,  // line #3433:   }
	BlankLine,         // line #3434:
	OverrideMethod,    // line #3435:   @override
	OverrideMethod,    // line #3436:   Map<String, dynamic> toJson() {
	OverrideMethod,    // line #3437:     var json = super.toJson();
	OverrideMethod,    // line #3438:     json['type'] = 'Error';
	OverrideMethod,    // line #3439:     json.addAll({
	OverrideMethod,    // line #3440:       'kind': kind,
	OverrideMethod,    // line #3441:       'message': message,
	OverrideMethod,    // line #3442:     });
	OverrideMethod,    // line #3443:     _setIfNotNull(json, 'exception', exception?.toJson());
	OverrideMethod,    // line #3444:     _setIfNotNull(json, 'stacktrace', stacktrace?.toJson());
	OverrideMethod,    // line #3445:     return json;
	OverrideMethod,    // line #3446:   }
	BlankLine,         // line #3447:
	OtherMethod,       // line #3448:   int get hashCode => id.hashCode;
	BlankLine,         // line #3449:
	OtherMethod,       // line #3450:   operator ==(other) => other is Error && id == other.id;
	BlankLine,         // line #3451:
	OtherMethod,       // line #3452:   String toString() =>
	OtherMethod,       // line #3453:       '[Error type: ${type}, id: ${id}, kind: ${kind}, message: ${message}]';
	BlankLine,         // line #3454:
}

var wantVMServiceClass37 = []EntityType{
	Unknown,           // line #3461: {
	OtherMethod,       // line #3462:   static Event parse(Map<String, dynamic> json) =>
	OtherMethod,       // line #3463:       json == null ? null : Event._fromJson(json);
	BlankLine,         // line #3464:
	SingleLineComment, // line #3465:   /// What kind of event is this?
	MultiLineComment,  // line #3466:   /*EventKind*/ String kind;
	BlankLine,         // line #3467:
	InstanceVariable,  // line #3468:   /// The isolate with which this event is associated.
	InstanceVariable,  // line #3469:   ///
	InstanceVariable,  // line #3470:   /// This is provided for all event kinds except for:
	InstanceVariable,  // line #3471:   ///  - VMUpdate, VMFlagUpdate
	InstanceVariable,  // line #3472:   @optional
	InstanceVariable,  // line #3473:   IsolateRef isolate;
	BlankLine,         // line #3474:
	InstanceVariable,  // line #3475:   /// The vm with which this event is associated.
	InstanceVariable,  // line #3476:   ///
	InstanceVariable,  // line #3477:   /// This is provided for the event kind:
	InstanceVariable,  // line #3478:   ///  - VMUpdate, VMFlagUpdate
	InstanceVariable,  // line #3479:   @optional
	InstanceVariable,  // line #3480:   VMRef vm;
	BlankLine,         // line #3481:
	InstanceVariable,  // line #3482:   /// The timestamp (in milliseconds since the epoch) associated with this
	InstanceVariable,  // line #3483:   /// event. For some isolate pause events, the timestamp is from when the
	InstanceVariable,  // line #3484:   /// isolate was paused. For other events, the timestamp is from when the event
	InstanceVariable,  // line #3485:   /// was created.
	InstanceVariable,  // line #3486:   int timestamp;
	BlankLine,         // line #3487:
	InstanceVariable,  // line #3488:   /// The breakpoint which was added, removed, or resolved.
	InstanceVariable,  // line #3489:   ///
	InstanceVariable,  // line #3490:   /// This is provided for the event kinds:
	InstanceVariable,  // line #3491:   ///  - PauseBreakpoint
	InstanceVariable,  // line #3492:   ///  - BreakpointAdded
	InstanceVariable,  // line #3493:   ///  - BreakpointRemoved
	InstanceVariable,  // line #3494:   ///  - BreakpointResolved
	InstanceVariable,  // line #3495:   @optional
	InstanceVariable,  // line #3496:   Breakpoint breakpoint;
	BlankLine,         // line #3497:
	InstanceVariable,  // line #3498:   /// The list of breakpoints at which we are currently paused for a
	InstanceVariable,  // line #3499:   /// PauseBreakpoint event.
	InstanceVariable,  // line #3500:   ///
	InstanceVariable,  // line #3501:   /// This list may be empty. For example, while single-stepping, the VM sends a
	InstanceVariable,  // line #3502:   /// PauseBreakpoint event with no breakpoints.
	InstanceVariable,  // line #3503:   ///
	InstanceVariable,  // line #3504:   /// If there is more than one breakpoint set at the program position, then all
	InstanceVariable,  // line #3505:   /// of them will be provided.
	InstanceVariable,  // line #3506:   ///
	InstanceVariable,  // line #3507:   /// This is provided for the event kinds:
	InstanceVariable,  // line #3508:   ///  - PauseBreakpoint
	InstanceVariable,  // line #3509:   @optional
	InstanceVariable,  // line #3510:   List<Breakpoint> pauseBreakpoints;
	BlankLine,         // line #3511:
	InstanceVariable,  // line #3512:   /// The top stack frame associated with this event, if applicable.
	InstanceVariable,  // line #3513:   ///
	InstanceVariable,  // line #3514:   /// This is provided for the event kinds:
	InstanceVariable,  // line #3515:   ///  - PauseBreakpoint
	InstanceVariable,  // line #3516:   ///  - PauseInterrupted
	InstanceVariable,  // line #3517:   ///  - PauseException
	InstanceVariable,  // line #3518:   ///
	InstanceVariable,  // line #3519:   /// For PauseInterrupted events, there will be no top frame if the isolate is
	InstanceVariable,  // line #3520:   /// idle (waiting in the message loop).
	InstanceVariable,  // line #3521:   ///
	InstanceVariable,  // line #3522:   /// For the Resume event, the top frame is provided at all times except for
	InstanceVariable,  // line #3523:   /// the initial resume event that is delivered when an isolate begins
	InstanceVariable,  // line #3524:   /// execution.
	InstanceVariable,  // line #3525:   @optional
	InstanceVariable,  // line #3526:   Frame topFrame;
	BlankLine,         // line #3527:
	InstanceVariable,  // line #3528:   /// The exception associated with this event, if this is a PauseException
	InstanceVariable,  // line #3529:   /// event.
	InstanceVariable,  // line #3530:   @optional
	InstanceVariable,  // line #3531:   InstanceRef exception;
	BlankLine,         // line #3532:
	InstanceVariable,  // line #3533:   /// An array of bytes, encoded as a base64 string.
	InstanceVariable,  // line #3534:   ///
	InstanceVariable,  // line #3535:   /// This is provided for the WriteEvent event.
	InstanceVariable,  // line #3536:   @optional
	InstanceVariable,  // line #3537:   String bytes;
	BlankLine,         // line #3538:
	InstanceVariable,  // line #3539:   /// The argument passed to dart:developer.inspect.
	InstanceVariable,  // line #3540:   ///
	InstanceVariable,  // line #3541:   /// This is provided for the Inspect event.
	InstanceVariable,  // line #3542:   @optional
	InstanceVariable,  // line #3543:   InstanceRef inspectee;
	BlankLine,         // line #3544:
	InstanceVariable,  // line #3545:   /// The RPC name of the extension that was added.
	InstanceVariable,  // line #3546:   ///
	InstanceVariable,  // line #3547:   /// This is provided for the ServiceExtensionAdded event.
	InstanceVariable,  // line #3548:   @optional
	InstanceVariable,  // line #3549:   String extensionRPC;
	BlankLine,         // line #3550:
	InstanceVariable,  // line #3551:   /// The extension event kind.
	InstanceVariable,  // line #3552:   ///
	InstanceVariable,  // line #3553:   /// This is provided for the Extension event.
	InstanceVariable,  // line #3554:   @optional
	InstanceVariable,  // line #3555:   String extensionKind;
	BlankLine,         // line #3556:
	InstanceVariable,  // line #3557:   /// The extension event data.
	InstanceVariable,  // line #3558:   ///
	InstanceVariable,  // line #3559:   /// This is provided for the Extension event.
	InstanceVariable,  // line #3560:   @optional
	InstanceVariable,  // line #3561:   ExtensionData extensionData;
	BlankLine,         // line #3562:
	InstanceVariable,  // line #3563:   /// An array of TimelineEvents
	InstanceVariable,  // line #3564:   ///
	InstanceVariable,  // line #3565:   /// This is provided for the TimelineEvents event.
	InstanceVariable,  // line #3566:   @optional
	InstanceVariable,  // line #3567:   List<TimelineEvent> timelineEvents;
	BlankLine,         // line #3568:
	InstanceVariable,  // line #3569:   /// The new set of recorded timeline streams.
	InstanceVariable,  // line #3570:   ///
	InstanceVariable,  // line #3571:   /// This is provided for the TimelineStreamSubscriptionsUpdate event.
	InstanceVariable,  // line #3572:   @optional
	InstanceVariable,  // line #3573:   List<String> updatedStreams;
	BlankLine,         // line #3574:
	InstanceVariable,  // line #3575:   /// Is the isolate paused at an await, yield, or yield* statement?
	InstanceVariable,  // line #3576:   ///
	InstanceVariable,  // line #3577:   /// This is provided for the event kinds:
	InstanceVariable,  // line #3578:   ///  - PauseBreakpoint
	InstanceVariable,  // line #3579:   ///  - PauseInterrupted
	InstanceVariable,  // line #3580:   @optional
	InstanceVariable,  // line #3581:   bool atAsyncSuspension;
	BlankLine,         // line #3582:
	InstanceVariable,  // line #3583:   /// The status (success or failure) related to the event. This is provided for
	InstanceVariable,  // line #3584:   /// the event kinds:
	InstanceVariable,  // line #3585:   ///  - IsolateReloaded
	InstanceVariable,  // line #3586:   @optional
	InstanceVariable,  // line #3587:   String status;
	BlankLine,         // line #3588:
	InstanceVariable,  // line #3589:   /// LogRecord data.
	InstanceVariable,  // line #3590:   ///
	InstanceVariable,  // line #3591:   /// This is provided for the Logging event.
	InstanceVariable,  // line #3592:   @optional
	InstanceVariable,  // line #3593:   LogRecord logRecord;
	BlankLine,         // line #3594:
	InstanceVariable,  // line #3595:   /// The service identifier.
	InstanceVariable,  // line #3596:   ///
	InstanceVariable,  // line #3597:   /// This is provided for the event kinds:
	InstanceVariable,  // line #3598:   ///  - ServiceRegistered
	InstanceVariable,  // line #3599:   ///  - ServiceUnregistered
	InstanceVariable,  // line #3600:   @optional
	InstanceVariable,  // line #3601:   String service;
	BlankLine,         // line #3602:
	InstanceVariable,  // line #3603:   /// The RPC method that should be used to invoke the service.
	InstanceVariable,  // line #3604:   ///
	InstanceVariable,  // line #3605:   /// This is provided for the event kinds:
	InstanceVariable,  // line #3606:   ///  - ServiceRegistered
	InstanceVariable,  // line #3607:   ///  - ServiceUnregistered
	InstanceVariable,  // line #3608:   @optional
	InstanceVariable,  // line #3609:   String method;
	BlankLine,         // line #3610:
	InstanceVariable,  // line #3611:   /// The alias of the registered service.
	InstanceVariable,  // line #3612:   ///
	InstanceVariable,  // line #3613:   /// This is provided for the event kinds:
	InstanceVariable,  // line #3614:   ///  - ServiceRegistered
	InstanceVariable,  // line #3615:   @optional
	InstanceVariable,  // line #3616:   String alias;
	BlankLine,         // line #3617:
	InstanceVariable,  // line #3618:   /// The name of the changed flag.
	InstanceVariable,  // line #3619:   ///
	InstanceVariable,  // line #3620:   /// This is provided for the event kinds:
	InstanceVariable,  // line #3621:   ///  - VMFlagUpdate
	InstanceVariable,  // line #3622:   @optional
	InstanceVariable,  // line #3623:   String flag;
	BlankLine,         // line #3624:
	InstanceVariable,  // line #3625:   /// The new value of the changed flag.
	InstanceVariable,  // line #3626:   ///
	InstanceVariable,  // line #3627:   /// This is provided for the event kinds:
	InstanceVariable,  // line #3628:   ///  - VMFlagUpdate
	InstanceVariable,  // line #3629:   @optional
	InstanceVariable,  // line #3630:   String newValue;
	BlankLine,         // line #3631:
	InstanceVariable,  // line #3632:   /// Specifies whether this event is the last of a group of events.
	InstanceVariable,  // line #3633:   ///
	InstanceVariable,  // line #3634:   /// This is provided for the event kinds:
	InstanceVariable,  // line #3635:   ///  - HeapSnapshot
	InstanceVariable,  // line #3636:   @optional
	InstanceVariable,  // line #3637:   bool last;
	BlankLine,         // line #3638:
	InstanceVariable,  // line #3639:   /// Binary data associated with the event.
	InstanceVariable,  // line #3640:   ///
	InstanceVariable,  // line #3641:   /// This is provided for the event kinds:
	InstanceVariable,  // line #3642:   ///   - HeapSnapshot
	InstanceVariable,  // line #3643:   @optional
	InstanceVariable,  // line #3644:   ByteData data;
	BlankLine,         // line #3645:
	MainConstructor,   // line #3646:   Event({
	MainConstructor,   // line #3647:     @required this.kind,
	MainConstructor,   // line #3648:     @required this.timestamp,
	MainConstructor,   // line #3649:     this.isolate,
	MainConstructor,   // line #3650:     this.vm,
	MainConstructor,   // line #3651:     this.breakpoint,
	MainConstructor,   // line #3652:     this.pauseBreakpoints,
	MainConstructor,   // line #3653:     this.topFrame,
	MainConstructor,   // line #3654:     this.exception,
	MainConstructor,   // line #3655:     this.bytes,
	MainConstructor,   // line #3656:     this.inspectee,
	MainConstructor,   // line #3657:     this.extensionRPC,
	MainConstructor,   // line #3658:     this.extensionKind,
	MainConstructor,   // line #3659:     this.extensionData,
	MainConstructor,   // line #3660:     this.timelineEvents,
	MainConstructor,   // line #3661:     this.updatedStreams,
	MainConstructor,   // line #3662:     this.atAsyncSuspension,
	MainConstructor,   // line #3663:     this.status,
	MainConstructor,   // line #3664:     this.logRecord,
	MainConstructor,   // line #3665:     this.service,
	MainConstructor,   // line #3666:     this.method,
	MainConstructor,   // line #3667:     this.alias,
	MainConstructor,   // line #3668:     this.flag,
	MainConstructor,   // line #3669:     this.newValue,
	MainConstructor,   // line #3670:     this.last,
	MainConstructor,   // line #3671:     this.data,
	MainConstructor,   // line #3672:   });
	BlankLine,         // line #3673:
	NamedConstructor,  // line #3674:   Event._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor,  // line #3675:     kind = json['kind'];
	NamedConstructor,  // line #3676:     isolate = createServiceObject(json['isolate'], const ['IsolateRef']);
	NamedConstructor,  // line #3677:     vm = createServiceObject(json['vm'], const ['VMRef']);
	NamedConstructor,  // line #3678:     timestamp = json['timestamp'];
	NamedConstructor,  // line #3679:     breakpoint = createServiceObject(json['breakpoint'], const ['Breakpoint']);
	NamedConstructor,  // line #3680:     pauseBreakpoints = json['pauseBreakpoints'] == null
	NamedConstructor,  // line #3681:         ? null
	NamedConstructor,  // line #3682:         : List<Breakpoint>.from(createServiceObject(
	NamedConstructor,  // line #3683:             json['pauseBreakpoints'], const ['Breakpoint']));
	NamedConstructor,  // line #3684:     topFrame = createServiceObject(json['topFrame'], const ['Frame']);
	NamedConstructor,  // line #3685:     exception = createServiceObject(json['exception'], const ['InstanceRef']);
	NamedConstructor,  // line #3686:     bytes = json['bytes'];
	NamedConstructor,  // line #3687:     inspectee = createServiceObject(json['inspectee'], const ['InstanceRef']);
	NamedConstructor,  // line #3688:     extensionRPC = json['extensionRPC'];
	NamedConstructor,  // line #3689:     extensionKind = json['extensionKind'];
	NamedConstructor,  // line #3690:     extensionData = ExtensionData.parse(json['extensionData']);
	NamedConstructor,  // line #3691:     timelineEvents = json['timelineEvents'] == null
	NamedConstructor,  // line #3692:         ? null
	NamedConstructor,  // line #3693:         : List<TimelineEvent>.from(createServiceObject(
	NamedConstructor,  // line #3694:             json['timelineEvents'], const ['TimelineEvent']));
	NamedConstructor,  // line #3695:     updatedStreams = json['updatedStreams'] == null
	NamedConstructor,  // line #3696:         ? null
	NamedConstructor,  // line #3697:         : List<String>.from(json['updatedStreams']);
	NamedConstructor,  // line #3698:     atAsyncSuspension = json['atAsyncSuspension'];
	NamedConstructor,  // line #3699:     status = json['status'];
	NamedConstructor,  // line #3700:     logRecord = createServiceObject(json['logRecord'], const ['LogRecord']);
	NamedConstructor,  // line #3701:     service = json['service'];
	NamedConstructor,  // line #3702:     method = json['method'];
	NamedConstructor,  // line #3703:     alias = json['alias'];
	NamedConstructor,  // line #3704:     flag = json['flag'];
	NamedConstructor,  // line #3705:     newValue = json['newValue'];
	NamedConstructor,  // line #3706:     last = json['last'];
	NamedConstructor,  // line #3707:     data = json['data'];
	NamedConstructor,  // line #3708:   }
	BlankLine,         // line #3709:
	OverrideMethod,    // line #3710:   @override
	OverrideMethod,    // line #3711:   Map<String, dynamic> toJson() {
	OverrideMethod,    // line #3712:     var json = <String, dynamic>{};
	OverrideMethod,    // line #3713:     json['type'] = 'Event';
	OverrideMethod,    // line #3714:     json.addAll({
	OverrideMethod,    // line #3715:       'kind': kind,
	OverrideMethod,    // line #3716:       'timestamp': timestamp,
	OverrideMethod,    // line #3717:     });
	OverrideMethod,    // line #3718:     _setIfNotNull(json, 'isolate', isolate?.toJson());
	OverrideMethod,    // line #3719:     _setIfNotNull(json, 'vm', vm?.toJson());
	OverrideMethod,    // line #3720:     _setIfNotNull(json, 'breakpoint', breakpoint?.toJson());
	OverrideMethod,    // line #3721:     _setIfNotNull(json, 'pauseBreakpoints',
	OverrideMethod,    // line #3722:         pauseBreakpoints?.map((f) => f?.toJson())?.toList());
	OverrideMethod,    // line #3723:     _setIfNotNull(json, 'topFrame', topFrame?.toJson());
	OverrideMethod,    // line #3724:     _setIfNotNull(json, 'exception', exception?.toJson());
	OverrideMethod,    // line #3725:     _setIfNotNull(json, 'bytes', bytes);
	OverrideMethod,    // line #3726:     _setIfNotNull(json, 'inspectee', inspectee?.toJson());
	OverrideMethod,    // line #3727:     _setIfNotNull(json, 'extensionRPC', extensionRPC);
	OverrideMethod,    // line #3728:     _setIfNotNull(json, 'extensionKind', extensionKind);
	OverrideMethod,    // line #3729:     _setIfNotNull(json, 'extensionData', extensionData?.data);
	OverrideMethod,    // line #3730:     _setIfNotNull(json, 'timelineEvents',
	OverrideMethod,    // line #3731:         timelineEvents?.map((f) => f?.toJson())?.toList());
	OverrideMethod,    // line #3732:     _setIfNotNull(
	OverrideMethod,    // line #3733:         json, 'updatedStreams', updatedStreams?.map((f) => f)?.toList());
	OverrideMethod,    // line #3734:     _setIfNotNull(json, 'atAsyncSuspension', atAsyncSuspension);
	OverrideMethod,    // line #3735:     _setIfNotNull(json, 'status', status);
	OverrideMethod,    // line #3736:     _setIfNotNull(json, 'logRecord', logRecord?.toJson());
	OverrideMethod,    // line #3737:     _setIfNotNull(json, 'service', service);
	OverrideMethod,    // line #3738:     _setIfNotNull(json, 'method', method);
	OverrideMethod,    // line #3739:     _setIfNotNull(json, 'alias', alias);
	OverrideMethod,    // line #3740:     _setIfNotNull(json, 'flag', flag);
	OverrideMethod,    // line #3741:     _setIfNotNull(json, 'newValue', newValue);
	OverrideMethod,    // line #3742:     _setIfNotNull(json, 'last', last);
	OverrideMethod,    // line #3743:     _setIfNotNull(json, 'data', data);
	OverrideMethod,    // line #3744:     return json;
	OverrideMethod,    // line #3745:   }
	BlankLine,         // line #3746:
	OtherMethod,       // line #3747:   String toString() =>
	OtherMethod,       // line #3748:       '[Event type: ${type}, kind: ${kind}, timestamp: ${timestamp}]';
	BlankLine,         // line #3749:
}

var wantVMServiceClass38 = []EntityType{
	Unknown,          // line #3752: {
	OtherMethod,      // line #3753:   static FieldRef parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #3754:       json == null ? null : FieldRef._fromJson(json);
	BlankLine,        // line #3755:
	InstanceVariable, // line #3756:   /// The name of this field.
	InstanceVariable, // line #3757:   String name;
	BlankLine,        // line #3758:
	InstanceVariable, // line #3759:   /// The owner of this field, which can be either a Library or a Class.
	InstanceVariable, // line #3760:   ObjRef owner;
	BlankLine,        // line #3761:
	InstanceVariable, // line #3762:   /// The declared type of this field.
	InstanceVariable, // line #3763:   ///
	InstanceVariable, // line #3764:   /// The value will always be of one of the kinds: Type, TypeRef,
	InstanceVariable, // line #3765:   /// TypeParameter, BoundedType.
	InstanceVariable, // line #3766:   InstanceRef declaredType;
	BlankLine,        // line #3767:
	InstanceVariable, // line #3768:   /// Is this field const?
	InstanceVariable, // line #3769:   bool isConst;
	BlankLine,        // line #3770:
	InstanceVariable, // line #3771:   /// Is this field final?
	InstanceVariable, // line #3772:   bool isFinal;
	BlankLine,        // line #3773:
	InstanceVariable, // line #3774:   /// Is this field static?
	InstanceVariable, // line #3775:   bool isStatic;
	BlankLine,        // line #3776:
	MainConstructor,  // line #3777:   FieldRef({
	MainConstructor,  // line #3778:     @required this.name,
	MainConstructor,  // line #3779:     @required this.owner,
	MainConstructor,  // line #3780:     @required this.declaredType,
	MainConstructor,  // line #3781:     @required this.isConst,
	MainConstructor,  // line #3782:     @required this.isFinal,
	MainConstructor,  // line #3783:     @required this.isStatic,
	MainConstructor,  // line #3784:     @required String id,
	MainConstructor,  // line #3785:   }) : super(id: id);
	BlankLine,        // line #3786:
	NamedConstructor, // line #3787:   FieldRef._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #3788:     name = json['name'];
	NamedConstructor, // line #3789:     owner = createServiceObject(json['owner'], const ['ObjRef']);
	NamedConstructor, // line #3790:     declaredType =
	NamedConstructor, // line #3791:         createServiceObject(json['declaredType'], const ['InstanceRef']);
	NamedConstructor, // line #3792:     isConst = json['const'];
	NamedConstructor, // line #3793:     isFinal = json['final'];
	NamedConstructor, // line #3794:     isStatic = json['static'];
	NamedConstructor, // line #3795:   }
	BlankLine,        // line #3796:
	OverrideMethod,   // line #3797:   @override
	OverrideMethod,   // line #3798:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #3799:     var json = super.toJson();
	OverrideMethod,   // line #3800:     json['type'] = '@Field';
	OverrideMethod,   // line #3801:     json.addAll({
	OverrideMethod,   // line #3802:       'name': name,
	OverrideMethod,   // line #3803:       'owner': owner.toJson(),
	OverrideMethod,   // line #3804:       'declaredType': declaredType.toJson(),
	OverrideMethod,   // line #3805:       'const': isConst,
	OverrideMethod,   // line #3806:       'final': isFinal,
	OverrideMethod,   // line #3807:       'static': isStatic,
	OverrideMethod,   // line #3808:     });
	OverrideMethod,   // line #3809:     return json;
	OverrideMethod,   // line #3810:   }
	BlankLine,        // line #3811:
	OtherMethod,      // line #3812:   int get hashCode => id.hashCode;
	BlankLine,        // line #3813:
	OtherMethod,      // line #3814:   operator ==(other) => other is FieldRef && id == other.id;
	BlankLine,        // line #3815:
	OtherMethod,      // line #3816:   String toString() => '[FieldRef]';
	BlankLine,        // line #3817:
}

var wantVMServiceClass39 = []EntityType{
	Unknown,          // line #3820: {
	OtherMethod,      // line #3821:   static Field parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #3822:       json == null ? null : Field._fromJson(json);
	BlankLine,        // line #3823:
	InstanceVariable, // line #3824:   /// The name of this field.
	InstanceVariable, // line #3825:   String name;
	BlankLine,        // line #3826:
	InstanceVariable, // line #3827:   /// The owner of this field, which can be either a Library or a Class.
	InstanceVariable, // line #3828:   ObjRef owner;
	BlankLine,        // line #3829:
	InstanceVariable, // line #3830:   /// The declared type of this field.
	InstanceVariable, // line #3831:   ///
	InstanceVariable, // line #3832:   /// The value will always be of one of the kinds: Type, TypeRef,
	InstanceVariable, // line #3833:   /// TypeParameter, BoundedType.
	InstanceVariable, // line #3834:   InstanceRef declaredType;
	BlankLine,        // line #3835:
	InstanceVariable, // line #3836:   /// Is this field const?
	InstanceVariable, // line #3837:   bool isConst;
	BlankLine,        // line #3838:
	InstanceVariable, // line #3839:   /// Is this field final?
	InstanceVariable, // line #3840:   bool isFinal;
	BlankLine,        // line #3841:
	InstanceVariable, // line #3842:   /// Is this field static?
	InstanceVariable, // line #3843:   bool isStatic;
	BlankLine,        // line #3844:
	InstanceVariable, // line #3845:   /// The value of this field, if the field is static.
	InstanceVariable, // line #3846:   @optional
	InstanceVariable, // line #3847:   InstanceRef staticValue;
	BlankLine,        // line #3848:
	InstanceVariable, // line #3849:   /// The location of this field in the source code.
	InstanceVariable, // line #3850:   @optional
	InstanceVariable, // line #3851:   SourceLocation location;
	BlankLine,        // line #3852:
	MainConstructor,  // line #3853:   Field({
	MainConstructor,  // line #3854:     @required this.name,
	MainConstructor,  // line #3855:     @required this.owner,
	MainConstructor,  // line #3856:     @required this.declaredType,
	MainConstructor,  // line #3857:     @required this.isConst,
	MainConstructor,  // line #3858:     @required this.isFinal,
	MainConstructor,  // line #3859:     @required this.isStatic,
	MainConstructor,  // line #3860:     @required String id,
	MainConstructor,  // line #3861:     this.staticValue,
	MainConstructor,  // line #3862:     this.location,
	MainConstructor,  // line #3863:   }) : super(id: id);
	BlankLine,        // line #3864:
	NamedConstructor, // line #3865:   Field._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #3866:     name = json['name'];
	NamedConstructor, // line #3867:     owner = createServiceObject(json['owner'], const ['ObjRef']);
	NamedConstructor, // line #3868:     declaredType =
	NamedConstructor, // line #3869:         createServiceObject(json['declaredType'], const ['InstanceRef']);
	NamedConstructor, // line #3870:     isConst = json['const'];
	NamedConstructor, // line #3871:     isFinal = json['final'];
	NamedConstructor, // line #3872:     isStatic = json['static'];
	NamedConstructor, // line #3873:     staticValue =
	NamedConstructor, // line #3874:         createServiceObject(json['staticValue'], const ['InstanceRef']);
	NamedConstructor, // line #3875:     location = createServiceObject(json['location'], const ['SourceLocation']);
	NamedConstructor, // line #3876:   }
	BlankLine,        // line #3877:
	OverrideMethod,   // line #3878:   @override
	OverrideMethod,   // line #3879:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #3880:     var json = super.toJson();
	OverrideMethod,   // line #3881:     json['type'] = 'Field';
	OverrideMethod,   // line #3882:     json.addAll({
	OverrideMethod,   // line #3883:       'name': name,
	OverrideMethod,   // line #3884:       'owner': owner.toJson(),
	OverrideMethod,   // line #3885:       'declaredType': declaredType.toJson(),
	OverrideMethod,   // line #3886:       'const': isConst,
	OverrideMethod,   // line #3887:       'final': isFinal,
	OverrideMethod,   // line #3888:       'static': isStatic,
	OverrideMethod,   // line #3889:     });
	OverrideMethod,   // line #3890:     _setIfNotNull(json, 'staticValue', staticValue?.toJson());
	OverrideMethod,   // line #3891:     _setIfNotNull(json, 'location', location?.toJson());
	OverrideMethod,   // line #3892:     return json;
	OverrideMethod,   // line #3893:   }
	BlankLine,        // line #3894:
	OtherMethod,      // line #3895:   int get hashCode => id.hashCode;
	BlankLine,        // line #3896:
	OtherMethod,      // line #3897:   operator ==(other) => other is Field && id == other.id;
	BlankLine,        // line #3898:
	OtherMethod,      // line #3899:   String toString() => '[Field]';
	BlankLine,        // line #3900:
}

var wantVMServiceClass40 = []EntityType{
	Unknown,          // line #3903: {
	OtherMethod,      // line #3904:   static Flag parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #3905:       json == null ? null : Flag._fromJson(json);
	BlankLine,        // line #3906:
	InstanceVariable, // line #3907:   /// The name of the flag.
	InstanceVariable, // line #3908:   String name;
	BlankLine,        // line #3909:
	InstanceVariable, // line #3910:   /// A description of the flag.
	InstanceVariable, // line #3911:   String comment;
	BlankLine,        // line #3912:
	InstanceVariable, // line #3913:   /// Has this flag been modified from its default setting?
	InstanceVariable, // line #3914:   bool modified;
	BlankLine,        // line #3915:
	InstanceVariable, // line #3916:   /// The value of this flag as a string.
	InstanceVariable, // line #3917:   ///
	InstanceVariable, // line #3918:   /// If this property is absent, then the value of the flag was NULL.
	InstanceVariable, // line #3919:   @optional
	InstanceVariable, // line #3920:   String valueAsString;
	BlankLine,        // line #3921:
	MainConstructor,  // line #3922:   Flag({
	MainConstructor,  // line #3923:     @required this.name,
	MainConstructor,  // line #3924:     @required this.comment,
	MainConstructor,  // line #3925:     @required this.modified,
	MainConstructor,  // line #3926:     this.valueAsString,
	MainConstructor,  // line #3927:   });
	BlankLine,        // line #3928:
	NamedConstructor, // line #3929:   Flag._fromJson(Map<String, dynamic> json) {
	NamedConstructor, // line #3930:     name = json['name'];
	NamedConstructor, // line #3931:     comment = json['comment'];
	NamedConstructor, // line #3932:     modified = json['modified'];
	NamedConstructor, // line #3933:     valueAsString = json['valueAsString'];
	NamedConstructor, // line #3934:   }
	BlankLine,        // line #3935:
	OtherMethod,      // line #3936:   Map<String, dynamic> toJson() {
	OtherMethod,      // line #3937:     var json = <String, dynamic>{};
	OtherMethod,      // line #3938:     json.addAll({
	OtherMethod,      // line #3939:       'name': name,
	OtherMethod,      // line #3940:       'comment': comment,
	OtherMethod,      // line #3941:       'modified': modified,
	OtherMethod,      // line #3942:     });
	OtherMethod,      // line #3943:     _setIfNotNull(json, 'valueAsString', valueAsString);
	OtherMethod,      // line #3944:     return json;
	OtherMethod,      // line #3945:   }
	BlankLine,        // line #3946:
	OtherMethod,      // line #3947:   String toString() =>
	OtherMethod,      // line #3948:       '[Flag name: ${name}, comment: ${comment}, modified: ${modified}]';
	BlankLine,        // line #3949:
}

var wantVMServiceClass41 = []EntityType{
	Unknown,          // line #3952: {
	OtherMethod,      // line #3953:   static FlagList parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #3954:       json == null ? null : FlagList._fromJson(json);
	BlankLine,        // line #3955:
	InstanceVariable, // line #3956:   /// A list of all flags in the VM.
	InstanceVariable, // line #3957:   List<Flag> flags;
	BlankLine,        // line #3958:
	MainConstructor,  // line #3959:   FlagList({
	MainConstructor,  // line #3960:     @required this.flags,
	MainConstructor,  // line #3961:   });
	BlankLine,        // line #3962:
	NamedConstructor, // line #3963:   FlagList._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #3964:     flags = List<Flag>.from(
	NamedConstructor, // line #3965:         createServiceObject(json['flags'], const ['Flag']) ?? []);
	NamedConstructor, // line #3966:   }
	BlankLine,        // line #3967:
	OverrideMethod,   // line #3968:   @override
	OverrideMethod,   // line #3969:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #3970:     var json = <String, dynamic>{};
	OverrideMethod,   // line #3971:     json['type'] = 'FlagList';
	OverrideMethod,   // line #3972:     json.addAll({
	OverrideMethod,   // line #3973:       'flags': flags.map((f) => f.toJson()).toList(),
	OverrideMethod,   // line #3974:     });
	OverrideMethod,   // line #3975:     return json;
	OverrideMethod,   // line #3976:   }
	BlankLine,        // line #3977:
	OtherMethod,      // line #3978:   String toString() => '[FlagList type: ${type}, flags: ${flags}]';
	BlankLine,        // line #3979:
}

var wantVMServiceClass42 = []EntityType{
	Unknown,          // line #3981: {
	OtherMethod,      // line #3982:   static Frame parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #3983:       json == null ? null : Frame._fromJson(json);
	BlankLine,        // line #3984:
	InstanceVariable, // line #3985:   int index;
	BlankLine,        // line #3986:
	InstanceVariable, // line #3987:   @optional
	InstanceVariable, // line #3988:   FuncRef function;
	BlankLine,        // line #3989:
	InstanceVariable, // line #3990:   @optional
	InstanceVariable, // line #3991:   CodeRef code;
	BlankLine,        // line #3992:
	InstanceVariable, // line #3993:   @optional
	InstanceVariable, // line #3994:   SourceLocation location;
	BlankLine,        // line #3995:
	InstanceVariable, // line #3996:   @optional
	InstanceVariable, // line #3997:   List<BoundVariable> vars;
	BlankLine,        // line #3998:
	InstanceVariable, // line #3999:   @optional
	InstanceVariable, // line #4000:   /*FrameKind*/
	InstanceVariable, // line #4001:   String kind;
	BlankLine,        // line #4002:
	MainConstructor,  // line #4003:   Frame({
	MainConstructor,  // line #4004:     @required this.index,
	MainConstructor,  // line #4005:     this.function,
	MainConstructor,  // line #4006:     this.code,
	MainConstructor,  // line #4007:     this.location,
	MainConstructor,  // line #4008:     this.vars,
	MainConstructor,  // line #4009:     this.kind,
	MainConstructor,  // line #4010:   });
	BlankLine,        // line #4011:
	NamedConstructor, // line #4012:   Frame._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #4013:     index = json['index'];
	NamedConstructor, // line #4014:     function = createServiceObject(json['function'], const ['FuncRef']);
	NamedConstructor, // line #4015:     code = createServiceObject(json['code'], const ['CodeRef']);
	NamedConstructor, // line #4016:     location = createServiceObject(json['location'], const ['SourceLocation']);
	NamedConstructor, // line #4017:     vars = json['vars'] == null
	NamedConstructor, // line #4018:         ? null
	NamedConstructor, // line #4019:         : List<BoundVariable>.from(
	NamedConstructor, // line #4020:             createServiceObject(json['vars'], const ['BoundVariable']));
	NamedConstructor, // line #4021:     kind = json['kind'];
	NamedConstructor, // line #4022:   }
	BlankLine,        // line #4023:
	OverrideMethod,   // line #4024:   @override
	OverrideMethod,   // line #4025:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #4026:     var json = <String, dynamic>{};
	OverrideMethod,   // line #4027:     json['type'] = 'Frame';
	OverrideMethod,   // line #4028:     json.addAll({
	OverrideMethod,   // line #4029:       'index': index,
	OverrideMethod,   // line #4030:     });
	OverrideMethod,   // line #4031:     _setIfNotNull(json, 'function', function?.toJson());
	OverrideMethod,   // line #4032:     _setIfNotNull(json, 'code', code?.toJson());
	OverrideMethod,   // line #4033:     _setIfNotNull(json, 'location', location?.toJson());
	OverrideMethod,   // line #4034:     _setIfNotNull(json, 'vars', vars?.map((f) => f?.toJson())?.toList());
	OverrideMethod,   // line #4035:     _setIfNotNull(json, 'kind', kind);
	OverrideMethod,   // line #4036:     return json;
	OverrideMethod,   // line #4037:   }
	BlankLine,        // line #4038:
	OtherMethod,      // line #4039:   String toString() => '[Frame type: ${type}, index: ${index}]';
	BlankLine,        // line #4040:
}

var wantVMServiceClass43 = []EntityType{
	Unknown,          // line #4043: {
	OtherMethod,      // line #4044:   static FuncRef parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #4045:       json == null ? null : FuncRef._fromJson(json);
	BlankLine,        // line #4046:
	InstanceVariable, // line #4047:   /// The name of this function.
	InstanceVariable, // line #4048:   String name;
	BlankLine,        // line #4049:
	InstanceVariable, // line #4050:   /// The owner of this function, which can be a Library, Class, or a Function.
	InstanceVariable, // line #4051:   ///
	InstanceVariable, // line #4052:   /// [owner] can be one of [LibraryRef], [ClassRef] or [FuncRef].
	InstanceVariable, // line #4053:   dynamic owner;
	BlankLine,        // line #4054:
	InstanceVariable, // line #4055:   /// Is this function static?
	InstanceVariable, // line #4056:   bool isStatic;
	BlankLine,        // line #4057:
	InstanceVariable, // line #4058:   /// Is this function const?
	InstanceVariable, // line #4059:   bool isConst;
	BlankLine,        // line #4060:
	MainConstructor,  // line #4061:   FuncRef({
	MainConstructor,  // line #4062:     @required this.name,
	MainConstructor,  // line #4063:     @required this.owner,
	MainConstructor,  // line #4064:     @required this.isStatic,
	MainConstructor,  // line #4065:     @required this.isConst,
	MainConstructor,  // line #4066:     @required String id,
	MainConstructor,  // line #4067:   }) : super(id: id);
	BlankLine,        // line #4068:
	NamedConstructor, // line #4069:   FuncRef._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #4070:     name = json['name'];
	NamedConstructor, // line #4071:     owner = createServiceObject(
	NamedConstructor, // line #4072:         json['owner'], const ['LibraryRef', 'ClassRef', 'FuncRef']);
	NamedConstructor, // line #4073:     isStatic = json['static'];
	NamedConstructor, // line #4074:     isConst = json['const'];
	NamedConstructor, // line #4075:   }
	BlankLine,        // line #4076:
	OverrideMethod,   // line #4077:   @override
	OverrideMethod,   // line #4078:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #4079:     var json = super.toJson();
	OverrideMethod,   // line #4080:     json['type'] = '@Function';
	OverrideMethod,   // line #4081:     json.addAll({
	OverrideMethod,   // line #4082:       'name': name,
	OverrideMethod,   // line #4083:       'owner': owner.toJson(),
	OverrideMethod,   // line #4084:       'static': isStatic,
	OverrideMethod,   // line #4085:       'const': isConst,
	OverrideMethod,   // line #4086:     });
	OverrideMethod,   // line #4087:     return json;
	OverrideMethod,   // line #4088:   }
	BlankLine,        // line #4089:
	OtherMethod,      // line #4090:   int get hashCode => id.hashCode;
	BlankLine,        // line #4091:
	OtherMethod,      // line #4092:   operator ==(other) => other is FuncRef && id == other.id;
	BlankLine,        // line #4093:
	OtherMethod,      // line #4094:   String toString() => '[FuncRef ' //
	OtherMethod,      // line #4095:       'type: ${type}, id: ${id}, name: ${name}, owner: ${owner}, ' //
	OtherMethod,      // line #4096:       'isStatic: ${isStatic}, isConst: ${isConst}]';
	BlankLine,        // line #4097:
}

var wantVMServiceClass44 = []EntityType{
	Unknown,          // line #4100: {
	OtherMethod,      // line #4101:   static Func parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #4102:       json == null ? null : Func._fromJson(json);
	BlankLine,        // line #4103:
	InstanceVariable, // line #4104:   /// The name of this function.
	InstanceVariable, // line #4105:   String name;
	BlankLine,        // line #4106:
	InstanceVariable, // line #4107:   /// The owner of this function, which can be a Library, Class, or a Function.
	InstanceVariable, // line #4108:   ///
	InstanceVariable, // line #4109:   /// [owner] can be one of [LibraryRef], [ClassRef] or [FuncRef].
	InstanceVariable, // line #4110:   dynamic owner;
	BlankLine,        // line #4111:
	InstanceVariable, // line #4112:   /// Is this function static?
	InstanceVariable, // line #4113:   bool isStatic;
	BlankLine,        // line #4114:
	InstanceVariable, // line #4115:   /// Is this function const?
	InstanceVariable, // line #4116:   bool isConst;
	BlankLine,        // line #4117:
	InstanceVariable, // line #4118:   /// The location of this function in the source code.
	InstanceVariable, // line #4119:   @optional
	InstanceVariable, // line #4120:   SourceLocation location;
	BlankLine,        // line #4121:
	InstanceVariable, // line #4122:   /// The compiled code associated with this function.
	InstanceVariable, // line #4123:   @optional
	InstanceVariable, // line #4124:   CodeRef code;
	BlankLine,        // line #4125:
	MainConstructor,  // line #4126:   Func({
	MainConstructor,  // line #4127:     @required this.name,
	MainConstructor,  // line #4128:     @required this.owner,
	MainConstructor,  // line #4129:     @required this.isStatic,
	MainConstructor,  // line #4130:     @required this.isConst,
	MainConstructor,  // line #4131:     @required String id,
	MainConstructor,  // line #4132:     this.location,
	MainConstructor,  // line #4133:     this.code,
	MainConstructor,  // line #4134:   }) : super(id: id);
	BlankLine,        // line #4135:
	NamedConstructor, // line #4136:   Func._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #4137:     name = json['name'];
	NamedConstructor, // line #4138:     owner = createServiceObject(
	NamedConstructor, // line #4139:         json['owner'], const ['LibraryRef', 'ClassRef', 'FuncRef']);
	NamedConstructor, // line #4140:     isStatic = json['static'];
	NamedConstructor, // line #4141:     isConst = json['const'];
	NamedConstructor, // line #4142:     location = createServiceObject(json['location'], const ['SourceLocation']);
	NamedConstructor, // line #4143:     code = createServiceObject(json['code'], const ['CodeRef']);
	NamedConstructor, // line #4144:   }
	BlankLine,        // line #4145:
	OverrideMethod,   // line #4146:   @override
	OverrideMethod,   // line #4147:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #4148:     var json = super.toJson();
	OverrideMethod,   // line #4149:     json['type'] = 'Function';
	OverrideMethod,   // line #4150:     json.addAll({
	OverrideMethod,   // line #4151:       'name': name,
	OverrideMethod,   // line #4152:       'owner': owner.toJson(),
	OverrideMethod,   // line #4153:       'static': isStatic,
	OverrideMethod,   // line #4154:       'const': isConst,
	OverrideMethod,   // line #4155:     });
	OverrideMethod,   // line #4156:     _setIfNotNull(json, 'location', location?.toJson());
	OverrideMethod,   // line #4157:     _setIfNotNull(json, 'code', code?.toJson());
	OverrideMethod,   // line #4158:     return json;
	OverrideMethod,   // line #4159:   }
	BlankLine,        // line #4160:
	OtherMethod,      // line #4161:   int get hashCode => id.hashCode;
	BlankLine,        // line #4162:
	OtherMethod,      // line #4163:   operator ==(other) => other is Func && id == other.id;
	BlankLine,        // line #4164:
	OtherMethod,      // line #4165:   String toString() => '[Func ' //
	OtherMethod,      // line #4166:       'type: ${type}, id: ${id}, name: ${name}, owner: ${owner}, ' //
	OtherMethod,      // line #4167:       'isStatic: ${isStatic}, isConst: ${isConst}]';
	BlankLine,        // line #4168:
}

var wantVMServiceClass45 = []EntityType{
	Unknown,           // line #4171: {
	OtherMethod,       // line #4172:   static InstanceRef parse(Map<String, dynamic> json) =>
	OtherMethod,       // line #4173:       json == null ? null : InstanceRef._fromJson(json);
	BlankLine,         // line #4174:
	SingleLineComment, // line #4175:   /// What kind of instance is this?
	MultiLineComment,  // line #4176:   /*InstanceKind*/ String kind;
	BlankLine,         // line #4177:
	InstanceVariable,  // line #4178:   /// Instance references always include their class.
	InstanceVariable,  // line #4179:   ClassRef classRef;
	BlankLine,         // line #4180:
	InstanceVariable,  // line #4181:   /// The value of this instance as a string.
	InstanceVariable,  // line #4182:   ///
	InstanceVariable,  // line #4183:   /// Provided for the instance kinds:
	InstanceVariable,  // line #4184:   ///  - Null (null)
	InstanceVariable,  // line #4185:   ///  - Bool (true or false)
	InstanceVariable,  // line #4186:   ///  - Double (suitable for passing to Double.parse())
	InstanceVariable,  // line #4187:   ///  - Int (suitable for passing to int.parse())
	InstanceVariable,  // line #4188:   ///  - String (value may be truncated)
	InstanceVariable,  // line #4189:   ///  - Float32x4
	InstanceVariable,  // line #4190:   ///  - Float64x2
	InstanceVariable,  // line #4191:   ///  - Int32x4
	InstanceVariable,  // line #4192:   ///  - StackTrace
	InstanceVariable,  // line #4193:   @optional
	InstanceVariable,  // line #4194:   String valueAsString;
	BlankLine,         // line #4195:
	InstanceVariable,  // line #4196:   /// The valueAsString for String references may be truncated. If so, this
	InstanceVariable,  // line #4197:   /// property is added with the value 'true'.
	InstanceVariable,  // line #4198:   ///
	InstanceVariable,  // line #4199:   /// New code should use 'length' and 'count' instead.
	InstanceVariable,  // line #4200:   @optional
	InstanceVariable,  // line #4201:   bool valueAsStringIsTruncated;
	BlankLine,         // line #4202:
	InstanceVariable,  // line #4203:   /// The length of a List or the number of associations in a Map or the number
	InstanceVariable,  // line #4204:   /// of codeunits in a String.
	InstanceVariable,  // line #4205:   ///
	InstanceVariable,  // line #4206:   /// Provided for instance kinds:
	InstanceVariable,  // line #4207:   ///  - String
	InstanceVariable,  // line #4208:   ///  - List
	InstanceVariable,  // line #4209:   ///  - Map
	InstanceVariable,  // line #4210:   ///  - Uint8ClampedList
	InstanceVariable,  // line #4211:   ///  - Uint8List
	InstanceVariable,  // line #4212:   ///  - Uint16List
	InstanceVariable,  // line #4213:   ///  - Uint32List
	InstanceVariable,  // line #4214:   ///  - Uint64List
	InstanceVariable,  // line #4215:   ///  - Int8List
	InstanceVariable,  // line #4216:   ///  - Int16List
	InstanceVariable,  // line #4217:   ///  - Int32List
	InstanceVariable,  // line #4218:   ///  - Int64List
	InstanceVariable,  // line #4219:   ///  - Float32List
	InstanceVariable,  // line #4220:   ///  - Float64List
	InstanceVariable,  // line #4221:   ///  - Int32x4List
	InstanceVariable,  // line #4222:   ///  - Float32x4List
	InstanceVariable,  // line #4223:   ///  - Float64x2List
	InstanceVariable,  // line #4224:   @optional
	InstanceVariable,  // line #4225:   int length;
	BlankLine,         // line #4226:
	InstanceVariable,  // line #4227:   /// The name of a Type instance.
	InstanceVariable,  // line #4228:   ///
	InstanceVariable,  // line #4229:   /// Provided for instance kinds:
	InstanceVariable,  // line #4230:   ///  - Type
	InstanceVariable,  // line #4231:   @optional
	InstanceVariable,  // line #4232:   String name;
	BlankLine,         // line #4233:
	InstanceVariable,  // line #4234:   /// The corresponding Class if this Type has a resolved typeClass.
	InstanceVariable,  // line #4235:   ///
	InstanceVariable,  // line #4236:   /// Provided for instance kinds:
	InstanceVariable,  // line #4237:   ///  - Type
	InstanceVariable,  // line #4238:   @optional
	InstanceVariable,  // line #4239:   ClassRef typeClass;
	BlankLine,         // line #4240:
	InstanceVariable,  // line #4241:   /// The parameterized class of a type parameter:
	InstanceVariable,  // line #4242:   ///
	InstanceVariable,  // line #4243:   /// Provided for instance kinds:
	InstanceVariable,  // line #4244:   ///  - TypeParameter
	InstanceVariable,  // line #4245:   @optional
	InstanceVariable,  // line #4246:   ClassRef parameterizedClass;
	BlankLine,         // line #4247:
	InstanceVariable,  // line #4248:   /// The pattern of a RegExp instance.
	InstanceVariable,  // line #4249:   ///
	InstanceVariable,  // line #4250:   /// The pattern is always an instance of kind String.
	InstanceVariable,  // line #4251:   ///
	InstanceVariable,  // line #4252:   /// Provided for instance kinds:
	InstanceVariable,  // line #4253:   ///  - RegExp
	InstanceVariable,  // line #4254:   @optional
	InstanceVariable,  // line #4255:   InstanceRef pattern;
	BlankLine,         // line #4256:
	InstanceVariable,  // line #4257:   /// The function associated with a Closure instance.
	InstanceVariable,  // line #4258:   ///
	InstanceVariable,  // line #4259:   /// Provided for instance kinds:
	InstanceVariable,  // line #4260:   ///  - Closure
	InstanceVariable,  // line #4261:   @optional
	InstanceVariable,  // line #4262:   FuncRef closureFunction;
	BlankLine,         // line #4263:
	InstanceVariable,  // line #4264:   /// The context associated with a Closure instance.
	InstanceVariable,  // line #4265:   ///
	InstanceVariable,  // line #4266:   /// Provided for instance kinds:
	InstanceVariable,  // line #4267:   ///  - Closure
	InstanceVariable,  // line #4268:   @optional
	InstanceVariable,  // line #4269:   ContextRef closureContext;
	BlankLine,         // line #4270:
	InstanceVariable,  // line #4271:   /// The port ID for a ReceivePort.
	InstanceVariable,  // line #4272:   ///
	InstanceVariable,  // line #4273:   /// Provided for instance kinds:
	InstanceVariable,  // line #4274:   ///  - ReceivePort
	InstanceVariable,  // line #4275:   @optional
	InstanceVariable,  // line #4276:   int portId;
	BlankLine,         // line #4277:
	InstanceVariable,  // line #4278:   /// The stack trace associated with the allocation of a ReceivePort.
	InstanceVariable,  // line #4279:   ///
	InstanceVariable,  // line #4280:   /// Provided for instance kinds:
	InstanceVariable,  // line #4281:   ///  - ReceivePort
	InstanceVariable,  // line #4282:   @optional
	InstanceVariable,  // line #4283:   InstanceRef allocationLocation;
	BlankLine,         // line #4284:
	InstanceVariable,  // line #4285:   /// A name associated with a ReceivePort used for debugging purposes.
	InstanceVariable,  // line #4286:   ///
	InstanceVariable,  // line #4287:   /// Provided for instance kinds:
	InstanceVariable,  // line #4288:   ///  - ReceivePort
	InstanceVariable,  // line #4289:   @optional
	InstanceVariable,  // line #4290:   String debugName;
	BlankLine,         // line #4291:
	MainConstructor,   // line #4292:   InstanceRef({
	MainConstructor,   // line #4293:     @required this.kind,
	MainConstructor,   // line #4294:     @required this.classRef,
	MainConstructor,   // line #4295:     @required String id,
	MainConstructor,   // line #4296:     this.valueAsString,
	MainConstructor,   // line #4297:     this.valueAsStringIsTruncated,
	MainConstructor,   // line #4298:     this.length,
	MainConstructor,   // line #4299:     this.name,
	MainConstructor,   // line #4300:     this.typeClass,
	MainConstructor,   // line #4301:     this.parameterizedClass,
	MainConstructor,   // line #4302:     this.pattern,
	MainConstructor,   // line #4303:     this.closureFunction,
	MainConstructor,   // line #4304:     this.closureContext,
	MainConstructor,   // line #4305:     this.portId,
	MainConstructor,   // line #4306:     this.allocationLocation,
	MainConstructor,   // line #4307:     this.debugName,
	MainConstructor,   // line #4308:   }) : super(id: id);
	BlankLine,         // line #4309:
	NamedConstructor,  // line #4310:   InstanceRef._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor,  // line #4311:     kind = json['kind'];
	NamedConstructor,  // line #4312:     classRef = createServiceObject(json['class'], const ['ClassRef']);
	NamedConstructor,  // line #4313:     valueAsString = json['valueAsString'];
	NamedConstructor,  // line #4314:     valueAsStringIsTruncated = json['valueAsStringIsTruncated'];
	NamedConstructor,  // line #4315:     length = json['length'];
	NamedConstructor,  // line #4316:     name = json['name'];
	NamedConstructor,  // line #4317:     typeClass = createServiceObject(json['typeClass'], const ['ClassRef']);
	NamedConstructor,  // line #4318:     parameterizedClass =
	NamedConstructor,  // line #4319:         createServiceObject(json['parameterizedClass'], const ['ClassRef']);
	NamedConstructor,  // line #4320:     pattern = createServiceObject(json['pattern'], const ['InstanceRef']);
	NamedConstructor,  // line #4321:     closureFunction =
	NamedConstructor,  // line #4322:         createServiceObject(json['closureFunction'], const ['FuncRef']);
	NamedConstructor,  // line #4323:     closureContext =
	NamedConstructor,  // line #4324:         createServiceObject(json['closureContext'], const ['ContextRef']);
	NamedConstructor,  // line #4325:     portId = json['portId'];
	NamedConstructor,  // line #4326:     allocationLocation =
	NamedConstructor,  // line #4327:         createServiceObject(json['allocationLocation'], const ['InstanceRef']);
	NamedConstructor,  // line #4328:     debugName = json['debugName'];
	NamedConstructor,  // line #4329:   }
	BlankLine,         // line #4330:
	OverrideMethod,    // line #4331:   @override
	OverrideMethod,    // line #4332:   Map<String, dynamic> toJson() {
	OverrideMethod,    // line #4333:     var json = super.toJson();
	OverrideMethod,    // line #4334:     json['type'] = '@Instance';
	OverrideMethod,    // line #4335:     json.addAll({
	OverrideMethod,    // line #4336:       'kind': kind,
	OverrideMethod,    // line #4337:       'class': classRef.toJson(),
	OverrideMethod,    // line #4338:     });
	OverrideMethod,    // line #4339:     _setIfNotNull(json, 'valueAsString', valueAsString);
	OverrideMethod,    // line #4340:     _setIfNotNull(json, 'valueAsStringIsTruncated', valueAsStringIsTruncated);
	OverrideMethod,    // line #4341:     _setIfNotNull(json, 'length', length);
	OverrideMethod,    // line #4342:     _setIfNotNull(json, 'name', name);
	OverrideMethod,    // line #4343:     _setIfNotNull(json, 'typeClass', typeClass?.toJson());
	OverrideMethod,    // line #4344:     _setIfNotNull(json, 'parameterizedClass', parameterizedClass?.toJson());
	OverrideMethod,    // line #4345:     _setIfNotNull(json, 'pattern', pattern?.toJson());
	OverrideMethod,    // line #4346:     _setIfNotNull(json, 'closureFunction', closureFunction?.toJson());
	OverrideMethod,    // line #4347:     _setIfNotNull(json, 'closureContext', closureContext?.toJson());
	OverrideMethod,    // line #4348:     _setIfNotNull(json, 'portId', portId);
	OverrideMethod,    // line #4349:     _setIfNotNull(json, 'allocationLocation', allocationLocation?.toJson());
	OverrideMethod,    // line #4350:     _setIfNotNull(json, 'debugName', debugName);
	OverrideMethod,    // line #4351:     return json;
	OverrideMethod,    // line #4352:   }
	BlankLine,         // line #4353:
	OtherMethod,       // line #4354:   int get hashCode => id.hashCode;
	BlankLine,         // line #4355:
	OtherMethod,       // line #4356:   operator ==(other) => other is InstanceRef && id == other.id;
	BlankLine,         // line #4357:
	OtherMethod,       // line #4358:   String toString() => '[InstanceRef ' //
	OtherMethod,       // line #4359:       'type: ${type}, id: ${id}, kind: ${kind}, classRef: ${classRef}]';
	BlankLine,         // line #4360:
}

var wantVMServiceClass46 = []EntityType{
	Unknown,           // line #4363: {
	OtherMethod,       // line #4364:   static Instance parse(Map<String, dynamic> json) =>
	OtherMethod,       // line #4365:       json == null ? null : Instance._fromJson(json);
	BlankLine,         // line #4366:
	SingleLineComment, // line #4367:   /// What kind of instance is this?
	MultiLineComment,  // line #4368:   /*InstanceKind*/ String kind;
	BlankLine,         // line #4369:
	OverrideVariable,  // line #4370:   /// Instance references always include their class.
	OverrideVariable,  // line #4371:   @override
	OverrideVariable,  // line #4372:   ClassRef classRef;
	BlankLine,         // line #4373:
	InstanceVariable,  // line #4374:   /// The value of this instance as a string.
	InstanceVariable,  // line #4375:   ///
	InstanceVariable,  // line #4376:   /// Provided for the instance kinds:
	InstanceVariable,  // line #4377:   ///  - Bool (true or false)
	InstanceVariable,  // line #4378:   ///  - Double (suitable for passing to Double.parse())
	InstanceVariable,  // line #4379:   ///  - Int (suitable for passing to int.parse())
	InstanceVariable,  // line #4380:   ///  - String (value may be truncated)
	InstanceVariable,  // line #4381:   ///  - StackTrace
	InstanceVariable,  // line #4382:   @optional
	InstanceVariable,  // line #4383:   String valueAsString;
	BlankLine,         // line #4384:
	InstanceVariable,  // line #4385:   /// The valueAsString for String references may be truncated. If so, this
	InstanceVariable,  // line #4386:   /// property is added with the value 'true'.
	InstanceVariable,  // line #4387:   ///
	InstanceVariable,  // line #4388:   /// New code should use 'length' and 'count' instead.
	InstanceVariable,  // line #4389:   @optional
	InstanceVariable,  // line #4390:   bool valueAsStringIsTruncated;
	BlankLine,         // line #4391:
	InstanceVariable,  // line #4392:   /// The length of a List or the number of associations in a Map or the number
	InstanceVariable,  // line #4393:   /// of codeunits in a String.
	InstanceVariable,  // line #4394:   ///
	InstanceVariable,  // line #4395:   /// Provided for instance kinds:
	InstanceVariable,  // line #4396:   ///  - String
	InstanceVariable,  // line #4397:   ///  - List
	InstanceVariable,  // line #4398:   ///  - Map
	InstanceVariable,  // line #4399:   ///  - Uint8ClampedList
	InstanceVariable,  // line #4400:   ///  - Uint8List
	InstanceVariable,  // line #4401:   ///  - Uint16List
	InstanceVariable,  // line #4402:   ///  - Uint32List
	InstanceVariable,  // line #4403:   ///  - Uint64List
	InstanceVariable,  // line #4404:   ///  - Int8List
	InstanceVariable,  // line #4405:   ///  - Int16List
	InstanceVariable,  // line #4406:   ///  - Int32List
	InstanceVariable,  // line #4407:   ///  - Int64List
	InstanceVariable,  // line #4408:   ///  - Float32List
	InstanceVariable,  // line #4409:   ///  - Float64List
	InstanceVariable,  // line #4410:   ///  - Int32x4List
	InstanceVariable,  // line #4411:   ///  - Float32x4List
	InstanceVariable,  // line #4412:   ///  - Float64x2List
	InstanceVariable,  // line #4413:   @optional
	InstanceVariable,  // line #4414:   int length;
	BlankLine,         // line #4415:
	InstanceVariable,  // line #4416:   /// The index of the first element or association or codeunit returned. This
	InstanceVariable,  // line #4417:   /// is only provided when it is non-zero.
	InstanceVariable,  // line #4418:   ///
	InstanceVariable,  // line #4419:   /// Provided for instance kinds:
	InstanceVariable,  // line #4420:   ///  - String
	InstanceVariable,  // line #4421:   ///  - List
	InstanceVariable,  // line #4422:   ///  - Map
	InstanceVariable,  // line #4423:   ///  - Uint8ClampedList
	InstanceVariable,  // line #4424:   ///  - Uint8List
	InstanceVariable,  // line #4425:   ///  - Uint16List
	InstanceVariable,  // line #4426:   ///  - Uint32List
	InstanceVariable,  // line #4427:   ///  - Uint64List
	InstanceVariable,  // line #4428:   ///  - Int8List
	InstanceVariable,  // line #4429:   ///  - Int16List
	InstanceVariable,  // line #4430:   ///  - Int32List
	InstanceVariable,  // line #4431:   ///  - Int64List
	InstanceVariable,  // line #4432:   ///  - Float32List
	InstanceVariable,  // line #4433:   ///  - Float64List
	InstanceVariable,  // line #4434:   ///  - Int32x4List
	InstanceVariable,  // line #4435:   ///  - Float32x4List
	InstanceVariable,  // line #4436:   ///  - Float64x2List
	InstanceVariable,  // line #4437:   @optional
	InstanceVariable,  // line #4438:   int offset;
	BlankLine,         // line #4439:
	InstanceVariable,  // line #4440:   /// The number of elements or associations or codeunits returned. This is only
	InstanceVariable,  // line #4441:   /// provided when it is less than length.
	InstanceVariable,  // line #4442:   ///
	InstanceVariable,  // line #4443:   /// Provided for instance kinds:
	InstanceVariable,  // line #4444:   ///  - String
	InstanceVariable,  // line #4445:   ///  - List
	InstanceVariable,  // line #4446:   ///  - Map
	InstanceVariable,  // line #4447:   ///  - Uint8ClampedList
	InstanceVariable,  // line #4448:   ///  - Uint8List
	InstanceVariable,  // line #4449:   ///  - Uint16List
	InstanceVariable,  // line #4450:   ///  - Uint32List
	InstanceVariable,  // line #4451:   ///  - Uint64List
	InstanceVariable,  // line #4452:   ///  - Int8List
	InstanceVariable,  // line #4453:   ///  - Int16List
	InstanceVariable,  // line #4454:   ///  - Int32List
	InstanceVariable,  // line #4455:   ///  - Int64List
	InstanceVariable,  // line #4456:   ///  - Float32List
	InstanceVariable,  // line #4457:   ///  - Float64List
	InstanceVariable,  // line #4458:   ///  - Int32x4List
	InstanceVariable,  // line #4459:   ///  - Float32x4List
	InstanceVariable,  // line #4460:   ///  - Float64x2List
	InstanceVariable,  // line #4461:   @optional
	InstanceVariable,  // line #4462:   int count;
	BlankLine,         // line #4463:
	InstanceVariable,  // line #4464:   /// The name of a Type instance.
	InstanceVariable,  // line #4465:   ///
	InstanceVariable,  // line #4466:   /// Provided for instance kinds:
	InstanceVariable,  // line #4467:   ///  - Type
	InstanceVariable,  // line #4468:   @optional
	InstanceVariable,  // line #4469:   String name;
	BlankLine,         // line #4470:
	InstanceVariable,  // line #4471:   /// The corresponding Class if this Type is canonical.
	InstanceVariable,  // line #4472:   ///
	InstanceVariable,  // line #4473:   /// Provided for instance kinds:
	InstanceVariable,  // line #4474:   ///  - Type
	InstanceVariable,  // line #4475:   @optional
	InstanceVariable,  // line #4476:   ClassRef typeClass;
	BlankLine,         // line #4477:
	InstanceVariable,  // line #4478:   /// The parameterized class of a type parameter:
	InstanceVariable,  // line #4479:   ///
	InstanceVariable,  // line #4480:   /// Provided for instance kinds:
	InstanceVariable,  // line #4481:   ///  - TypeParameter
	InstanceVariable,  // line #4482:   @optional
	InstanceVariable,  // line #4483:   ClassRef parameterizedClass;
	BlankLine,         // line #4484:
	InstanceVariable,  // line #4485:   /// The fields of this Instance.
	InstanceVariable,  // line #4486:   @optional
	InstanceVariable,  // line #4487:   List<BoundField> fields;
	BlankLine,         // line #4488:
	InstanceVariable,  // line #4489:   /// The elements of a List instance.
	InstanceVariable,  // line #4490:   ///
	InstanceVariable,  // line #4491:   /// Provided for instance kinds:
	InstanceVariable,  // line #4492:   ///  - List
	InstanceVariable,  // line #4493:   @optional
	InstanceVariable,  // line #4494:   List<dynamic> elements;
	BlankLine,         // line #4495:
	InstanceVariable,  // line #4496:   /// The elements of a Map instance.
	InstanceVariable,  // line #4497:   ///
	InstanceVariable,  // line #4498:   /// Provided for instance kinds:
	InstanceVariable,  // line #4499:   ///  - Map
	InstanceVariable,  // line #4500:   @optional
	InstanceVariable,  // line #4501:   List<MapAssociation> associations;
	BlankLine,         // line #4502:
	InstanceVariable,  // line #4503:   /// The bytes of a TypedData instance.
	InstanceVariable,  // line #4504:   ///
	InstanceVariable,  // line #4505:   /// The data is provided as a Base64 encoded string.
	InstanceVariable,  // line #4506:   ///
	InstanceVariable,  // line #4507:   /// Provided for instance kinds:
	InstanceVariable,  // line #4508:   ///  - Uint8ClampedList
	InstanceVariable,  // line #4509:   ///  - Uint8List
	InstanceVariable,  // line #4510:   ///  - Uint16List
	InstanceVariable,  // line #4511:   ///  - Uint32List
	InstanceVariable,  // line #4512:   ///  - Uint64List
	InstanceVariable,  // line #4513:   ///  - Int8List
	InstanceVariable,  // line #4514:   ///  - Int16List
	InstanceVariable,  // line #4515:   ///  - Int32List
	InstanceVariable,  // line #4516:   ///  - Int64List
	InstanceVariable,  // line #4517:   ///  - Float32List
	InstanceVariable,  // line #4518:   ///  - Float64List
	InstanceVariable,  // line #4519:   ///  - Int32x4List
	InstanceVariable,  // line #4520:   ///  - Float32x4List
	InstanceVariable,  // line #4521:   ///  - Float64x2List
	InstanceVariable,  // line #4522:   @optional
	InstanceVariable,  // line #4523:   String bytes;
	BlankLine,         // line #4524:
	InstanceVariable,  // line #4525:   /// The referent of a MirrorReference instance.
	InstanceVariable,  // line #4526:   ///
	InstanceVariable,  // line #4527:   /// Provided for instance kinds:
	InstanceVariable,  // line #4528:   ///  - MirrorReference
	InstanceVariable,  // line #4529:   @optional
	InstanceVariable,  // line #4530:   InstanceRef mirrorReferent;
	BlankLine,         // line #4531:
	InstanceVariable,  // line #4532:   /// The pattern of a RegExp instance.
	InstanceVariable,  // line #4533:   ///
	InstanceVariable,  // line #4534:   /// Provided for instance kinds:
	InstanceVariable,  // line #4535:   ///  - RegExp
	InstanceVariable,  // line #4536:   @optional
	InstanceVariable,  // line #4537:   InstanceRef pattern;
	BlankLine,         // line #4538:
	InstanceVariable,  // line #4539:   /// The function associated with a Closure instance.
	InstanceVariable,  // line #4540:   ///
	InstanceVariable,  // line #4541:   /// Provided for instance kinds:
	InstanceVariable,  // line #4542:   ///  - Closure
	InstanceVariable,  // line #4543:   @optional
	InstanceVariable,  // line #4544:   FuncRef closureFunction;
	BlankLine,         // line #4545:
	InstanceVariable,  // line #4546:   /// The context associated with a Closure instance.
	InstanceVariable,  // line #4547:   ///
	InstanceVariable,  // line #4548:   /// Provided for instance kinds:
	InstanceVariable,  // line #4549:   ///  - Closure
	InstanceVariable,  // line #4550:   @optional
	InstanceVariable,  // line #4551:   ContextRef closureContext;
	BlankLine,         // line #4552:
	InstanceVariable,  // line #4553:   /// Whether this regular expression is case sensitive.
	InstanceVariable,  // line #4554:   ///
	InstanceVariable,  // line #4555:   /// Provided for instance kinds:
	InstanceVariable,  // line #4556:   ///  - RegExp
	InstanceVariable,  // line #4557:   @optional
	InstanceVariable,  // line #4558:   bool isCaseSensitive;
	BlankLine,         // line #4559:
	InstanceVariable,  // line #4560:   /// Whether this regular expression matches multiple lines.
	InstanceVariable,  // line #4561:   ///
	InstanceVariable,  // line #4562:   /// Provided for instance kinds:
	InstanceVariable,  // line #4563:   ///  - RegExp
	InstanceVariable,  // line #4564:   @optional
	InstanceVariable,  // line #4565:   bool isMultiLine;
	BlankLine,         // line #4566:
	InstanceVariable,  // line #4567:   /// The key for a WeakProperty instance.
	InstanceVariable,  // line #4568:   ///
	InstanceVariable,  // line #4569:   /// Provided for instance kinds:
	InstanceVariable,  // line #4570:   ///  - WeakProperty
	InstanceVariable,  // line #4571:   @optional
	InstanceVariable,  // line #4572:   InstanceRef propertyKey;
	BlankLine,         // line #4573:
	InstanceVariable,  // line #4574:   /// The key for a WeakProperty instance.
	InstanceVariable,  // line #4575:   ///
	InstanceVariable,  // line #4576:   /// Provided for instance kinds:
	InstanceVariable,  // line #4577:   ///  - WeakProperty
	InstanceVariable,  // line #4578:   @optional
	InstanceVariable,  // line #4579:   InstanceRef propertyValue;
	BlankLine,         // line #4580:
	InstanceVariable,  // line #4581:   /// The type arguments for this type.
	InstanceVariable,  // line #4582:   ///
	InstanceVariable,  // line #4583:   /// Provided for instance kinds:
	InstanceVariable,  // line #4584:   ///  - Type
	InstanceVariable,  // line #4585:   @optional
	InstanceVariable,  // line #4586:   TypeArgumentsRef typeArguments;
	BlankLine,         // line #4587:
	InstanceVariable,  // line #4588:   /// The index of a TypeParameter instance.
	InstanceVariable,  // line #4589:   ///
	InstanceVariable,  // line #4590:   /// Provided for instance kinds:
	InstanceVariable,  // line #4591:   ///  - TypeParameter
	InstanceVariable,  // line #4592:   @optional
	InstanceVariable,  // line #4593:   int parameterIndex;
	BlankLine,         // line #4594:
	InstanceVariable,  // line #4595:   /// The type bounded by a BoundedType instance - or - the referent of a
	InstanceVariable,  // line #4596:   /// TypeRef instance.
	InstanceVariable,  // line #4597:   ///
	InstanceVariable,  // line #4598:   /// The value will always be of one of the kinds: Type, TypeRef,
	InstanceVariable,  // line #4599:   /// TypeParameter, BoundedType.
	InstanceVariable,  // line #4600:   ///
	InstanceVariable,  // line #4601:   /// Provided for instance kinds:
	InstanceVariable,  // line #4602:   ///  - BoundedType
	InstanceVariable,  // line #4603:   ///  - TypeRef
	InstanceVariable,  // line #4604:   @optional
	InstanceVariable,  // line #4605:   InstanceRef targetType;
	BlankLine,         // line #4606:
	InstanceVariable,  // line #4607:   /// The bound of a TypeParameter or BoundedType.
	InstanceVariable,  // line #4608:   ///
	InstanceVariable,  // line #4609:   /// The value will always be of one of the kinds: Type, TypeRef,
	InstanceVariable,  // line #4610:   /// TypeParameter, BoundedType.
	InstanceVariable,  // line #4611:   ///
	InstanceVariable,  // line #4612:   /// Provided for instance kinds:
	InstanceVariable,  // line #4613:   ///  - BoundedType
	InstanceVariable,  // line #4614:   ///  - TypeParameter
	InstanceVariable,  // line #4615:   @optional
	InstanceVariable,  // line #4616:   InstanceRef bound;
	BlankLine,         // line #4617:
	InstanceVariable,  // line #4618:   /// The port ID for a ReceivePort.
	InstanceVariable,  // line #4619:   ///
	InstanceVariable,  // line #4620:   /// Provided for instance kinds:
	InstanceVariable,  // line #4621:   ///  - ReceivePort
	InstanceVariable,  // line #4622:   @optional
	InstanceVariable,  // line #4623:   int portId;
	BlankLine,         // line #4624:
	InstanceVariable,  // line #4625:   /// The stack trace associated with the allocation of a ReceivePort.
	InstanceVariable,  // line #4626:   ///
	InstanceVariable,  // line #4627:   /// Provided for instance kinds:
	InstanceVariable,  // line #4628:   ///  - ReceivePort
	InstanceVariable,  // line #4629:   @optional
	InstanceVariable,  // line #4630:   InstanceRef allocationLocation;
	BlankLine,         // line #4631:
	InstanceVariable,  // line #4632:   /// A name associated with a ReceivePort used for debugging purposes.
	InstanceVariable,  // line #4633:   ///
	InstanceVariable,  // line #4634:   /// Provided for instance kinds:
	InstanceVariable,  // line #4635:   ///  - ReceivePort
	InstanceVariable,  // line #4636:   @optional
	InstanceVariable,  // line #4637:   String debugName;
	BlankLine,         // line #4638:
	MainConstructor,   // line #4639:   Instance({
	MainConstructor,   // line #4640:     @required this.kind,
	MainConstructor,   // line #4641:     @required this.classRef,
	MainConstructor,   // line #4642:     @required String id,
	MainConstructor,   // line #4643:     this.valueAsString,
	MainConstructor,   // line #4644:     this.valueAsStringIsTruncated,
	MainConstructor,   // line #4645:     this.length,
	MainConstructor,   // line #4646:     this.offset,
	MainConstructor,   // line #4647:     this.count,
	MainConstructor,   // line #4648:     this.name,
	MainConstructor,   // line #4649:     this.typeClass,
	MainConstructor,   // line #4650:     this.parameterizedClass,
	MainConstructor,   // line #4651:     this.fields,
	MainConstructor,   // line #4652:     this.elements,
	MainConstructor,   // line #4653:     this.associations,
	MainConstructor,   // line #4654:     this.bytes,
	MainConstructor,   // line #4655:     this.mirrorReferent,
	MainConstructor,   // line #4656:     this.pattern,
	MainConstructor,   // line #4657:     this.closureFunction,
	MainConstructor,   // line #4658:     this.closureContext,
	MainConstructor,   // line #4659:     this.isCaseSensitive,
	MainConstructor,   // line #4660:     this.isMultiLine,
	MainConstructor,   // line #4661:     this.propertyKey,
	MainConstructor,   // line #4662:     this.propertyValue,
	MainConstructor,   // line #4663:     this.typeArguments,
	MainConstructor,   // line #4664:     this.parameterIndex,
	MainConstructor,   // line #4665:     this.targetType,
	MainConstructor,   // line #4666:     this.bound,
	MainConstructor,   // line #4667:     this.portId,
	MainConstructor,   // line #4668:     this.allocationLocation,
	MainConstructor,   // line #4669:     this.debugName,
	MainConstructor,   // line #4670:   }) : super(id: id);
	BlankLine,         // line #4671:
	NamedConstructor,  // line #4672:   Instance._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor,  // line #4673:     kind = json['kind'];
	NamedConstructor,  // line #4674:     classRef = createServiceObject(json['class'], const ['ClassRef']);
	NamedConstructor,  // line #4675:     valueAsString = json['valueAsString'];
	NamedConstructor,  // line #4676:     valueAsStringIsTruncated = json['valueAsStringIsTruncated'];
	NamedConstructor,  // line #4677:     length = json['length'];
	NamedConstructor,  // line #4678:     offset = json['offset'];
	NamedConstructor,  // line #4679:     count = json['count'];
	NamedConstructor,  // line #4680:     name = json['name'];
	NamedConstructor,  // line #4681:     typeClass = createServiceObject(json['typeClass'], const ['ClassRef']);
	NamedConstructor,  // line #4682:     parameterizedClass =
	NamedConstructor,  // line #4683:         createServiceObject(json['parameterizedClass'], const ['ClassRef']);
	NamedConstructor,  // line #4684:     fields = json['fields'] == null
	NamedConstructor,  // line #4685:         ? null
	NamedConstructor,  // line #4686:         : List<BoundField>.from(
	NamedConstructor,  // line #4687:             createServiceObject(json['fields'], const ['BoundField']));
	NamedConstructor,  // line #4688:     elements = json['elements'] == null
	NamedConstructor,  // line #4689:         ? null
	NamedConstructor,  // line #4690:         : List<dynamic>.from(
	NamedConstructor,  // line #4691:             createServiceObject(json['elements'], const ['dynamic']));
	NamedConstructor,  // line #4692:     associations = json['associations'] == null
	NamedConstructor,  // line #4693:         ? null
	NamedConstructor,  // line #4694:         : List<MapAssociation>.from(
	NamedConstructor,  // line #4695:             _createSpecificObject(json['associations'], MapAssociation.parse));
	NamedConstructor,  // line #4696:     bytes = json['bytes'];
	NamedConstructor,  // line #4697:     mirrorReferent =
	NamedConstructor,  // line #4698:         createServiceObject(json['mirrorReferent'], const ['InstanceRef']);
	NamedConstructor,  // line #4699:     pattern = createServiceObject(json['pattern'], const ['InstanceRef']);
	NamedConstructor,  // line #4700:     closureFunction =
	NamedConstructor,  // line #4701:         createServiceObject(json['closureFunction'], const ['FuncRef']);
	NamedConstructor,  // line #4702:     closureContext =
	NamedConstructor,  // line #4703:         createServiceObject(json['closureContext'], const ['ContextRef']);
	NamedConstructor,  // line #4704:     isCaseSensitive = json['isCaseSensitive'];
	NamedConstructor,  // line #4705:     isMultiLine = json['isMultiLine'];
	NamedConstructor,  // line #4706:     propertyKey =
	NamedConstructor,  // line #4707:         createServiceObject(json['propertyKey'], const ['InstanceRef']);
	NamedConstructor,  // line #4708:     propertyValue =
	NamedConstructor,  // line #4709:         createServiceObject(json['propertyValue'], const ['InstanceRef']);
	NamedConstructor,  // line #4710:     typeArguments =
	NamedConstructor,  // line #4711:         createServiceObject(json['typeArguments'], const ['TypeArgumentsRef']);
	NamedConstructor,  // line #4712:     parameterIndex = json['parameterIndex'];
	NamedConstructor,  // line #4713:     targetType = createServiceObject(json['targetType'], const ['InstanceRef']);
	NamedConstructor,  // line #4714:     bound = createServiceObject(json['bound'], const ['InstanceRef']);
	NamedConstructor,  // line #4715:     portId = json['portId'];
	NamedConstructor,  // line #4716:     allocationLocation =
	NamedConstructor,  // line #4717:         createServiceObject(json['allocationLocation'], const ['InstanceRef']);
	NamedConstructor,  // line #4718:     debugName = json['debugName'];
	NamedConstructor,  // line #4719:   }
	BlankLine,         // line #4720:
	OverrideMethod,    // line #4721:   @override
	OverrideMethod,    // line #4722:   Map<String, dynamic> toJson() {
	OverrideMethod,    // line #4723:     var json = super.toJson();
	OverrideMethod,    // line #4724:     json['type'] = 'Instance';
	OverrideMethod,    // line #4725:     json.addAll({
	OverrideMethod,    // line #4726:       'kind': kind,
	OverrideMethod,    // line #4727:       'class': classRef.toJson(),
	OverrideMethod,    // line #4728:     });
	OverrideMethod,    // line #4729:     _setIfNotNull(json, 'valueAsString', valueAsString);
	OverrideMethod,    // line #4730:     _setIfNotNull(json, 'valueAsStringIsTruncated', valueAsStringIsTruncated);
	OverrideMethod,    // line #4731:     _setIfNotNull(json, 'length', length);
	OverrideMethod,    // line #4732:     _setIfNotNull(json, 'offset', offset);
	OverrideMethod,    // line #4733:     _setIfNotNull(json, 'count', count);
	OverrideMethod,    // line #4734:     _setIfNotNull(json, 'name', name);
	OverrideMethod,    // line #4735:     _setIfNotNull(json, 'typeClass', typeClass?.toJson());
	OverrideMethod,    // line #4736:     _setIfNotNull(json, 'parameterizedClass', parameterizedClass?.toJson());
	OverrideMethod,    // line #4737:     _setIfNotNull(json, 'fields', fields?.map((f) => f?.toJson())?.toList());
	OverrideMethod,    // line #4738:     _setIfNotNull(
	OverrideMethod,    // line #4739:         json, 'elements', elements?.map((f) => f?.toJson())?.toList());
	OverrideMethod,    // line #4740:     _setIfNotNull(
	OverrideMethod,    // line #4741:         json, 'associations', associations?.map((f) => f?.toJson())?.toList());
	OverrideMethod,    // line #4742:     _setIfNotNull(json, 'bytes', bytes);
	OverrideMethod,    // line #4743:     _setIfNotNull(json, 'mirrorReferent', mirrorReferent?.toJson());
	OverrideMethod,    // line #4744:     _setIfNotNull(json, 'pattern', pattern?.toJson());
	OverrideMethod,    // line #4745:     _setIfNotNull(json, 'closureFunction', closureFunction?.toJson());
	OverrideMethod,    // line #4746:     _setIfNotNull(json, 'closureContext', closureContext?.toJson());
	OverrideMethod,    // line #4747:     _setIfNotNull(json, 'isCaseSensitive', isCaseSensitive);
	OverrideMethod,    // line #4748:     _setIfNotNull(json, 'isMultiLine', isMultiLine);
	OverrideMethod,    // line #4749:     _setIfNotNull(json, 'propertyKey', propertyKey?.toJson());
	OverrideMethod,    // line #4750:     _setIfNotNull(json, 'propertyValue', propertyValue?.toJson());
	OverrideMethod,    // line #4751:     _setIfNotNull(json, 'typeArguments', typeArguments?.toJson());
	OverrideMethod,    // line #4752:     _setIfNotNull(json, 'parameterIndex', parameterIndex);
	OverrideMethod,    // line #4753:     _setIfNotNull(json, 'targetType', targetType?.toJson());
	OverrideMethod,    // line #4754:     _setIfNotNull(json, 'bound', bound?.toJson());
	OverrideMethod,    // line #4755:     _setIfNotNull(json, 'portId', portId);
	OverrideMethod,    // line #4756:     _setIfNotNull(json, 'allocationLocation', allocationLocation?.toJson());
	OverrideMethod,    // line #4757:     _setIfNotNull(json, 'debugName', debugName);
	OverrideMethod,    // line #4758:     return json;
	OverrideMethod,    // line #4759:   }
	BlankLine,         // line #4760:
	OtherMethod,       // line #4761:   int get hashCode => id.hashCode;
	BlankLine,         // line #4762:
	OtherMethod,       // line #4763:   operator ==(other) => other is Instance && id == other.id;
	BlankLine,         // line #4764:
	OtherMethod,       // line #4765:   String toString() => '[Instance ' //
	OtherMethod,       // line #4766:       'type: ${type}, id: ${id}, kind: ${kind}, classRef: ${classRef}]';
	BlankLine,         // line #4767:
}

var wantVMServiceClass47 = []EntityType{
	Unknown,          // line #4770: {
	OtherMethod,      // line #4771:   static IsolateRef parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #4772:       json == null ? null : IsolateRef._fromJson(json);
	BlankLine,        // line #4773:
	InstanceVariable, // line #4774:   /// The id which is passed to the getIsolate RPC to load this isolate.
	InstanceVariable, // line #4775:   String id;
	BlankLine,        // line #4776:
	InstanceVariable, // line #4777:   /// A numeric id for this isolate, represented as a string. Unique.
	InstanceVariable, // line #4778:   String number;
	BlankLine,        // line #4779:
	InstanceVariable, // line #4780:   /// A name identifying this isolate. Not guaranteed to be unique.
	InstanceVariable, // line #4781:   String name;
	BlankLine,        // line #4782:
	InstanceVariable, // line #4783:   /// Specifies whether the isolate was spawned by the VM or embedder for
	InstanceVariable, // line #4784:   /// internal use. If `false`, this isolate is likely running user code.
	InstanceVariable, // line #4785:   bool isSystemIsolate;
	BlankLine,        // line #4786:
	MainConstructor,  // line #4787:   IsolateRef({
	MainConstructor,  // line #4788:     @required this.id,
	MainConstructor,  // line #4789:     @required this.number,
	MainConstructor,  // line #4790:     @required this.name,
	MainConstructor,  // line #4791:     @required this.isSystemIsolate,
	MainConstructor,  // line #4792:   });
	BlankLine,        // line #4793:
	NamedConstructor, // line #4794:   IsolateRef._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #4795:     id = json['id'];
	NamedConstructor, // line #4796:     number = json['number'];
	NamedConstructor, // line #4797:     name = json['name'];
	NamedConstructor, // line #4798:     isSystemIsolate = json['isSystemIsolate'];
	NamedConstructor, // line #4799:   }
	BlankLine,        // line #4800:
	OverrideMethod,   // line #4801:   @override
	OverrideMethod,   // line #4802:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #4803:     var json = <String, dynamic>{};
	OverrideMethod,   // line #4804:     json['type'] = '@Isolate';
	OverrideMethod,   // line #4805:     json.addAll({
	OverrideMethod,   // line #4806:       'id': id,
	OverrideMethod,   // line #4807:       'number': number,
	OverrideMethod,   // line #4808:       'name': name,
	OverrideMethod,   // line #4809:       'isSystemIsolate': isSystemIsolate,
	OverrideMethod,   // line #4810:     });
	OverrideMethod,   // line #4811:     return json;
	OverrideMethod,   // line #4812:   }
	BlankLine,        // line #4813:
	OtherMethod,      // line #4814:   int get hashCode => id.hashCode;
	BlankLine,        // line #4815:
	OtherMethod,      // line #4816:   operator ==(other) => other is IsolateRef && id == other.id;
	BlankLine,        // line #4817:
	OtherMethod,      // line #4818:   String toString() => '[IsolateRef ' //
	OtherMethod,      // line #4819:       'type: ${type}, id: ${id}, number: ${number}, name: ${name}, ' //
	OtherMethod,      // line #4820:       'isSystemIsolate: ${isSystemIsolate}]';
	BlankLine,        // line #4821:
}

var wantVMServiceClass48 = []EntityType{
	Unknown,           // line #4824: {
	OtherMethod,       // line #4825:   static Isolate parse(Map<String, dynamic> json) =>
	OtherMethod,       // line #4826:       json == null ? null : Isolate._fromJson(json);
	BlankLine,         // line #4827:
	InstanceVariable,  // line #4828:   /// The id which is passed to the getIsolate RPC to reload this isolate.
	InstanceVariable,  // line #4829:   String id;
	BlankLine,         // line #4830:
	InstanceVariable,  // line #4831:   /// A numeric id for this isolate, represented as a string. Unique.
	InstanceVariable,  // line #4832:   String number;
	BlankLine,         // line #4833:
	InstanceVariable,  // line #4834:   /// A name identifying this isolate. Not guaranteed to be unique.
	InstanceVariable,  // line #4835:   String name;
	BlankLine,         // line #4836:
	InstanceVariable,  // line #4837:   /// Specifies whether the isolate was spawned by the VM or embedder for
	InstanceVariable,  // line #4838:   /// internal use. If `false`, this isolate is likely running user code.
	InstanceVariable,  // line #4839:   bool isSystemIsolate;
	BlankLine,         // line #4840:
	InstanceVariable,  // line #4841:   /// The list of isolate flags provided to this isolate. See Dart_IsolateFlags
	InstanceVariable,  // line #4842:   /// in dart_api.h for the list of accepted isolate flags.
	InstanceVariable,  // line #4843:   List<IsolateFlag> isolateFlags;
	BlankLine,         // line #4844:
	InstanceVariable,  // line #4845:   /// The time that the VM started in milliseconds since the epoch.
	InstanceVariable,  // line #4846:   ///
	InstanceVariable,  // line #4847:   /// Suitable to pass to DateTime.fromMillisecondsSinceEpoch.
	InstanceVariable,  // line #4848:   int startTime;
	BlankLine,         // line #4849:
	InstanceVariable,  // line #4850:   /// Is the isolate in a runnable state?
	InstanceVariable,  // line #4851:   bool runnable;
	BlankLine,         // line #4852:
	InstanceVariable,  // line #4853:   /// The number of live ports for this isolate.
	InstanceVariable,  // line #4854:   int livePorts;
	BlankLine,         // line #4855:
	InstanceVariable,  // line #4856:   /// Will this isolate pause when exiting?
	InstanceVariable,  // line #4857:   bool pauseOnExit;
	BlankLine,         // line #4858:
	InstanceVariable,  // line #4859:   /// The last pause event delivered to the isolate. If the isolate is running,
	InstanceVariable,  // line #4860:   /// this will be a resume event.
	InstanceVariable,  // line #4861:   Event pauseEvent;
	BlankLine,         // line #4862:
	InstanceVariable,  // line #4863:   /// The root library for this isolate.
	InstanceVariable,  // line #4864:   ///
	InstanceVariable,  // line #4865:   /// Guaranteed to be initialized when the IsolateRunnable event fires.
	InstanceVariable,  // line #4866:   @optional
	InstanceVariable,  // line #4867:   LibraryRef rootLib;
	BlankLine,         // line #4868:
	InstanceVariable,  // line #4869:   /// A list of all libraries for this isolate.
	InstanceVariable,  // line #4870:   ///
	InstanceVariable,  // line #4871:   /// Guaranteed to be initialized when the IsolateRunnable event fires.
	InstanceVariable,  // line #4872:   List<LibraryRef> libraries;
	BlankLine,         // line #4873:
	InstanceVariable,  // line #4874:   /// A list of all breakpoints for this isolate.
	InstanceVariable,  // line #4875:   List<Breakpoint> breakpoints;
	BlankLine,         // line #4876:
	InstanceVariable,  // line #4877:   /// The error that is causing this isolate to exit, if applicable.
	InstanceVariable,  // line #4878:   @optional
	InstanceVariable,  // line #4879:   Error error;
	BlankLine,         // line #4880:
	SingleLineComment, // line #4881:   /// The current pause on exception mode for this isolate.
	MultiLineComment,  // line #4882:   /*ExceptionPauseMode*/ String exceptionPauseMode;
	BlankLine,         // line #4883:
	InstanceVariable,  // line #4884:   /// The list of service extension RPCs that are registered for this isolate,
	InstanceVariable,  // line #4885:   /// if any.
	InstanceVariable,  // line #4886:   @optional
	InstanceVariable,  // line #4887:   List<String> extensionRPCs;
	BlankLine,         // line #4888:
	MainConstructor,   // line #4889:   Isolate({
	MainConstructor,   // line #4890:     @required this.id,
	MainConstructor,   // line #4891:     @required this.number,
	MainConstructor,   // line #4892:     @required this.name,
	MainConstructor,   // line #4893:     @required this.isSystemIsolate,
	MainConstructor,   // line #4894:     @required this.isolateFlags,
	MainConstructor,   // line #4895:     @required this.startTime,
	MainConstructor,   // line #4896:     @required this.runnable,
	MainConstructor,   // line #4897:     @required this.livePorts,
	MainConstructor,   // line #4898:     @required this.pauseOnExit,
	MainConstructor,   // line #4899:     @required this.pauseEvent,
	MainConstructor,   // line #4900:     @required this.libraries,
	MainConstructor,   // line #4901:     @required this.breakpoints,
	MainConstructor,   // line #4902:     @required this.exceptionPauseMode,
	MainConstructor,   // line #4903:     this.rootLib,
	MainConstructor,   // line #4904:     this.error,
	MainConstructor,   // line #4905:     this.extensionRPCs,
	MainConstructor,   // line #4906:   });
	BlankLine,         // line #4907:
	NamedConstructor,  // line #4908:   Isolate._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor,  // line #4909:     id = json['id'];
	NamedConstructor,  // line #4910:     number = json['number'];
	NamedConstructor,  // line #4911:     name = json['name'];
	NamedConstructor,  // line #4912:     isSystemIsolate = json['isSystemIsolate'];
	NamedConstructor,  // line #4913:     isolateFlags = List<IsolateFlag>.from(
	NamedConstructor,  // line #4914:         createServiceObject(json['isolateFlags'], const ['IsolateFlag']) ?? []);
	NamedConstructor,  // line #4915:     startTime = json['startTime'];
	NamedConstructor,  // line #4916:     runnable = json['runnable'];
	NamedConstructor,  // line #4917:     livePorts = json['livePorts'];
	NamedConstructor,  // line #4918:     pauseOnExit = json['pauseOnExit'];
	NamedConstructor,  // line #4919:     pauseEvent = createServiceObject(json['pauseEvent'], const ['Event']);
	NamedConstructor,  // line #4920:     rootLib = createServiceObject(json['rootLib'], const ['LibraryRef']);
	NamedConstructor,  // line #4921:     libraries = List<LibraryRef>.from(
	NamedConstructor,  // line #4922:         createServiceObject(json['libraries'], const ['LibraryRef']) ?? []);
	NamedConstructor,  // line #4923:     breakpoints = List<Breakpoint>.from(
	NamedConstructor,  // line #4924:         createServiceObject(json['breakpoints'], const ['Breakpoint']) ?? []);
	NamedConstructor,  // line #4925:     error = createServiceObject(json['error'], const ['Error']);
	NamedConstructor,  // line #4926:     exceptionPauseMode = json['exceptionPauseMode'];
	NamedConstructor,  // line #4927:     extensionRPCs = json['extensionRPCs'] == null
	NamedConstructor,  // line #4928:         ? null
	NamedConstructor,  // line #4929:         : List<String>.from(json['extensionRPCs']);
	NamedConstructor,  // line #4930:   }
	BlankLine,         // line #4931:
	OverrideMethod,    // line #4932:   @override
	OverrideMethod,    // line #4933:   Map<String, dynamic> toJson() {
	OverrideMethod,    // line #4934:     var json = <String, dynamic>{};
	OverrideMethod,    // line #4935:     json['type'] = 'Isolate';
	OverrideMethod,    // line #4936:     json.addAll({
	OverrideMethod,    // line #4937:       'id': id,
	OverrideMethod,    // line #4938:       'number': number,
	OverrideMethod,    // line #4939:       'name': name,
	OverrideMethod,    // line #4940:       'isSystemIsolate': isSystemIsolate,
	OverrideMethod,    // line #4941:       'isolateFlags': isolateFlags.map((f) => f.toJson()).toList(),
	OverrideMethod,    // line #4942:       'startTime': startTime,
	OverrideMethod,    // line #4943:       'runnable': runnable,
	OverrideMethod,    // line #4944:       'livePorts': livePorts,
	OverrideMethod,    // line #4945:       'pauseOnExit': pauseOnExit,
	OverrideMethod,    // line #4946:       'pauseEvent': pauseEvent.toJson(),
	OverrideMethod,    // line #4947:       'libraries': libraries.map((f) => f.toJson()).toList(),
	OverrideMethod,    // line #4948:       'breakpoints': breakpoints.map((f) => f.toJson()).toList(),
	OverrideMethod,    // line #4949:       'exceptionPauseMode': exceptionPauseMode,
	OverrideMethod,    // line #4950:     });
	OverrideMethod,    // line #4951:     _setIfNotNull(json, 'rootLib', rootLib?.toJson());
	OverrideMethod,    // line #4952:     _setIfNotNull(json, 'error', error?.toJson());
	OverrideMethod,    // line #4953:     _setIfNotNull(
	OverrideMethod,    // line #4954:         json, 'extensionRPCs', extensionRPCs?.map((f) => f)?.toList());
	OverrideMethod,    // line #4955:     return json;
	OverrideMethod,    // line #4956:   }
	BlankLine,         // line #4957:
	OtherMethod,       // line #4958:   int get hashCode => id.hashCode;
	BlankLine,         // line #4959:
	OtherMethod,       // line #4960:   operator ==(other) => other is Isolate && id == other.id;
	BlankLine,         // line #4961:
	OtherMethod,       // line #4962:   String toString() => '[Isolate]';
	BlankLine,         // line #4963:
}

var wantVMServiceClass49 = []EntityType{
	Unknown,          // line #4966: {
	OtherMethod,      // line #4967:   static IsolateFlag parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #4968:       json == null ? null : IsolateFlag._fromJson(json);
	BlankLine,        // line #4969:
	InstanceVariable, // line #4970:   /// The name of the flag.
	InstanceVariable, // line #4971:   String name;
	BlankLine,        // line #4972:
	InstanceVariable, // line #4973:   /// The value of this flag as a string.
	InstanceVariable, // line #4974:   String valueAsString;
	BlankLine,        // line #4975:
	MainConstructor,  // line #4976:   IsolateFlag({
	MainConstructor,  // line #4977:     @required this.name,
	MainConstructor,  // line #4978:     @required this.valueAsString,
	MainConstructor,  // line #4979:   });
	BlankLine,        // line #4980:
	NamedConstructor, // line #4981:   IsolateFlag._fromJson(Map<String, dynamic> json) {
	NamedConstructor, // line #4982:     name = json['name'];
	NamedConstructor, // line #4983:     valueAsString = json['valueAsString'];
	NamedConstructor, // line #4984:   }
	BlankLine,        // line #4985:
	OtherMethod,      // line #4986:   Map<String, dynamic> toJson() {
	OtherMethod,      // line #4987:     var json = <String, dynamic>{};
	OtherMethod,      // line #4988:     json.addAll({
	OtherMethod,      // line #4989:       'name': name,
	OtherMethod,      // line #4990:       'valueAsString': valueAsString,
	OtherMethod,      // line #4991:     });
	OtherMethod,      // line #4992:     return json;
	OtherMethod,      // line #4993:   }
	BlankLine,        // line #4994:
	OtherMethod,      // line #4995:   String toString() =>
	OtherMethod,      // line #4996:       '[IsolateFlag name: ${name}, valueAsString: ${valueAsString}]';
	BlankLine,        // line #4997:
}

var wantVMServiceClass50 = []EntityType{
	Unknown,          // line #5000: {
	OtherMethod,      // line #5001:   static IsolateGroupRef parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #5002:       json == null ? null : IsolateGroupRef._fromJson(json);
	BlankLine,        // line #5003:
	InstanceVariable, // line #5004:   /// The id which is passed to the getIsolateGroup RPC to load this isolate
	InstanceVariable, // line #5005:   /// group.
	InstanceVariable, // line #5006:   String id;
	BlankLine,        // line #5007:
	InstanceVariable, // line #5008:   /// A numeric id for this isolate group, represented as a string. Unique.
	InstanceVariable, // line #5009:   String number;
	BlankLine,        // line #5010:
	InstanceVariable, // line #5011:   /// A name identifying this isolate group. Not guaranteed to be unique.
	InstanceVariable, // line #5012:   String name;
	BlankLine,        // line #5013:
	InstanceVariable, // line #5014:   /// Specifies whether the isolate group was spawned by the VM or embedder for
	InstanceVariable, // line #5015:   /// internal use. If `false`, this isolate group is likely running user code.
	InstanceVariable, // line #5016:   bool isSystemIsolateGroup;
	BlankLine,        // line #5017:
	MainConstructor,  // line #5018:   IsolateGroupRef({
	MainConstructor,  // line #5019:     @required this.id,
	MainConstructor,  // line #5020:     @required this.number,
	MainConstructor,  // line #5021:     @required this.name,
	MainConstructor,  // line #5022:     @required this.isSystemIsolateGroup,
	MainConstructor,  // line #5023:   });
	BlankLine,        // line #5024:
	NamedConstructor, // line #5025:   IsolateGroupRef._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #5026:     id = json['id'];
	NamedConstructor, // line #5027:     number = json['number'];
	NamedConstructor, // line #5028:     name = json['name'];
	NamedConstructor, // line #5029:     isSystemIsolateGroup = json['isSystemIsolateGroup'];
	NamedConstructor, // line #5030:   }
	BlankLine,        // line #5031:
	OverrideMethod,   // line #5032:   @override
	OverrideMethod,   // line #5033:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #5034:     var json = <String, dynamic>{};
	OverrideMethod,   // line #5035:     json['type'] = '@IsolateGroup';
	OverrideMethod,   // line #5036:     json.addAll({
	OverrideMethod,   // line #5037:       'id': id,
	OverrideMethod,   // line #5038:       'number': number,
	OverrideMethod,   // line #5039:       'name': name,
	OverrideMethod,   // line #5040:       'isSystemIsolateGroup': isSystemIsolateGroup,
	OverrideMethod,   // line #5041:     });
	OverrideMethod,   // line #5042:     return json;
	OverrideMethod,   // line #5043:   }
	BlankLine,        // line #5044:
	OtherMethod,      // line #5045:   int get hashCode => id.hashCode;
	BlankLine,        // line #5046:
	OtherMethod,      // line #5047:   operator ==(other) => other is IsolateGroupRef && id == other.id;
	BlankLine,        // line #5048:
	OtherMethod,      // line #5049:   String toString() => '[IsolateGroupRef ' //
	OtherMethod,      // line #5050:       'type: ${type}, id: ${id}, number: ${number}, name: ${name}, ' //
	OtherMethod,      // line #5051:       'isSystemIsolateGroup: ${isSystemIsolateGroup}]';
	BlankLine,        // line #5052:
}

var wantVMServiceClass51 = []EntityType{
	Unknown,          // line #5055: {
	OtherMethod,      // line #5056:   static IsolateGroup parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #5057:       json == null ? null : IsolateGroup._fromJson(json);
	BlankLine,        // line #5058:
	InstanceVariable, // line #5059:   /// The id which is passed to the getIsolate RPC to reload this isolate.
	InstanceVariable, // line #5060:   String id;
	BlankLine,        // line #5061:
	InstanceVariable, // line #5062:   /// A numeric id for this isolate, represented as a string. Unique.
	InstanceVariable, // line #5063:   String number;
	BlankLine,        // line #5064:
	InstanceVariable, // line #5065:   /// A name identifying this isolate. Not guaranteed to be unique.
	InstanceVariable, // line #5066:   String name;
	BlankLine,        // line #5067:
	InstanceVariable, // line #5068:   /// Specifies whether the isolate group was spawned by the VM or embedder for
	InstanceVariable, // line #5069:   /// internal use. If `false`, this isolate group is likely running user code.
	InstanceVariable, // line #5070:   bool isSystemIsolateGroup;
	BlankLine,        // line #5071:
	InstanceVariable, // line #5072:   /// A list of all isolates in this isolate group.
	InstanceVariable, // line #5073:   List<IsolateRef> isolates;
	BlankLine,        // line #5074:
	MainConstructor,  // line #5075:   IsolateGroup({
	MainConstructor,  // line #5076:     @required this.id,
	MainConstructor,  // line #5077:     @required this.number,
	MainConstructor,  // line #5078:     @required this.name,
	MainConstructor,  // line #5079:     @required this.isSystemIsolateGroup,
	MainConstructor,  // line #5080:     @required this.isolates,
	MainConstructor,  // line #5081:   });
	BlankLine,        // line #5082:
	NamedConstructor, // line #5083:   IsolateGroup._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #5084:     id = json['id'];
	NamedConstructor, // line #5085:     number = json['number'];
	NamedConstructor, // line #5086:     name = json['name'];
	NamedConstructor, // line #5087:     isSystemIsolateGroup = json['isSystemIsolateGroup'];
	NamedConstructor, // line #5088:     isolates = List<IsolateRef>.from(
	NamedConstructor, // line #5089:         createServiceObject(json['isolates'], const ['IsolateRef']) ?? []);
	NamedConstructor, // line #5090:   }
	BlankLine,        // line #5091:
	OverrideMethod,   // line #5092:   @override
	OverrideMethod,   // line #5093:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #5094:     var json = <String, dynamic>{};
	OverrideMethod,   // line #5095:     json['type'] = 'IsolateGroup';
	OverrideMethod,   // line #5096:     json.addAll({
	OverrideMethod,   // line #5097:       'id': id,
	OverrideMethod,   // line #5098:       'number': number,
	OverrideMethod,   // line #5099:       'name': name,
	OverrideMethod,   // line #5100:       'isSystemIsolateGroup': isSystemIsolateGroup,
	OverrideMethod,   // line #5101:       'isolates': isolates.map((f) => f.toJson()).toList(),
	OverrideMethod,   // line #5102:     });
	OverrideMethod,   // line #5103:     return json;
	OverrideMethod,   // line #5104:   }
	BlankLine,        // line #5105:
	OtherMethod,      // line #5106:   int get hashCode => id.hashCode;
	BlankLine,        // line #5107:
	OtherMethod,      // line #5108:   operator ==(other) => other is IsolateGroup && id == other.id;
	BlankLine,        // line #5109:
	OtherMethod,      // line #5110:   String toString() => '[IsolateGroup ' //
	OtherMethod,      // line #5111:       'type: ${type}, id: ${id}, number: ${number}, name: ${name}, ' //
	OtherMethod,      // line #5112:       'isSystemIsolateGroup: ${isSystemIsolateGroup}, isolates: ${isolates}]';
	BlankLine,        // line #5113:
}

var wantVMServiceClass52 = []EntityType{
	Unknown,          // line #5116: {
	OtherMethod,      // line #5117:   static InboundReferences parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #5118:       json == null ? null : InboundReferences._fromJson(json);
	BlankLine,        // line #5119:
	InstanceVariable, // line #5120:   /// An array of inbound references to an object.
	InstanceVariable, // line #5121:   List<InboundReference> references;
	BlankLine,        // line #5122:
	MainConstructor,  // line #5123:   InboundReferences({
	MainConstructor,  // line #5124:     @required this.references,
	MainConstructor,  // line #5125:   });
	BlankLine,        // line #5126:
	NamedConstructor, // line #5127:   InboundReferences._fromJson(Map<String, dynamic> json)
	NamedConstructor, // line #5128:       : super._fromJson(json) {
	NamedConstructor, // line #5129:     references = List<InboundReference>.from(
	NamedConstructor, // line #5130:         createServiceObject(json['references'], const ['InboundReference']) ??
	NamedConstructor, // line #5131:             []);
	NamedConstructor, // line #5132:   }
	BlankLine,        // line #5133:
	OverrideMethod,   // line #5134:   @override
	OverrideMethod,   // line #5135:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #5136:     var json = <String, dynamic>{};
	OverrideMethod,   // line #5137:     json['type'] = 'InboundReferences';
	OverrideMethod,   // line #5138:     json.addAll({
	OverrideMethod,   // line #5139:       'references': references.map((f) => f.toJson()).toList(),
	OverrideMethod,   // line #5140:     });
	OverrideMethod,   // line #5141:     return json;
	OverrideMethod,   // line #5142:   }
	BlankLine,        // line #5143:
	OtherMethod,      // line #5144:   String toString() =>
	OtherMethod,      // line #5145:       '[InboundReferences type: ${type}, references: ${references}]';
	BlankLine,        // line #5146:
}

var wantVMServiceClass53 = []EntityType{
	Unknown,          // line #5149: {
	OtherMethod,      // line #5150:   static InboundReference parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #5151:       json == null ? null : InboundReference._fromJson(json);
	BlankLine,        // line #5152:
	InstanceVariable, // line #5153:   /// The object holding the inbound reference.
	InstanceVariable, // line #5154:   ObjRef source;
	BlankLine,        // line #5155:
	InstanceVariable, // line #5156:   /// If source is a List, parentListIndex is the index of the inbound
	InstanceVariable, // line #5157:   /// reference.
	InstanceVariable, // line #5158:   @optional
	InstanceVariable, // line #5159:   int parentListIndex;
	BlankLine,        // line #5160:
	InstanceVariable, // line #5161:   /// If source is a field of an object, parentField is the field containing the
	InstanceVariable, // line #5162:   /// inbound reference.
	InstanceVariable, // line #5163:   @optional
	InstanceVariable, // line #5164:   FieldRef parentField;
	BlankLine,        // line #5165:
	MainConstructor,  // line #5166:   InboundReference({
	MainConstructor,  // line #5167:     @required this.source,
	MainConstructor,  // line #5168:     this.parentListIndex,
	MainConstructor,  // line #5169:     this.parentField,
	MainConstructor,  // line #5170:   });
	BlankLine,        // line #5171:
	NamedConstructor, // line #5172:   InboundReference._fromJson(Map<String, dynamic> json) {
	NamedConstructor, // line #5173:     source = createServiceObject(json['source'], const ['ObjRef']);
	NamedConstructor, // line #5174:     parentListIndex = json['parentListIndex'];
	NamedConstructor, // line #5175:     parentField = createServiceObject(json['parentField'], const ['FieldRef']);
	NamedConstructor, // line #5176:   }
	BlankLine,        // line #5177:
	OtherMethod,      // line #5178:   Map<String, dynamic> toJson() {
	OtherMethod,      // line #5179:     var json = <String, dynamic>{};
	OtherMethod,      // line #5180:     json.addAll({
	OtherMethod,      // line #5181:       'source': source.toJson(),
	OtherMethod,      // line #5182:     });
	OtherMethod,      // line #5183:     _setIfNotNull(json, 'parentListIndex', parentListIndex);
	OtherMethod,      // line #5184:     _setIfNotNull(json, 'parentField', parentField?.toJson());
	OtherMethod,      // line #5185:     return json;
	OtherMethod,      // line #5186:   }
	BlankLine,        // line #5187:
	OtherMethod,      // line #5188:   String toString() => '[InboundReference source: ${source}]';
	BlankLine,        // line #5189:
}

var wantVMServiceClass54 = []EntityType{
	Unknown,          // line #5192: {
	OtherMethod,      // line #5193:   static InstanceSet parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #5194:       json == null ? null : InstanceSet._fromJson(json);
	BlankLine,        // line #5195:
	InstanceVariable, // line #5196:   /// The number of instances of the requested type currently allocated.
	InstanceVariable, // line #5197:   int totalCount;
	BlankLine,        // line #5198:
	InstanceVariable, // line #5199:   /// An array of instances of the requested type.
	InstanceVariable, // line #5200:   List<ObjRef> instances;
	BlankLine,        // line #5201:
	MainConstructor,  // line #5202:   InstanceSet({
	MainConstructor,  // line #5203:     @required this.totalCount,
	MainConstructor,  // line #5204:     @required this.instances,
	MainConstructor,  // line #5205:   });
	BlankLine,        // line #5206:
	NamedConstructor, // line #5207:   InstanceSet._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #5208:     totalCount = json['totalCount'];
	NamedConstructor, // line #5209:     instances = List<ObjRef>.from(createServiceObject(
	NamedConstructor, // line #5210:         json['instances'] ?? json['samples'], const ['ObjRef']));
	NamedConstructor, // line #5211:   }
	BlankLine,        // line #5212:
	OverrideMethod,   // line #5213:   @override
	OverrideMethod,   // line #5214:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #5215:     var json = <String, dynamic>{};
	OverrideMethod,   // line #5216:     json['type'] = 'InstanceSet';
	OverrideMethod,   // line #5217:     json.addAll({
	OverrideMethod,   // line #5218:       'totalCount': totalCount,
	OverrideMethod,   // line #5219:       'instances': instances.map((f) => f.toJson()).toList(),
	OverrideMethod,   // line #5220:     });
	OverrideMethod,   // line #5221:     return json;
	OverrideMethod,   // line #5222:   }
	BlankLine,        // line #5223:
	OtherMethod,      // line #5224:   String toString() => '[InstanceSet ' //
	OtherMethod,      // line #5225:       'type: ${type}, totalCount: ${totalCount}, instances: ${instances}]';
	BlankLine,        // line #5226:
}

var wantVMServiceClass55 = []EntityType{
	Unknown,          // line #5229: {
	OtherMethod,      // line #5230:   static LibraryRef parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #5231:       json == null ? null : LibraryRef._fromJson(json);
	BlankLine,        // line #5232:
	InstanceVariable, // line #5233:   /// The name of this library.
	InstanceVariable, // line #5234:   String name;
	BlankLine,        // line #5235:
	InstanceVariable, // line #5236:   /// The uri of this library.
	InstanceVariable, // line #5237:   String uri;
	BlankLine,        // line #5238:
	MainConstructor,  // line #5239:   LibraryRef({
	MainConstructor,  // line #5240:     @required this.name,
	MainConstructor,  // line #5241:     @required this.uri,
	MainConstructor,  // line #5242:     @required String id,
	MainConstructor,  // line #5243:   }) : super(id: id);
	BlankLine,        // line #5244:
	NamedConstructor, // line #5245:   LibraryRef._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #5246:     name = json['name'];
	NamedConstructor, // line #5247:     uri = json['uri'];
	NamedConstructor, // line #5248:   }
	BlankLine,        // line #5249:
	OverrideMethod,   // line #5250:   @override
	OverrideMethod,   // line #5251:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #5252:     var json = super.toJson();
	OverrideMethod,   // line #5253:     json['type'] = '@Library';
	OverrideMethod,   // line #5254:     json.addAll({
	OverrideMethod,   // line #5255:       'name': name,
	OverrideMethod,   // line #5256:       'uri': uri,
	OverrideMethod,   // line #5257:     });
	OverrideMethod,   // line #5258:     return json;
	OverrideMethod,   // line #5259:   }
	BlankLine,        // line #5260:
	OtherMethod,      // line #5261:   int get hashCode => id.hashCode;
	BlankLine,        // line #5262:
	OtherMethod,      // line #5263:   operator ==(other) => other is LibraryRef && id == other.id;
	BlankLine,        // line #5264:
	OtherMethod,      // line #5265:   String toString() =>
	OtherMethod,      // line #5266:       '[LibraryRef type: ${type}, id: ${id}, name: ${name}, uri: ${uri}]';
	BlankLine,        // line #5267:
}

var wantVMServiceClass56 = []EntityType{
	Unknown,          // line #5272: {
	OtherMethod,      // line #5273:   static Library parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #5274:       json == null ? null : Library._fromJson(json);
	BlankLine,        // line #5275:
	InstanceVariable, // line #5276:   /// The name of this library.
	InstanceVariable, // line #5277:   String name;
	BlankLine,        // line #5278:
	InstanceVariable, // line #5279:   /// The uri of this library.
	InstanceVariable, // line #5280:   String uri;
	BlankLine,        // line #5281:
	InstanceVariable, // line #5282:   /// Is this library debuggable? Default true.
	InstanceVariable, // line #5283:   bool debuggable;
	BlankLine,        // line #5284:
	InstanceVariable, // line #5285:   /// A list of the imports for this library.
	InstanceVariable, // line #5286:   List<LibraryDependency> dependencies;
	BlankLine,        // line #5287:
	InstanceVariable, // line #5288:   /// A list of the scripts which constitute this library.
	InstanceVariable, // line #5289:   List<ScriptRef> scripts;
	BlankLine,        // line #5290:
	InstanceVariable, // line #5291:   /// A list of the top-level variables in this library.
	InstanceVariable, // line #5292:   List<FieldRef> variables;
	BlankLine,        // line #5293:
	InstanceVariable, // line #5294:   /// A list of the top-level functions in this library.
	InstanceVariable, // line #5295:   List<FuncRef> functions;
	BlankLine,        // line #5296:
	InstanceVariable, // line #5297:   /// A list of all classes in this library.
	InstanceVariable, // line #5298:   List<ClassRef> classes;
	BlankLine,        // line #5299:
	MainConstructor,  // line #5300:   Library({
	MainConstructor,  // line #5301:     @required this.name,
	MainConstructor,  // line #5302:     @required this.uri,
	MainConstructor,  // line #5303:     @required this.debuggable,
	MainConstructor,  // line #5304:     @required this.dependencies,
	MainConstructor,  // line #5305:     @required this.scripts,
	MainConstructor,  // line #5306:     @required this.variables,
	MainConstructor,  // line #5307:     @required this.functions,
	MainConstructor,  // line #5308:     @required this.classes,
	MainConstructor,  // line #5309:     @required String id,
	MainConstructor,  // line #5310:   }) : super(id: id);
	BlankLine,        // line #5311:
	NamedConstructor, // line #5312:   Library._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #5313:     name = json['name'];
	NamedConstructor, // line #5314:     uri = json['uri'];
	NamedConstructor, // line #5315:     debuggable = json['debuggable'];
	NamedConstructor, // line #5316:     dependencies = List<LibraryDependency>.from(
	NamedConstructor, // line #5317:         _createSpecificObject(json['dependencies'], LibraryDependency.parse));
	NamedConstructor, // line #5318:     scripts = List<ScriptRef>.from(
	NamedConstructor, // line #5319:         createServiceObject(json['scripts'], const ['ScriptRef']) ?? []);
	NamedConstructor, // line #5320:     variables = List<FieldRef>.from(
	NamedConstructor, // line #5321:         createServiceObject(json['variables'], const ['FieldRef']) ?? []);
	NamedConstructor, // line #5322:     functions = List<FuncRef>.from(
	NamedConstructor, // line #5323:         createServiceObject(json['functions'], const ['FuncRef']) ?? []);
	NamedConstructor, // line #5324:     classes = List<ClassRef>.from(
	NamedConstructor, // line #5325:         createServiceObject(json['classes'], const ['ClassRef']) ?? []);
	NamedConstructor, // line #5326:   }
	BlankLine,        // line #5327:
	OverrideMethod,   // line #5328:   @override
	OverrideMethod,   // line #5329:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #5330:     var json = super.toJson();
	OverrideMethod,   // line #5331:     json['type'] = 'Library';
	OverrideMethod,   // line #5332:     json.addAll({
	OverrideMethod,   // line #5333:       'name': name,
	OverrideMethod,   // line #5334:       'uri': uri,
	OverrideMethod,   // line #5335:       'debuggable': debuggable,
	OverrideMethod,   // line #5336:       'dependencies': dependencies.map((f) => f.toJson()).toList(),
	OverrideMethod,   // line #5337:       'scripts': scripts.map((f) => f.toJson()).toList(),
	OverrideMethod,   // line #5338:       'variables': variables.map((f) => f.toJson()).toList(),
	OverrideMethod,   // line #5339:       'functions': functions.map((f) => f.toJson()).toList(),
	OverrideMethod,   // line #5340:       'classes': classes.map((f) => f.toJson()).toList(),
	OverrideMethod,   // line #5341:     });
	OverrideMethod,   // line #5342:     return json;
	OverrideMethod,   // line #5343:   }
	BlankLine,        // line #5344:
	OtherMethod,      // line #5345:   int get hashCode => id.hashCode;
	BlankLine,        // line #5346:
	OtherMethod,      // line #5347:   operator ==(other) => other is Library && id == other.id;
	BlankLine,        // line #5348:
	OtherMethod,      // line #5349:   String toString() => '[Library]';
	BlankLine,        // line #5350:
}

var wantVMServiceClass57 = []EntityType{
	Unknown,          // line #5353: {
	OtherMethod,      // line #5354:   static LibraryDependency parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #5355:       json == null ? null : LibraryDependency._fromJson(json);
	BlankLine,        // line #5356:
	InstanceVariable, // line #5357:   /// Is this dependency an import (rather than an export)?
	InstanceVariable, // line #5358:   bool isImport;
	BlankLine,        // line #5359:
	InstanceVariable, // line #5360:   /// Is this dependency deferred?
	InstanceVariable, // line #5361:   bool isDeferred;
	BlankLine,        // line #5362:
	InstanceVariable, // line #5363:   /// The prefix of an 'as' import, or null.
	InstanceVariable, // line #5364:   String prefix;
	BlankLine,        // line #5365:
	InstanceVariable, // line #5366:   /// The library being imported or exported.
	InstanceVariable, // line #5367:   LibraryRef target;
	BlankLine,        // line #5368:
	MainConstructor,  // line #5369:   LibraryDependency({
	MainConstructor,  // line #5370:     @required this.isImport,
	MainConstructor,  // line #5371:     @required this.isDeferred,
	MainConstructor,  // line #5372:     @required this.prefix,
	MainConstructor,  // line #5373:     @required this.target,
	MainConstructor,  // line #5374:   });
	BlankLine,        // line #5375:
	NamedConstructor, // line #5376:   LibraryDependency._fromJson(Map<String, dynamic> json) {
	NamedConstructor, // line #5377:     isImport = json['isImport'];
	NamedConstructor, // line #5378:     isDeferred = json['isDeferred'];
	NamedConstructor, // line #5379:     prefix = json['prefix'];
	NamedConstructor, // line #5380:     target = createServiceObject(json['target'], const ['LibraryRef']);
	NamedConstructor, // line #5381:   }
	BlankLine,        // line #5382:
	OtherMethod,      // line #5383:   Map<String, dynamic> toJson() {
	OtherMethod,      // line #5384:     var json = <String, dynamic>{};
	OtherMethod,      // line #5385:     json.addAll({
	OtherMethod,      // line #5386:       'isImport': isImport,
	OtherMethod,      // line #5387:       'isDeferred': isDeferred,
	OtherMethod,      // line #5388:       'prefix': prefix,
	OtherMethod,      // line #5389:       'target': target.toJson(),
	OtherMethod,      // line #5390:     });
	OtherMethod,      // line #5391:     return json;
	OtherMethod,      // line #5392:   }
	BlankLine,        // line #5393:
	OtherMethod,      // line #5394:   String toString() => '[LibraryDependency ' //
	OtherMethod,      // line #5395:       'isImport: ${isImport}, isDeferred: ${isDeferred}, prefix: ${prefix}, ' //
	OtherMethod,      // line #5396:       'target: ${target}]';
	BlankLine,        // line #5397:
}

var wantVMServiceClass58 = []EntityType{
	Unknown,          // line #5399: {
	OtherMethod,      // line #5400:   static LogRecord parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #5401:       json == null ? null : LogRecord._fromJson(json);
	BlankLine,        // line #5402:
	InstanceVariable, // line #5403:   /// The log message.
	InstanceVariable, // line #5404:   InstanceRef message;
	BlankLine,        // line #5405:
	InstanceVariable, // line #5406:   /// The timestamp.
	InstanceVariable, // line #5407:   int time;
	BlankLine,        // line #5408:
	InstanceVariable, // line #5409:   /// The severity level (a value between 0 and 2000).
	InstanceVariable, // line #5410:   ///
	InstanceVariable, // line #5411:   /// See the package:logging `Level` class for an overview of the possible
	InstanceVariable, // line #5412:   /// values.
	InstanceVariable, // line #5413:   int level;
	BlankLine,        // line #5414:
	InstanceVariable, // line #5415:   /// A monotonically increasing sequence number.
	InstanceVariable, // line #5416:   int sequenceNumber;
	BlankLine,        // line #5417:
	InstanceVariable, // line #5418:   /// The name of the source of the log message.
	InstanceVariable, // line #5419:   InstanceRef loggerName;
	BlankLine,        // line #5420:
	InstanceVariable, // line #5421:   /// The zone where the log was emitted.
	InstanceVariable, // line #5422:   InstanceRef zone;
	BlankLine,        // line #5423:
	InstanceVariable, // line #5424:   /// An error object associated with this log event.
	InstanceVariable, // line #5425:   InstanceRef error;
	BlankLine,        // line #5426:
	InstanceVariable, // line #5427:   /// A stack trace associated with this log event.
	InstanceVariable, // line #5428:   InstanceRef stackTrace;
	BlankLine,        // line #5429:
	MainConstructor,  // line #5430:   LogRecord({
	MainConstructor,  // line #5431:     @required this.message,
	MainConstructor,  // line #5432:     @required this.time,
	MainConstructor,  // line #5433:     @required this.level,
	MainConstructor,  // line #5434:     @required this.sequenceNumber,
	MainConstructor,  // line #5435:     @required this.loggerName,
	MainConstructor,  // line #5436:     @required this.zone,
	MainConstructor,  // line #5437:     @required this.error,
	MainConstructor,  // line #5438:     @required this.stackTrace,
	MainConstructor,  // line #5439:   });
	BlankLine,        // line #5440:
	NamedConstructor, // line #5441:   LogRecord._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #5442:     message = createServiceObject(json['message'], const ['InstanceRef']);
	NamedConstructor, // line #5443:     time = json['time'];
	NamedConstructor, // line #5444:     level = json['level'];
	NamedConstructor, // line #5445:     sequenceNumber = json['sequenceNumber'];
	NamedConstructor, // line #5446:     loggerName = createServiceObject(json['loggerName'], const ['InstanceRef']);
	NamedConstructor, // line #5447:     zone = createServiceObject(json['zone'], const ['InstanceRef']);
	NamedConstructor, // line #5448:     error = createServiceObject(json['error'], const ['InstanceRef']);
	NamedConstructor, // line #5449:     stackTrace = createServiceObject(json['stackTrace'], const ['InstanceRef']);
	NamedConstructor, // line #5450:   }
	BlankLine,        // line #5451:
	OverrideMethod,   // line #5452:   @override
	OverrideMethod,   // line #5453:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #5454:     var json = <String, dynamic>{};
	OverrideMethod,   // line #5455:     json['type'] = 'LogRecord';
	OverrideMethod,   // line #5456:     json.addAll({
	OverrideMethod,   // line #5457:       'message': message.toJson(),
	OverrideMethod,   // line #5458:       'time': time,
	OverrideMethod,   // line #5459:       'level': level,
	OverrideMethod,   // line #5460:       'sequenceNumber': sequenceNumber,
	OverrideMethod,   // line #5461:       'loggerName': loggerName.toJson(),
	OverrideMethod,   // line #5462:       'zone': zone.toJson(),
	OverrideMethod,   // line #5463:       'error': error.toJson(),
	OverrideMethod,   // line #5464:       'stackTrace': stackTrace.toJson(),
	OverrideMethod,   // line #5465:     });
	OverrideMethod,   // line #5466:     return json;
	OverrideMethod,   // line #5467:   }
	BlankLine,        // line #5468:
	OtherMethod,      // line #5469:   String toString() => '[LogRecord]';
	BlankLine,        // line #5470:
}

var wantVMServiceClass59 = []EntityType{
	Unknown,          // line #5472: {
	OtherMethod,      // line #5473:   static MapAssociation parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #5474:       json == null ? null : MapAssociation._fromJson(json);
	BlankLine,        // line #5475:
	InstanceVariable, // line #5476:   /// [key] can be one of [InstanceRef] or [Sentinel].
	InstanceVariable, // line #5477:   dynamic key;
	BlankLine,        // line #5478:
	InstanceVariable, // line #5479:   /// [value] can be one of [InstanceRef] or [Sentinel].
	InstanceVariable, // line #5480:   dynamic value;
	BlankLine,        // line #5481:
	MainConstructor,  // line #5482:   MapAssociation({
	MainConstructor,  // line #5483:     @required this.key,
	MainConstructor,  // line #5484:     @required this.value,
	MainConstructor,  // line #5485:   });
	BlankLine,        // line #5486:
	NamedConstructor, // line #5487:   MapAssociation._fromJson(Map<String, dynamic> json) {
	NamedConstructor, // line #5488:     key = createServiceObject(json['key'], const ['InstanceRef', 'Sentinel']);
	NamedConstructor, // line #5489:     value =
	NamedConstructor, // line #5490:         createServiceObject(json['value'], const ['InstanceRef', 'Sentinel']);
	NamedConstructor, // line #5491:   }
	BlankLine,        // line #5492:
	OtherMethod,      // line #5493:   Map<String, dynamic> toJson() {
	OtherMethod,      // line #5494:     var json = <String, dynamic>{};
	OtherMethod,      // line #5495:     json.addAll({
	OtherMethod,      // line #5496:       'key': key.toJson(),
	OtherMethod,      // line #5497:       'value': value.toJson(),
	OtherMethod,      // line #5498:     });
	OtherMethod,      // line #5499:     return json;
	OtherMethod,      // line #5500:   }
	BlankLine,        // line #5501:
	OtherMethod,      // line #5502:   String toString() => '[MapAssociation key: ${key}, value: ${value}]';
	BlankLine,        // line #5503:
}

var wantVMServiceClass60 = []EntityType{
	Unknown,          // line #5507: {
	OtherMethod,      // line #5508:   static MemoryUsage parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #5509:       json == null ? null : MemoryUsage._fromJson(json);
	BlankLine,        // line #5510:
	InstanceVariable, // line #5511:   /// The amount of non-Dart memory that is retained by Dart objects. For
	InstanceVariable, // line #5512:   /// example, memory associated with Dart objects through APIs such as
	InstanceVariable, // line #5513:   /// Dart_NewFinalizableHandle, Dart_NewWeakPersistentHandle and
	InstanceVariable, // line #5514:   /// Dart_NewExternalTypedData.  This usage is only as accurate as the values
	InstanceVariable, // line #5515:   /// supplied to these APIs from the VM embedder or native extensions. This
	InstanceVariable, // line #5516:   /// external memory applies GC pressure, but is separate from heapUsage and
	InstanceVariable, // line #5517:   /// heapCapacity.
	InstanceVariable, // line #5518:   int externalUsage;
	BlankLine,        // line #5519:
	InstanceVariable, // line #5520:   /// The total capacity of the heap in bytes. This is the amount of memory used
	InstanceVariable, // line #5521:   /// by the Dart heap from the perspective of the operating system.
	InstanceVariable, // line #5522:   int heapCapacity;
	BlankLine,        // line #5523:
	InstanceVariable, // line #5524:   /// The current heap memory usage in bytes. Heap usage is always less than or
	InstanceVariable, // line #5525:   /// equal to the heap capacity.
	InstanceVariable, // line #5526:   int heapUsage;
	BlankLine,        // line #5527:
	MainConstructor,  // line #5528:   MemoryUsage({
	MainConstructor,  // line #5529:     @required this.externalUsage,
	MainConstructor,  // line #5530:     @required this.heapCapacity,
	MainConstructor,  // line #5531:     @required this.heapUsage,
	MainConstructor,  // line #5532:   });
	BlankLine,        // line #5533:
	NamedConstructor, // line #5534:   MemoryUsage._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #5535:     externalUsage = json['externalUsage'];
	NamedConstructor, // line #5536:     heapCapacity = json['heapCapacity'];
	NamedConstructor, // line #5537:     heapUsage = json['heapUsage'];
	NamedConstructor, // line #5538:   }
	BlankLine,        // line #5539:
	OverrideMethod,   // line #5540:   @override
	OverrideMethod,   // line #5541:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #5542:     var json = <String, dynamic>{};
	OverrideMethod,   // line #5543:     json['type'] = 'MemoryUsage';
	OverrideMethod,   // line #5544:     json.addAll({
	OverrideMethod,   // line #5545:       'externalUsage': externalUsage,
	OverrideMethod,   // line #5546:       'heapCapacity': heapCapacity,
	OverrideMethod,   // line #5547:       'heapUsage': heapUsage,
	OverrideMethod,   // line #5548:     });
	OverrideMethod,   // line #5549:     return json;
	OverrideMethod,   // line #5550:   }
	BlankLine,        // line #5551:
	OtherMethod,      // line #5552:   String toString() => '[MemoryUsage ' //
	OtherMethod,      // line #5553:       'type: ${type}, externalUsage: ${externalUsage}, heapCapacity: ${heapCapacity}, ' //
	OtherMethod,      // line #5554:       'heapUsage: ${heapUsage}]';
	BlankLine,        // line #5555:
}

var wantVMServiceClass61 = []EntityType{
	Unknown,          // line #5559: {
	OtherMethod,      // line #5560:   static Message parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #5561:       json == null ? null : Message._fromJson(json);
	BlankLine,        // line #5562:
	InstanceVariable, // line #5563:   /// The index in the isolate's message queue. The 0th message being the next
	InstanceVariable, // line #5564:   /// message to be processed.
	InstanceVariable, // line #5565:   int index;
	BlankLine,        // line #5566:
	InstanceVariable, // line #5567:   /// An advisory name describing this message.
	InstanceVariable, // line #5568:   String name;
	BlankLine,        // line #5569:
	InstanceVariable, // line #5570:   /// An instance id for the decoded message. This id can be passed to other
	InstanceVariable, // line #5571:   /// RPCs, for example, getObject or evaluate.
	InstanceVariable, // line #5572:   String messageObjectId;
	BlankLine,        // line #5573:
	InstanceVariable, // line #5574:   /// The size (bytes) of the encoded message.
	InstanceVariable, // line #5575:   int size;
	BlankLine,        // line #5576:
	InstanceVariable, // line #5577:   /// A reference to the function that will be invoked to handle this message.
	InstanceVariable, // line #5578:   @optional
	InstanceVariable, // line #5579:   FuncRef handler;
	BlankLine,        // line #5580:
	InstanceVariable, // line #5581:   /// The source location of handler.
	InstanceVariable, // line #5582:   @optional
	InstanceVariable, // line #5583:   SourceLocation location;
	BlankLine,        // line #5584:
	MainConstructor,  // line #5585:   Message({
	MainConstructor,  // line #5586:     @required this.index,
	MainConstructor,  // line #5587:     @required this.name,
	MainConstructor,  // line #5588:     @required this.messageObjectId,
	MainConstructor,  // line #5589:     @required this.size,
	MainConstructor,  // line #5590:     this.handler,
	MainConstructor,  // line #5591:     this.location,
	MainConstructor,  // line #5592:   });
	BlankLine,        // line #5593:
	NamedConstructor, // line #5594:   Message._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #5595:     index = json['index'];
	NamedConstructor, // line #5596:     name = json['name'];
	NamedConstructor, // line #5597:     messageObjectId = json['messageObjectId'];
	NamedConstructor, // line #5598:     size = json['size'];
	NamedConstructor, // line #5599:     handler = createServiceObject(json['handler'], const ['FuncRef']);
	NamedConstructor, // line #5600:     location = createServiceObject(json['location'], const ['SourceLocation']);
	NamedConstructor, // line #5601:   }
	BlankLine,        // line #5602:
	OverrideMethod,   // line #5603:   @override
	OverrideMethod,   // line #5604:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #5605:     var json = <String, dynamic>{};
	OverrideMethod,   // line #5606:     json['type'] = 'Message';
	OverrideMethod,   // line #5607:     json.addAll({
	OverrideMethod,   // line #5608:       'index': index,
	OverrideMethod,   // line #5609:       'name': name,
	OverrideMethod,   // line #5610:       'messageObjectId': messageObjectId,
	OverrideMethod,   // line #5611:       'size': size,
	OverrideMethod,   // line #5612:     });
	OverrideMethod,   // line #5613:     _setIfNotNull(json, 'handler', handler?.toJson());
	OverrideMethod,   // line #5614:     _setIfNotNull(json, 'location', location?.toJson());
	OverrideMethod,   // line #5615:     return json;
	OverrideMethod,   // line #5616:   }
	BlankLine,        // line #5617:
	OtherMethod,      // line #5618:   String toString() => '[Message ' //
	OtherMethod,      // line #5619:       'type: ${type}, index: ${index}, name: ${name}, messageObjectId: ${messageObjectId}, ' //
	OtherMethod,      // line #5620:       'size: ${size}]';
	BlankLine,        // line #5621:
}

var wantVMServiceClass62 = []EntityType{
	Unknown,          // line #5625: {
	OtherMethod,      // line #5626:   static NativeFunction parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #5627:       json == null ? null : NativeFunction._fromJson(json);
	BlankLine,        // line #5628:
	InstanceVariable, // line #5629:   /// The name of the native function this object represents.
	InstanceVariable, // line #5630:   String name;
	BlankLine,        // line #5631:
	MainConstructor,  // line #5632:   NativeFunction({
	MainConstructor,  // line #5633:     @required this.name,
	MainConstructor,  // line #5634:   });
	BlankLine,        // line #5635:
	NamedConstructor, // line #5636:   NativeFunction._fromJson(Map<String, dynamic> json) {
	NamedConstructor, // line #5637:     name = json['name'];
	NamedConstructor, // line #5638:   }
	BlankLine,        // line #5639:
	OtherMethod,      // line #5640:   Map<String, dynamic> toJson() {
	OtherMethod,      // line #5641:     var json = <String, dynamic>{};
	OtherMethod,      // line #5642:     json.addAll({
	OtherMethod,      // line #5643:       'name': name,
	OtherMethod,      // line #5644:     });
	OtherMethod,      // line #5645:     return json;
	OtherMethod,      // line #5646:   }
	BlankLine,        // line #5647:
	OtherMethod,      // line #5648:   String toString() => '[NativeFunction name: ${name}]';
	BlankLine,        // line #5649:
}

var wantVMServiceClass63 = []EntityType{
	Unknown,          // line #5652: {
	OtherMethod,      // line #5653:   static NullValRef parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #5654:       json == null ? null : NullValRef._fromJson(json);
	BlankLine,        // line #5655:
	OverrideVariable, // line #5656:   /// Always 'null'.
	OverrideVariable, // line #5657:   @override
	OverrideVariable, // line #5658:   String valueAsString;
	BlankLine,        // line #5659:
	MainConstructor,  // line #5660:   NullValRef({
	MainConstructor,  // line #5661:     @required this.valueAsString,
	MainConstructor,  // line #5662:   });
	BlankLine,        // line #5663:
	NamedConstructor, // line #5664:   NullValRef._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #5665:     valueAsString = json['valueAsString'];
	NamedConstructor, // line #5666:   }
	BlankLine,        // line #5667:
	OverrideMethod,   // line #5668:   @override
	OverrideMethod,   // line #5669:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #5670:     var json = super.toJson();
	OverrideMethod,   // line #5671:     json['type'] = '@Null';
	OverrideMethod,   // line #5672:     json.addAll({
	OverrideMethod,   // line #5673:       'valueAsString': valueAsString,
	OverrideMethod,   // line #5674:     });
	OverrideMethod,   // line #5675:     return json;
	OverrideMethod,   // line #5676:   }
	BlankLine,        // line #5677:
	OtherMethod,      // line #5678:   int get hashCode => id.hashCode;
	BlankLine,        // line #5679:
	OtherMethod,      // line #5680:   operator ==(other) => other is NullValRef && id == other.id;
	BlankLine,        // line #5681:
	OtherMethod,      // line #5682:   String toString() => '[NullValRef ' //
	OtherMethod,      // line #5683:       'type: ${type}, id: ${id}, kind: ${kind}, classRef: ${classRef}, ' //
	OtherMethod,      // line #5684:       'valueAsString: ${valueAsString}]';
	BlankLine,        // line #5685:
}

var wantVMServiceClass64 = []EntityType{
	Unknown,          // line #5688: {
	OtherMethod,      // line #5689:   static NullVal parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #5690:       json == null ? null : NullVal._fromJson(json);
	BlankLine,        // line #5691:
	OverrideVariable, // line #5692:   /// Always 'null'.
	OverrideVariable, // line #5693:   @override
	OverrideVariable, // line #5694:   String valueAsString;
	BlankLine,        // line #5695:
	MainConstructor,  // line #5696:   NullVal({
	MainConstructor,  // line #5697:     @required this.valueAsString,
	MainConstructor,  // line #5698:   });
	BlankLine,        // line #5699:
	NamedConstructor, // line #5700:   NullVal._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #5701:     valueAsString = json['valueAsString'];
	NamedConstructor, // line #5702:   }
	BlankLine,        // line #5703:
	OverrideMethod,   // line #5704:   @override
	OverrideMethod,   // line #5705:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #5706:     var json = super.toJson();
	OverrideMethod,   // line #5707:     json['type'] = 'Null';
	OverrideMethod,   // line #5708:     json.addAll({
	OverrideMethod,   // line #5709:       'valueAsString': valueAsString,
	OverrideMethod,   // line #5710:     });
	OverrideMethod,   // line #5711:     return json;
	OverrideMethod,   // line #5712:   }
	BlankLine,        // line #5713:
	OtherMethod,      // line #5714:   int get hashCode => id.hashCode;
	BlankLine,        // line #5715:
	OtherMethod,      // line #5716:   operator ==(other) => other is NullVal && id == other.id;
	BlankLine,        // line #5717:
	OtherMethod,      // line #5718:   String toString() => '[NullVal ' //
	OtherMethod,      // line #5719:       'type: ${type}, id: ${id}, kind: ${kind}, classRef: ${classRef}, ' //
	OtherMethod,      // line #5720:       'valueAsString: ${valueAsString}]';
	BlankLine,        // line #5721:
}

var wantVMServiceClass65 = []EntityType{
	Unknown,          // line #5724: {
	OtherMethod,      // line #5725:   static ObjRef parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #5726:       json == null ? null : ObjRef._fromJson(json);
	BlankLine,        // line #5727:
	InstanceVariable, // line #5728:   /// A unique identifier for an Object. Passed to the getObject RPC to load
	InstanceVariable, // line #5729:   /// this Object.
	InstanceVariable, // line #5730:   String id;
	BlankLine,        // line #5731:
	InstanceVariable, // line #5732:   /// Provided and set to true if the id of an Object is fixed. If true, the id
	InstanceVariable, // line #5733:   /// of an Object is guaranteed not to change or expire. The object may,
	InstanceVariable, // line #5734:   /// however, still be _Collected_.
	InstanceVariable, // line #5735:   @optional
	InstanceVariable, // line #5736:   bool fixedId;
	BlankLine,        // line #5737:
	MainConstructor,  // line #5738:   ObjRef({
	MainConstructor,  // line #5739:     @required this.id,
	MainConstructor,  // line #5740:     this.fixedId,
	MainConstructor,  // line #5741:   });
	BlankLine,        // line #5742:
	NamedConstructor, // line #5743:   ObjRef._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #5744:     id = json['id'];
	NamedConstructor, // line #5745:     fixedId = json['fixedId'];
	NamedConstructor, // line #5746:   }
	BlankLine,        // line #5747:
	OverrideMethod,   // line #5748:   @override
	OverrideMethod,   // line #5749:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #5750:     var json = <String, dynamic>{};
	OverrideMethod,   // line #5751:     json['type'] = '@Object';
	OverrideMethod,   // line #5752:     json.addAll({
	OverrideMethod,   // line #5753:       'id': id,
	OverrideMethod,   // line #5754:     });
	OverrideMethod,   // line #5755:     _setIfNotNull(json, 'fixedId', fixedId);
	OverrideMethod,   // line #5756:     return json;
	OverrideMethod,   // line #5757:   }
	BlankLine,        // line #5758:
	OtherMethod,      // line #5759:   int get hashCode => id.hashCode;
	BlankLine,        // line #5760:
	OtherMethod,      // line #5761:   operator ==(other) => other is ObjRef && id == other.id;
	BlankLine,        // line #5762:
	OtherMethod,      // line #5763:   String toString() => '[ObjRef type: ${type}, id: ${id}]';
	BlankLine,        // line #5764:
}

var wantVMServiceClass66 = []EntityType{
	Unknown,          // line #5767: {
	OtherMethod,      // line #5768:   static Obj parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #5769:       json == null ? null : Obj._fromJson(json);
	BlankLine,        // line #5770:
	InstanceVariable, // line #5771:   /// A unique identifier for an Object. Passed to the getObject RPC to reload
	InstanceVariable, // line #5772:   /// this Object.
	InstanceVariable, // line #5773:   ///
	InstanceVariable, // line #5774:   /// Some objects may get a new id when they are reloaded.
	InstanceVariable, // line #5775:   String id;
	BlankLine,        // line #5776:
	InstanceVariable, // line #5777:   /// Provided and set to true if the id of an Object is fixed. If true, the id
	InstanceVariable, // line #5778:   /// of an Object is guaranteed not to change or expire. The object may,
	InstanceVariable, // line #5779:   /// however, still be _Collected_.
	InstanceVariable, // line #5780:   @optional
	InstanceVariable, // line #5781:   bool fixedId;
	BlankLine,        // line #5782:
	InstanceVariable, // line #5783:   /// If an object is allocated in the Dart heap, it will have a corresponding
	InstanceVariable, // line #5784:   /// class object.
	InstanceVariable, // line #5785:   ///
	InstanceVariable, // line #5786:   /// The class of a non-instance is not a Dart class, but is instead an
	InstanceVariable, // line #5787:   /// internal vm object.
	InstanceVariable, // line #5788:   ///
	InstanceVariable, // line #5789:   /// Moving an Object into or out of the heap is considered a backwards
	InstanceVariable, // line #5790:   /// compatible change for types other than Instance.
	InstanceVariable, // line #5791:   @optional
	InstanceVariable, // line #5792:   ClassRef classRef;
	BlankLine,        // line #5793:
	InstanceVariable, // line #5794:   /// The size of this object in the heap.
	InstanceVariable, // line #5795:   ///
	InstanceVariable, // line #5796:   /// If an object is not heap-allocated, then this field is omitted.
	InstanceVariable, // line #5797:   ///
	InstanceVariable, // line #5798:   /// Note that the size can be zero for some objects. In the current VM
	InstanceVariable, // line #5799:   /// implementation, this occurs for small integers, which are stored entirely
	InstanceVariable, // line #5800:   /// within their object pointers.
	InstanceVariable, // line #5801:   @optional
	InstanceVariable, // line #5802:   int size;
	BlankLine,        // line #5803:
	MainConstructor,  // line #5804:   Obj({
	MainConstructor,  // line #5805:     @required this.id,
	MainConstructor,  // line #5806:     this.fixedId,
	MainConstructor,  // line #5807:     this.classRef,
	MainConstructor,  // line #5808:     this.size,
	MainConstructor,  // line #5809:   });
	BlankLine,        // line #5810:
	NamedConstructor, // line #5811:   Obj._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #5812:     id = json['id'];
	NamedConstructor, // line #5813:     fixedId = json['fixedId'];
	NamedConstructor, // line #5814:     classRef = createServiceObject(json['class'], const ['ClassRef']);
	NamedConstructor, // line #5815:     size = json['size'];
	NamedConstructor, // line #5816:   }
	BlankLine,        // line #5817:
	OverrideMethod,   // line #5818:   @override
	OverrideMethod,   // line #5819:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #5820:     var json = <String, dynamic>{};
	OverrideMethod,   // line #5821:     json['type'] = 'Object';
	OverrideMethod,   // line #5822:     json.addAll({
	OverrideMethod,   // line #5823:       'id': id,
	OverrideMethod,   // line #5824:     });
	OverrideMethod,   // line #5825:     _setIfNotNull(json, 'fixedId', fixedId);
	OverrideMethod,   // line #5826:     _setIfNotNull(json, 'class', classRef?.toJson());
	OverrideMethod,   // line #5827:     _setIfNotNull(json, 'size', size);
	OverrideMethod,   // line #5828:     return json;
	OverrideMethod,   // line #5829:   }
	BlankLine,        // line #5830:
	OtherMethod,      // line #5831:   int get hashCode => id.hashCode;
	BlankLine,        // line #5832:
	OtherMethod,      // line #5833:   operator ==(other) => other is Obj && id == other.id;
	BlankLine,        // line #5834:
	OtherMethod,      // line #5835:   String toString() => '[Obj type: ${type}, id: ${id}]';
	BlankLine,        // line #5836:
}

var wantVMServiceClass67 = []EntityType{
	Unknown,          // line #5841: {
	OtherMethod,      // line #5842:   static PortList parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #5843:       json == null ? null : PortList._fromJson(json);
	BlankLine,        // line #5844:
	InstanceVariable, // line #5845:   List<InstanceRef> ports;
	BlankLine,        // line #5846:
	MainConstructor,  // line #5847:   PortList({
	MainConstructor,  // line #5848:     @required this.ports,
	MainConstructor,  // line #5849:   });
	BlankLine,        // line #5850:
	NamedConstructor, // line #5851:   PortList._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #5852:     ports = List<InstanceRef>.from(
	NamedConstructor, // line #5853:         createServiceObject(json['ports'], const ['InstanceRef']) ?? []);
	NamedConstructor, // line #5854:   }
	BlankLine,        // line #5855:
	OverrideMethod,   // line #5856:   @override
	OverrideMethod,   // line #5857:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #5858:     var json = <String, dynamic>{};
	OverrideMethod,   // line #5859:     json['type'] = 'PortList';
	OverrideMethod,   // line #5860:     json.addAll({
	OverrideMethod,   // line #5861:       'ports': ports.map((f) => f.toJson()).toList(),
	OverrideMethod,   // line #5862:     });
	OverrideMethod,   // line #5863:     return json;
	OverrideMethod,   // line #5864:   }
	BlankLine,        // line #5865:
	OtherMethod,      // line #5866:   String toString() => '[PortList type: ${type}, ports: ${ports}]';
	BlankLine,        // line #5867:
}

var wantVMServiceClass68 = []EntityType{
	Unknown,          // line #5873: {
	OtherMethod,      // line #5874:   static ProfileFunction parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #5875:       json == null ? null : ProfileFunction._fromJson(json);
	BlankLine,        // line #5876:
	InstanceVariable, // line #5877:   /// The kind of function this object represents.
	InstanceVariable, // line #5878:   String kind;
	BlankLine,        // line #5879:
	InstanceVariable, // line #5880:   /// The number of times function appeared on the stack during sampling events.
	InstanceVariable, // line #5881:   int inclusiveTicks;
	BlankLine,        // line #5882:
	InstanceVariable, // line #5883:   /// The number of times function appeared on the top of the stack during
	InstanceVariable, // line #5884:   /// sampling events.
	InstanceVariable, // line #5885:   int exclusiveTicks;
	BlankLine,        // line #5886:
	InstanceVariable, // line #5887:   /// The resolved URL for the script containing function.
	InstanceVariable, // line #5888:   String resolvedUrl;
	BlankLine,        // line #5889:
	InstanceVariable, // line #5890:   /// The function captured during profiling.
	InstanceVariable, // line #5891:   dynamic function;
	BlankLine,        // line #5892:
	MainConstructor,  // line #5893:   ProfileFunction({
	MainConstructor,  // line #5894:     @required this.kind,
	MainConstructor,  // line #5895:     @required this.inclusiveTicks,
	MainConstructor,  // line #5896:     @required this.exclusiveTicks,
	MainConstructor,  // line #5897:     @required this.resolvedUrl,
	MainConstructor,  // line #5898:     @required this.function,
	MainConstructor,  // line #5899:   });
	BlankLine,        // line #5900:
	NamedConstructor, // line #5901:   ProfileFunction._fromJson(Map<String, dynamic> json) {
	NamedConstructor, // line #5902:     kind = json['kind'];
	NamedConstructor, // line #5903:     inclusiveTicks = json['inclusiveTicks'];
	NamedConstructor, // line #5904:     exclusiveTicks = json['exclusiveTicks'];
	NamedConstructor, // line #5905:     resolvedUrl = json['resolvedUrl'];
	NamedConstructor, // line #5906:     function = createServiceObject(json['function'], const ['dynamic']);
	NamedConstructor, // line #5907:   }
	BlankLine,        // line #5908:
	OtherMethod,      // line #5909:   Map<String, dynamic> toJson() {
	OtherMethod,      // line #5910:     var json = <String, dynamic>{};
	OtherMethod,      // line #5911:     json.addAll({
	OtherMethod,      // line #5912:       'kind': kind,
	OtherMethod,      // line #5913:       'inclusiveTicks': inclusiveTicks,
	OtherMethod,      // line #5914:       'exclusiveTicks': exclusiveTicks,
	OtherMethod,      // line #5915:       'resolvedUrl': resolvedUrl,
	OtherMethod,      // line #5916:       'function': function.toJson(),
	OtherMethod,      // line #5917:     });
	OtherMethod,      // line #5918:     return json;
	OtherMethod,      // line #5919:   }
	BlankLine,        // line #5920:
	OtherMethod,      // line #5921:   String toString() => '[ProfileFunction ' //
	OtherMethod,      // line #5922:       'kind: ${kind}, inclusiveTicks: ${inclusiveTicks}, exclusiveTicks: ${exclusiveTicks}, ' //
	OtherMethod,      // line #5923:       'resolvedUrl: ${resolvedUrl}, function: ${function}]';
	BlankLine,        // line #5924:
}

var wantVMServiceClass69 = []EntityType{
	Unknown,          // line #5930: {
	OtherMethod,      // line #5931:   static ProtocolList parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #5932:       json == null ? null : ProtocolList._fromJson(json);
	BlankLine,        // line #5933:
	InstanceVariable, // line #5934:   /// A list of supported protocols provided by this service.
	InstanceVariable, // line #5935:   List<Protocol> protocols;
	BlankLine,        // line #5936:
	MainConstructor,  // line #5937:   ProtocolList({
	MainConstructor,  // line #5938:     @required this.protocols,
	MainConstructor,  // line #5939:   });
	BlankLine,        // line #5940:
	NamedConstructor, // line #5941:   ProtocolList._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #5942:     protocols = List<Protocol>.from(
	NamedConstructor, // line #5943:         createServiceObject(json['protocols'], const ['Protocol']) ?? []);
	NamedConstructor, // line #5944:   }
	BlankLine,        // line #5945:
	OverrideMethod,   // line #5946:   @override
	OverrideMethod,   // line #5947:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #5948:     var json = <String, dynamic>{};
	OverrideMethod,   // line #5949:     json['type'] = 'ProtocolList';
	OverrideMethod,   // line #5950:     json.addAll({
	OverrideMethod,   // line #5951:       'protocols': protocols.map((f) => f.toJson()).toList(),
	OverrideMethod,   // line #5952:     });
	OverrideMethod,   // line #5953:     return json;
	OverrideMethod,   // line #5954:   }
	BlankLine,        // line #5955:
	OtherMethod,      // line #5956:   String toString() => '[ProtocolList type: ${type}, protocols: ${protocols}]';
	BlankLine,        // line #5957:
}

var wantVMServiceClass70 = []EntityType{
	Unknown,          // line #5960: {
	OtherMethod,      // line #5961:   static Protocol parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #5962:       json == null ? null : Protocol._fromJson(json);
	BlankLine,        // line #5963:
	InstanceVariable, // line #5964:   /// The name of the supported protocol.
	InstanceVariable, // line #5965:   String protocolName;
	BlankLine,        // line #5966:
	InstanceVariable, // line #5967:   /// The major revision of the protocol.
	InstanceVariable, // line #5968:   int major;
	BlankLine,        // line #5969:
	InstanceVariable, // line #5970:   /// The minor revision of the protocol.
	InstanceVariable, // line #5971:   int minor;
	BlankLine,        // line #5972:
	MainConstructor,  // line #5973:   Protocol({
	MainConstructor,  // line #5974:     @required this.protocolName,
	MainConstructor,  // line #5975:     @required this.major,
	MainConstructor,  // line #5976:     @required this.minor,
	MainConstructor,  // line #5977:   });
	BlankLine,        // line #5978:
	NamedConstructor, // line #5979:   Protocol._fromJson(Map<String, dynamic> json) {
	NamedConstructor, // line #5980:     protocolName = json['protocolName'];
	NamedConstructor, // line #5981:     major = json['major'];
	NamedConstructor, // line #5982:     minor = json['minor'];
	NamedConstructor, // line #5983:   }
	BlankLine,        // line #5984:
	OtherMethod,      // line #5985:   Map<String, dynamic> toJson() {
	OtherMethod,      // line #5986:     var json = <String, dynamic>{};
	OtherMethod,      // line #5987:     json.addAll({
	OtherMethod,      // line #5988:       'protocolName': protocolName,
	OtherMethod,      // line #5989:       'major': major,
	OtherMethod,      // line #5990:       'minor': minor,
	OtherMethod,      // line #5991:     });
	OtherMethod,      // line #5992:     return json;
	OtherMethod,      // line #5993:   }
	BlankLine,        // line #5994:
	OtherMethod,      // line #5995:   String toString() => '[Protocol ' //
	OtherMethod,      // line #5996:       'protocolName: ${protocolName}, major: ${major}, minor: ${minor}]';
	BlankLine,        // line #5997:
}

var wantVMServiceClass71 = []EntityType{
	Unknown,          // line #6000: {
	OtherMethod,      // line #6001:   static ProcessMemoryUsage parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #6002:       json == null ? null : ProcessMemoryUsage._fromJson(json);
	BlankLine,        // line #6003:
	InstanceVariable, // line #6004:   ProcessMemoryItem root;
	BlankLine,        // line #6005:
	MainConstructor,  // line #6006:   ProcessMemoryUsage({
	MainConstructor,  // line #6007:     @required this.root,
	MainConstructor,  // line #6008:   });
	BlankLine,        // line #6009:
	NamedConstructor, // line #6010:   ProcessMemoryUsage._fromJson(Map<String, dynamic> json)
	NamedConstructor, // line #6011:       : super._fromJson(json) {
	NamedConstructor, // line #6012:     root = createServiceObject(json['root'], const ['ProcessMemoryItem']);
	NamedConstructor, // line #6013:   }
	BlankLine,        // line #6014:
	OverrideMethod,   // line #6015:   @override
	OverrideMethod,   // line #6016:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #6017:     var json = <String, dynamic>{};
	OverrideMethod,   // line #6018:     json['type'] = 'ProcessMemoryUsage';
	OverrideMethod,   // line #6019:     json.addAll({
	OverrideMethod,   // line #6020:       'root': root.toJson(),
	OverrideMethod,   // line #6021:     });
	OverrideMethod,   // line #6022:     return json;
	OverrideMethod,   // line #6023:   }
	BlankLine,        // line #6024:
	OtherMethod,      // line #6025:   String toString() => '[ProcessMemoryUsage type: ${type}, root: ${root}]';
	BlankLine,        // line #6026:
}

var wantVMServiceClass72 = []EntityType{
	Unknown,          // line #6028: {
	OtherMethod,      // line #6029:   static ProcessMemoryItem parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #6030:       json == null ? null : ProcessMemoryItem._fromJson(json);
	BlankLine,        // line #6031:
	InstanceVariable, // line #6032:   /// A short name for this bucket of memory.
	InstanceVariable, // line #6033:   String name;
	BlankLine,        // line #6034:
	InstanceVariable, // line #6035:   /// A longer description for this item.
	InstanceVariable, // line #6036:   String description;
	BlankLine,        // line #6037:
	InstanceVariable, // line #6038:   /// The amount of memory in bytes. This is a retained size, not a shallow
	InstanceVariable, // line #6039:   /// size. That is, it includes the size of children.
	InstanceVariable, // line #6040:   int size;
	BlankLine,        // line #6041:
	InstanceVariable, // line #6042:   /// Subdivisons of this bucket of memory.
	InstanceVariable, // line #6043:   List<ProcessMemoryItem> children;
	BlankLine,        // line #6044:
	MainConstructor,  // line #6045:   ProcessMemoryItem({
	MainConstructor,  // line #6046:     @required this.name,
	MainConstructor,  // line #6047:     @required this.description,
	MainConstructor,  // line #6048:     @required this.size,
	MainConstructor,  // line #6049:     @required this.children,
	MainConstructor,  // line #6050:   });
	BlankLine,        // line #6051:
	NamedConstructor, // line #6052:   ProcessMemoryItem._fromJson(Map<String, dynamic> json) {
	NamedConstructor, // line #6053:     name = json['name'];
	NamedConstructor, // line #6054:     description = json['description'];
	NamedConstructor, // line #6055:     size = json['size'];
	NamedConstructor, // line #6056:     children = List<ProcessMemoryItem>.from(
	NamedConstructor, // line #6057:         createServiceObject(json['children'], const ['ProcessMemoryItem']) ??
	NamedConstructor, // line #6058:             []);
	NamedConstructor, // line #6059:   }
	BlankLine,        // line #6060:
	OtherMethod,      // line #6061:   Map<String, dynamic> toJson() {
	OtherMethod,      // line #6062:     var json = <String, dynamic>{};
	OtherMethod,      // line #6063:     json.addAll({
	OtherMethod,      // line #6064:       'name': name,
	OtherMethod,      // line #6065:       'description': description,
	OtherMethod,      // line #6066:       'size': size,
	OtherMethod,      // line #6067:       'children': children.map((f) => f.toJson()).toList(),
	OtherMethod,      // line #6068:     });
	OtherMethod,      // line #6069:     return json;
	OtherMethod,      // line #6070:   }
	BlankLine,        // line #6071:
	OtherMethod,      // line #6072:   String toString() => '[ProcessMemoryItem ' //
	OtherMethod,      // line #6073:       'name: ${name}, description: ${description}, size: ${size}, ' //
	OtherMethod,      // line #6074:       'children: ${children}]';
	BlankLine,        // line #6075:
}

var wantVMServiceClass73 = []EntityType{
	Unknown,          // line #6077: {
	OtherMethod,      // line #6078:   static ReloadReport parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #6079:       json == null ? null : ReloadReport._fromJson(json);
	BlankLine,        // line #6080:
	InstanceVariable, // line #6081:   /// Did the reload succeed or fail?
	InstanceVariable, // line #6082:   bool success;
	BlankLine,        // line #6083:
	MainConstructor,  // line #6084:   ReloadReport({
	MainConstructor,  // line #6085:     @required this.success,
	MainConstructor,  // line #6086:   });
	BlankLine,        // line #6087:
	NamedConstructor, // line #6088:   ReloadReport._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #6089:     success = json['success'];
	NamedConstructor, // line #6090:   }
	BlankLine,        // line #6091:
	OverrideMethod,   // line #6092:   @override
	OverrideMethod,   // line #6093:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #6094:     var json = <String, dynamic>{};
	OverrideMethod,   // line #6095:     json['type'] = 'ReloadReport';
	OverrideMethod,   // line #6096:     json.addAll({
	OverrideMethod,   // line #6097:       'success': success,
	OverrideMethod,   // line #6098:     });
	OverrideMethod,   // line #6099:     return json;
	OverrideMethod,   // line #6100:   }
	BlankLine,        // line #6101:
	OtherMethod,      // line #6102:   String toString() => '[ReloadReport type: ${type}, success: ${success}]';
	BlankLine,        // line #6103:
}

var wantVMServiceClass74 = []EntityType{
	Unknown,          // line #6106: {
	OtherMethod,      // line #6107:   static RetainingObject parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #6108:       json == null ? null : RetainingObject._fromJson(json);
	BlankLine,        // line #6109:
	InstanceVariable, // line #6110:   /// An object that is part of a retaining path.
	InstanceVariable, // line #6111:   ObjRef value;
	BlankLine,        // line #6112:
	InstanceVariable, // line #6113:   /// The offset of the retaining object in a containing list.
	InstanceVariable, // line #6114:   @optional
	InstanceVariable, // line #6115:   int parentListIndex;
	BlankLine,        // line #6116:
	InstanceVariable, // line #6117:   /// The key mapping to the retaining object in a containing map.
	InstanceVariable, // line #6118:   @optional
	InstanceVariable, // line #6119:   ObjRef parentMapKey;
	BlankLine,        // line #6120:
	InstanceVariable, // line #6121:   /// The name of the field containing the retaining object within an object.
	InstanceVariable, // line #6122:   @optional
	InstanceVariable, // line #6123:   String parentField;
	BlankLine,        // line #6124:
	MainConstructor,  // line #6125:   RetainingObject({
	MainConstructor,  // line #6126:     @required this.value,
	MainConstructor,  // line #6127:     this.parentListIndex,
	MainConstructor,  // line #6128:     this.parentMapKey,
	MainConstructor,  // line #6129:     this.parentField,
	MainConstructor,  // line #6130:   });
	BlankLine,        // line #6131:
	NamedConstructor, // line #6132:   RetainingObject._fromJson(Map<String, dynamic> json) {
	NamedConstructor, // line #6133:     value = createServiceObject(json['value'], const ['ObjRef']);
	NamedConstructor, // line #6134:     parentListIndex = json['parentListIndex'];
	NamedConstructor, // line #6135:     parentMapKey = createServiceObject(json['parentMapKey'], const ['ObjRef']);
	NamedConstructor, // line #6136:     parentField = json['parentField'];
	NamedConstructor, // line #6137:   }
	BlankLine,        // line #6138:
	OtherMethod,      // line #6139:   Map<String, dynamic> toJson() {
	OtherMethod,      // line #6140:     var json = <String, dynamic>{};
	OtherMethod,      // line #6141:     json.addAll({
	OtherMethod,      // line #6142:       'value': value.toJson(),
	OtherMethod,      // line #6143:     });
	OtherMethod,      // line #6144:     _setIfNotNull(json, 'parentListIndex', parentListIndex);
	OtherMethod,      // line #6145:     _setIfNotNull(json, 'parentMapKey', parentMapKey?.toJson());
	OtherMethod,      // line #6146:     _setIfNotNull(json, 'parentField', parentField);
	OtherMethod,      // line #6147:     return json;
	OtherMethod,      // line #6148:   }
	BlankLine,        // line #6149:
	OtherMethod,      // line #6150:   String toString() => '[RetainingObject value: ${value}]';
	BlankLine,        // line #6151:
}

var wantVMServiceClass75 = []EntityType{
	Unknown,          // line #6154: {
	OtherMethod,      // line #6155:   static RetainingPath parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #6156:       json == null ? null : RetainingPath._fromJson(json);
	BlankLine,        // line #6157:
	InstanceVariable, // line #6158:   /// The length of the retaining path.
	InstanceVariable, // line #6159:   int length;
	BlankLine,        // line #6160:
	InstanceVariable, // line #6161:   /// The type of GC root which is holding a reference to the specified object.
	InstanceVariable, // line #6162:   /// Possible values include:  * class table  * local handle  * persistent
	InstanceVariable, // line #6163:   /// handle  * stack  * user global  * weak persistent handle  * unknown
	InstanceVariable, // line #6164:   String gcRootType;
	BlankLine,        // line #6165:
	InstanceVariable, // line #6166:   /// The chain of objects which make up the retaining path.
	InstanceVariable, // line #6167:   List<RetainingObject> elements;
	BlankLine,        // line #6168:
	MainConstructor,  // line #6169:   RetainingPath({
	MainConstructor,  // line #6170:     @required this.length,
	MainConstructor,  // line #6171:     @required this.gcRootType,
	MainConstructor,  // line #6172:     @required this.elements,
	MainConstructor,  // line #6173:   });
	BlankLine,        // line #6174:
	NamedConstructor, // line #6175:   RetainingPath._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #6176:     length = json['length'];
	NamedConstructor, // line #6177:     gcRootType = json['gcRootType'];
	NamedConstructor, // line #6178:     elements = List<RetainingObject>.from(
	NamedConstructor, // line #6179:         createServiceObject(json['elements'], const ['RetainingObject']) ?? []);
	NamedConstructor, // line #6180:   }
	BlankLine,        // line #6181:
	OverrideMethod,   // line #6182:   @override
	OverrideMethod,   // line #6183:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #6184:     var json = <String, dynamic>{};
	OverrideMethod,   // line #6185:     json['type'] = 'RetainingPath';
	OverrideMethod,   // line #6186:     json.addAll({
	OverrideMethod,   // line #6187:       'length': length,
	OverrideMethod,   // line #6188:       'gcRootType': gcRootType,
	OverrideMethod,   // line #6189:       'elements': elements.map((f) => f.toJson()).toList(),
	OverrideMethod,   // line #6190:     });
	OverrideMethod,   // line #6191:     return json;
	OverrideMethod,   // line #6192:   }
	BlankLine,        // line #6193:
	OtherMethod,      // line #6194:   String toString() => '[RetainingPath ' //
	OtherMethod,      // line #6195:       'type: ${type}, length: ${length}, gcRootType: ${gcRootType}, ' //
	OtherMethod,      // line #6196:       'elements: ${elements}]';
	BlankLine,        // line #6197:
}

var wantVMServiceClass76 = []EntityType{
	Unknown,          // line #6202: {
	OtherMethod,      // line #6203:   static Response parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #6204:       json == null ? null : Response._fromJson(json);
	BlankLine,        // line #6205:
	InstanceVariable, // line #6206:   Map<String, dynamic> json;
	BlankLine,        // line #6207:
	InstanceVariable, // line #6208:   /// Every response returned by the VM Service has the type property. This
	InstanceVariable, // line #6209:   /// allows the client distinguish between different kinds of responses.
	InstanceVariable, // line #6210:   String type;
	BlankLine,        // line #6211:
	MainConstructor,  // line #6212:   Response({
	MainConstructor,  // line #6213:     @required this.type,
	MainConstructor,  // line #6214:   });
	BlankLine,        // line #6215:
	NamedConstructor, // line #6216:   Response._fromJson(this.json) {
	NamedConstructor, // line #6217:     type = json['type'];
	NamedConstructor, // line #6218:   }
	BlankLine,        // line #6219:
	OtherMethod,      // line #6220:   Map<String, dynamic> toJson() {
	OtherMethod,      // line #6221:     var result = json == null ? <String, dynamic>{} : Map.of(json);
	OtherMethod,      // line #6222:     result['type'] = type ?? 'Response';
	OtherMethod,      // line #6223:     return result;
	OtherMethod,      // line #6224:   }
	BlankLine,        // line #6225:
	OtherMethod,      // line #6226:   String toString() => '[Response type: ${type}]';
	BlankLine,        // line #6227:
}

var wantVMServiceClass77 = []EntityType{
	Unknown,           // line #6233: {
	OtherMethod,       // line #6234:   static Sentinel parse(Map<String, dynamic> json) =>
	OtherMethod,       // line #6235:       json == null ? null : Sentinel._fromJson(json);
	BlankLine,         // line #6236:
	SingleLineComment, // line #6237:   /// What kind of sentinel is this?
	MultiLineComment,  // line #6238:   /*SentinelKind*/ String kind;
	BlankLine,         // line #6239:
	InstanceVariable,  // line #6240:   /// A reasonable string representation of this sentinel.
	InstanceVariable,  // line #6241:   String valueAsString;
	BlankLine,         // line #6242:
	MainConstructor,   // line #6243:   Sentinel({
	MainConstructor,   // line #6244:     @required this.kind,
	MainConstructor,   // line #6245:     @required this.valueAsString,
	MainConstructor,   // line #6246:   });
	BlankLine,         // line #6247:
	NamedConstructor,  // line #6248:   Sentinel._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor,  // line #6249:     kind = json['kind'];
	NamedConstructor,  // line #6250:     valueAsString = json['valueAsString'];
	NamedConstructor,  // line #6251:   }
	BlankLine,         // line #6252:
	OverrideMethod,    // line #6253:   @override
	OverrideMethod,    // line #6254:   Map<String, dynamic> toJson() {
	OverrideMethod,    // line #6255:     var json = <String, dynamic>{};
	OverrideMethod,    // line #6256:     json['type'] = 'Sentinel';
	OverrideMethod,    // line #6257:     json.addAll({
	OverrideMethod,    // line #6258:       'kind': kind,
	OverrideMethod,    // line #6259:       'valueAsString': valueAsString,
	OverrideMethod,    // line #6260:     });
	OverrideMethod,    // line #6261:     return json;
	OverrideMethod,    // line #6262:   }
	BlankLine,         // line #6263:
	OtherMethod,       // line #6264:   String toString() => '[Sentinel ' //
	OtherMethod,       // line #6265:       'type: ${type}, kind: ${kind}, valueAsString: ${valueAsString}]';
	BlankLine,         // line #6266:
}

var wantVMServiceClass78 = []EntityType{
	Unknown,          // line #6269: {
	OtherMethod,      // line #6270:   static ScriptRef parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #6271:       json == null ? null : ScriptRef._fromJson(json);
	BlankLine,        // line #6272:
	InstanceVariable, // line #6273:   /// The uri from which this script was loaded.
	InstanceVariable, // line #6274:   String uri;
	BlankLine,        // line #6275:
	MainConstructor,  // line #6276:   ScriptRef({
	MainConstructor,  // line #6277:     @required this.uri,
	MainConstructor,  // line #6278:     @required String id,
	MainConstructor,  // line #6279:   }) : super(id: id);
	BlankLine,        // line #6280:
	NamedConstructor, // line #6281:   ScriptRef._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #6282:     uri = json['uri'];
	NamedConstructor, // line #6283:   }
	BlankLine,        // line #6284:
	OverrideMethod,   // line #6285:   @override
	OverrideMethod,   // line #6286:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #6287:     var json = super.toJson();
	OverrideMethod,   // line #6288:     json['type'] = '@Script';
	OverrideMethod,   // line #6289:     json.addAll({
	OverrideMethod,   // line #6290:       'uri': uri,
	OverrideMethod,   // line #6291:     });
	OverrideMethod,   // line #6292:     return json;
	OverrideMethod,   // line #6293:   }
	BlankLine,        // line #6294:
	OtherMethod,      // line #6295:   int get hashCode => id.hashCode;
	BlankLine,        // line #6296:
	OtherMethod,      // line #6297:   operator ==(other) => other is ScriptRef && id == other.id;
	BlankLine,        // line #6298:
	OtherMethod,      // line #6299:   String toString() => '[ScriptRef type: ${type}, id: ${id}, uri: ${uri}]';
	BlankLine,        // line #6300:
}

var wantVMServiceClass79 = []EntityType{
	Unknown,                 // line #6332: {
	OtherMethod,             // line #6333:   static Script parse(Map<String, dynamic> json) =>
	OtherMethod,             // line #6334:       json == null ? null : Script._fromJson(json);
	BlankLine,               // line #6335:
	PrivateInstanceVariable, // line #6336:   final _tokenToLine = <int, int>{};
	PrivateInstanceVariable, // line #6337:   final _tokenToColumn = <int, int>{};
	BlankLine,               // line #6338:
	InstanceVariable,        // line #6339:   /// The uri from which this script was loaded.
	InstanceVariable,        // line #6340:   String uri;
	BlankLine,               // line #6341:
	InstanceVariable,        // line #6342:   /// The library which owns this script.
	InstanceVariable,        // line #6343:   LibraryRef library;
	BlankLine,               // line #6344:
	InstanceVariable,        // line #6345:   @optional
	InstanceVariable,        // line #6346:   int lineOffset;
	BlankLine,               // line #6347:
	InstanceVariable,        // line #6348:   @optional
	InstanceVariable,        // line #6349:   int columnOffset;
	BlankLine,               // line #6350:
	InstanceVariable,        // line #6351:   /// The source code for this script. This can be null for certain built-in
	InstanceVariable,        // line #6352:   /// scripts.
	InstanceVariable,        // line #6353:   @optional
	InstanceVariable,        // line #6354:   String source;
	BlankLine,               // line #6355:
	InstanceVariable,        // line #6356:   /// A table encoding a mapping from token position to line and column. This
	InstanceVariable,        // line #6357:   /// field is null if sources aren't available.
	InstanceVariable,        // line #6358:   @optional
	InstanceVariable,        // line #6359:   List<List<int>> tokenPosTable;
	BlankLine,               // line #6360:
	MainConstructor,         // line #6361:   Script({
	MainConstructor,         // line #6362:     @required this.uri,
	MainConstructor,         // line #6363:     @required this.library,
	MainConstructor,         // line #6364:     @required String id,
	MainConstructor,         // line #6365:     this.lineOffset,
	MainConstructor,         // line #6366:     this.columnOffset,
	MainConstructor,         // line #6367:     this.source,
	MainConstructor,         // line #6368:     this.tokenPosTable,
	MainConstructor,         // line #6369:   }) : super(id: id);
	BlankLine,               // line #6370:
	NamedConstructor,        // line #6371:   Script._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor,        // line #6372:     uri = json['uri'];
	NamedConstructor,        // line #6373:     library = createServiceObject(json['library'], const ['LibraryRef']);
	NamedConstructor,        // line #6374:     lineOffset = json['lineOffset'];
	NamedConstructor,        // line #6375:     columnOffset = json['columnOffset'];
	NamedConstructor,        // line #6376:     source = json['source'];
	NamedConstructor,        // line #6377:     tokenPosTable = json['tokenPosTable'] == null
	NamedConstructor,        // line #6378:         ? null
	NamedConstructor,        // line #6379:         : List<List<int>>.from(
	NamedConstructor,        // line #6380:             json['tokenPosTable'].map((dynamic list) => List<int>.from(list)));
	NamedConstructor,        // line #6381:     _parseTokenPosTable();
	NamedConstructor,        // line #6382:   }
	BlankLine,               // line #6383:
	OtherMethod,             // line #6384:   /// This function maps a token position to a line number.
	OtherMethod,             // line #6385:   /// The VM considers the first line to be line 1.
	OtherMethod,             // line #6386:   int getLineNumberFromTokenPos(int tokenPos) => _tokenToLine[tokenPos];
	BlankLine,               // line #6387:
	OtherMethod,             // line #6388:   /// This function maps a token position to a column number.
	OtherMethod,             // line #6389:   /// The VM considers the first column to be column 1.
	OtherMethod,             // line #6390:   int getColumnNumberFromTokenPos(int tokenPos) => _tokenToColumn[tokenPos];
	BlankLine,               // line #6391:
	OtherMethod,             // line #6392:   void _parseTokenPosTable() {
	OtherMethod,             // line #6393:     if (tokenPosTable == null) {
	OtherMethod,             // line #6394:       return;
	OtherMethod,             // line #6395:     }
	OtherMethod,             // line #6396:     final lineSet = Set<int>();
	OtherMethod,             // line #6397:     for (List line in tokenPosTable) {
	OtherMethod,             // line #6398:       // Each entry begins with a line number...
	OtherMethod,             // line #6399:       int lineNumber = line[0];
	OtherMethod,             // line #6400:       lineSet.add(lineNumber);
	OtherMethod,             // line #6401:       for (var pos = 1; pos < line.length; pos += 2) {
	OtherMethod,             // line #6402:         // ...and is followed by (token offset, col number) pairs.
	OtherMethod,             // line #6403:         final int tokenOffset = line[pos];
	OtherMethod,             // line #6404:         final int colNumber = line[pos + 1];
	OtherMethod,             // line #6405:         _tokenToLine[tokenOffset] = lineNumber;
	OtherMethod,             // line #6406:         _tokenToColumn[tokenOffset] = colNumber;
	OtherMethod,             // line #6407:       }
	OtherMethod,             // line #6408:     }
	OtherMethod,             // line #6409:   }
	BlankLine,               // line #6410:
	OverrideMethod,          // line #6411:   @override
	OverrideMethod,          // line #6412:   Map<String, dynamic> toJson() {
	OverrideMethod,          // line #6413:     var json = super.toJson();
	OverrideMethod,          // line #6414:     json['type'] = 'Script';
	OverrideMethod,          // line #6415:     json.addAll({
	OverrideMethod,          // line #6416:       'uri': uri,
	OverrideMethod,          // line #6417:       'library': library.toJson(),
	OverrideMethod,          // line #6418:     });
	OverrideMethod,          // line #6419:     _setIfNotNull(json, 'lineOffset', lineOffset);
	OverrideMethod,          // line #6420:     _setIfNotNull(json, 'columnOffset', columnOffset);
	OverrideMethod,          // line #6421:     _setIfNotNull(json, 'source', source);
	OverrideMethod,          // line #6422:     _setIfNotNull(json, 'tokenPosTable',
	OverrideMethod,          // line #6423:         tokenPosTable?.map((f) => f?.toList())?.toList());
	OverrideMethod,          // line #6424:     return json;
	OverrideMethod,          // line #6425:   }
	BlankLine,               // line #6426:
	OtherMethod,             // line #6427:   int get hashCode => id.hashCode;
	BlankLine,               // line #6428:
	OtherMethod,             // line #6429:   operator ==(other) => other is Script && id == other.id;
	BlankLine,               // line #6430:
	OtherMethod,             // line #6431:   String toString() =>
	OtherMethod,             // line #6432:       '[Script type: ${type}, id: ${id}, uri: ${uri}, library: ${library}]';
	BlankLine,               // line #6433:
}

var wantVMServiceClass80 = []EntityType{
	Unknown,          // line #6435: {
	OtherMethod,      // line #6436:   static ScriptList parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #6437:       json == null ? null : ScriptList._fromJson(json);
	BlankLine,        // line #6438:
	InstanceVariable, // line #6439:   List<ScriptRef> scripts;
	BlankLine,        // line #6440:
	MainConstructor,  // line #6441:   ScriptList({
	MainConstructor,  // line #6442:     @required this.scripts,
	MainConstructor,  // line #6443:   });
	BlankLine,        // line #6444:
	NamedConstructor, // line #6445:   ScriptList._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #6446:     scripts = List<ScriptRef>.from(
	NamedConstructor, // line #6447:         createServiceObject(json['scripts'], const ['ScriptRef']) ?? []);
	NamedConstructor, // line #6448:   }
	BlankLine,        // line #6449:
	OverrideMethod,   // line #6450:   @override
	OverrideMethod,   // line #6451:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #6452:     var json = <String, dynamic>{};
	OverrideMethod,   // line #6453:     json['type'] = 'ScriptList';
	OverrideMethod,   // line #6454:     json.addAll({
	OverrideMethod,   // line #6455:       'scripts': scripts.map((f) => f.toJson()).toList(),
	OverrideMethod,   // line #6456:     });
	OverrideMethod,   // line #6457:     return json;
	OverrideMethod,   // line #6458:   }
	BlankLine,        // line #6459:
	OtherMethod,      // line #6460:   String toString() => '[ScriptList type: ${type}, scripts: ${scripts}]';
	BlankLine,        // line #6461:
}

var wantVMServiceClass81 = []EntityType{
	Unknown,          // line #6465: {
	OtherMethod,      // line #6466:   static SourceLocation parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #6467:       json == null ? null : SourceLocation._fromJson(json);
	BlankLine,        // line #6468:
	InstanceVariable, // line #6469:   /// The script containing the source location.
	InstanceVariable, // line #6470:   ScriptRef script;
	BlankLine,        // line #6471:
	InstanceVariable, // line #6472:   /// The first token of the location.
	InstanceVariable, // line #6473:   int tokenPos;
	BlankLine,        // line #6474:
	InstanceVariable, // line #6475:   /// The last token of the location if this is a range.
	InstanceVariable, // line #6476:   @optional
	InstanceVariable, // line #6477:   int endTokenPos;
	BlankLine,        // line #6478:
	MainConstructor,  // line #6479:   SourceLocation({
	MainConstructor,  // line #6480:     @required this.script,
	MainConstructor,  // line #6481:     @required this.tokenPos,
	MainConstructor,  // line #6482:     this.endTokenPos,
	MainConstructor,  // line #6483:   });
	BlankLine,        // line #6484:
	NamedConstructor, // line #6485:   SourceLocation._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #6486:     script = createServiceObject(json['script'], const ['ScriptRef']);
	NamedConstructor, // line #6487:     tokenPos = json['tokenPos'];
	NamedConstructor, // line #6488:     endTokenPos = json['endTokenPos'];
	NamedConstructor, // line #6489:   }
	BlankLine,        // line #6490:
	OverrideMethod,   // line #6491:   @override
	OverrideMethod,   // line #6492:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #6493:     var json = <String, dynamic>{};
	OverrideMethod,   // line #6494:     json['type'] = 'SourceLocation';
	OverrideMethod,   // line #6495:     json.addAll({
	OverrideMethod,   // line #6496:       'script': script.toJson(),
	OverrideMethod,   // line #6497:       'tokenPos': tokenPos,
	OverrideMethod,   // line #6498:     });
	OverrideMethod,   // line #6499:     _setIfNotNull(json, 'endTokenPos', endTokenPos);
	OverrideMethod,   // line #6500:     return json;
	OverrideMethod,   // line #6501:   }
	BlankLine,        // line #6502:
	OtherMethod,      // line #6503:   String toString() =>
	OtherMethod,      // line #6504:       '[SourceLocation type: ${type}, script: ${script}, tokenPos: ${tokenPos}]';
	BlankLine,        // line #6505:
}

var wantVMServiceClass82 = []EntityType{
	Unknown,          // line #6509: {
	OtherMethod,      // line #6510:   static SourceReport parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #6511:       json == null ? null : SourceReport._fromJson(json);
	BlankLine,        // line #6512:
	InstanceVariable, // line #6513:   /// A list of ranges in the program source.  These ranges correspond to ranges
	InstanceVariable, // line #6514:   /// of executable code in the user's program (functions, methods,
	InstanceVariable, // line #6515:   /// constructors, etc.)
	InstanceVariable, // line #6516:   ///
	InstanceVariable, // line #6517:   /// Note that ranges may nest in other ranges, in the case of nested
	InstanceVariable, // line #6518:   /// functions.
	InstanceVariable, // line #6519:   ///
	InstanceVariable, // line #6520:   /// Note that ranges may be duplicated, in the case of mixins.
	InstanceVariable, // line #6521:   List<SourceReportRange> ranges;
	BlankLine,        // line #6522:
	InstanceVariable, // line #6523:   /// A list of scripts, referenced by index in the report's ranges.
	InstanceVariable, // line #6524:   List<ScriptRef> scripts;
	BlankLine,        // line #6525:
	MainConstructor,  // line #6526:   SourceReport({
	MainConstructor,  // line #6527:     @required this.ranges,
	MainConstructor,  // line #6528:     @required this.scripts,
	MainConstructor,  // line #6529:   });
	BlankLine,        // line #6530:
	NamedConstructor, // line #6531:   SourceReport._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #6532:     ranges = List<SourceReportRange>.from(
	NamedConstructor, // line #6533:         _createSpecificObject(json['ranges'], SourceReportRange.parse));
	NamedConstructor, // line #6534:     scripts = List<ScriptRef>.from(
	NamedConstructor, // line #6535:         createServiceObject(json['scripts'], const ['ScriptRef']) ?? []);
	NamedConstructor, // line #6536:   }
	BlankLine,        // line #6537:
	OverrideMethod,   // line #6538:   @override
	OverrideMethod,   // line #6539:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #6540:     var json = <String, dynamic>{};
	OverrideMethod,   // line #6541:     json['type'] = 'SourceReport';
	OverrideMethod,   // line #6542:     json.addAll({
	OverrideMethod,   // line #6543:       'ranges': ranges.map((f) => f.toJson()).toList(),
	OverrideMethod,   // line #6544:       'scripts': scripts.map((f) => f.toJson()).toList(),
	OverrideMethod,   // line #6545:     });
	OverrideMethod,   // line #6546:     return json;
	OverrideMethod,   // line #6547:   }
	BlankLine,        // line #6548:
	OtherMethod,      // line #6549:   String toString() =>
	OtherMethod,      // line #6550:       '[SourceReport type: ${type}, ranges: ${ranges}, scripts: ${scripts}]';
	BlankLine,        // line #6551:
}

var wantVMServiceClass83 = []EntityType{
	Unknown,          // line #6558: {
	OtherMethod,      // line #6559:   static SourceReportCoverage parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #6560:       json == null ? null : SourceReportCoverage._fromJson(json);
	BlankLine,        // line #6561:
	InstanceVariable, // line #6562:   /// A list of token positions in a SourceReportRange which have been executed.
	InstanceVariable, // line #6563:   /// The list is sorted.
	InstanceVariable, // line #6564:   List<int> hits;
	BlankLine,        // line #6565:
	InstanceVariable, // line #6566:   /// A list of token positions in a SourceReportRange which have not been
	InstanceVariable, // line #6567:   /// executed.  The list is sorted.
	InstanceVariable, // line #6568:   List<int> misses;
	BlankLine,        // line #6569:
	MainConstructor,  // line #6570:   SourceReportCoverage({
	MainConstructor,  // line #6571:     @required this.hits,
	MainConstructor,  // line #6572:     @required this.misses,
	MainConstructor,  // line #6573:   });
	BlankLine,        // line #6574:
	NamedConstructor, // line #6575:   SourceReportCoverage._fromJson(Map<String, dynamic> json) {
	NamedConstructor, // line #6576:     hits = List<int>.from(json['hits']);
	NamedConstructor, // line #6577:     misses = List<int>.from(json['misses']);
	NamedConstructor, // line #6578:   }
	BlankLine,        // line #6579:
	OtherMethod,      // line #6580:   Map<String, dynamic> toJson() {
	OtherMethod,      // line #6581:     var json = <String, dynamic>{};
	OtherMethod,      // line #6582:     json.addAll({
	OtherMethod,      // line #6583:       'hits': hits.map((f) => f).toList(),
	OtherMethod,      // line #6584:       'misses': misses.map((f) => f).toList(),
	OtherMethod,      // line #6585:     });
	OtherMethod,      // line #6586:     return json;
	OtherMethod,      // line #6587:   }
	BlankLine,        // line #6588:
	OtherMethod,      // line #6589:   String toString() =>
	OtherMethod,      // line #6590:       '[SourceReportCoverage hits: ${hits}, misses: ${misses}]';
	BlankLine,        // line #6591:
}

var wantVMServiceClass84 = []EntityType{
	Unknown,          // line #6599: {
	OtherMethod,      // line #6600:   static SourceReportRange parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #6601:       json == null ? null : SourceReportRange._fromJson(json);
	BlankLine,        // line #6602:
	InstanceVariable, // line #6603:   /// An index into the script table of the SourceReport, indicating which
	InstanceVariable, // line #6604:   /// script contains this range of code.
	InstanceVariable, // line #6605:   int scriptIndex;
	BlankLine,        // line #6606:
	InstanceVariable, // line #6607:   /// The token position at which this range begins.
	InstanceVariable, // line #6608:   int startPos;
	BlankLine,        // line #6609:
	InstanceVariable, // line #6610:   /// The token position at which this range ends.  Inclusive.
	InstanceVariable, // line #6611:   int endPos;
	BlankLine,        // line #6612:
	InstanceVariable, // line #6613:   /// Has this range been compiled by the Dart VM?
	InstanceVariable, // line #6614:   bool compiled;
	BlankLine,        // line #6615:
	InstanceVariable, // line #6616:   /// The error while attempting to compile this range, if this report was
	InstanceVariable, // line #6617:   /// generated with forceCompile=true.
	InstanceVariable, // line #6618:   @optional
	InstanceVariable, // line #6619:   ErrorRef error;
	BlankLine,        // line #6620:
	InstanceVariable, // line #6621:   /// Code coverage information for this range.  Provided only when the Coverage
	InstanceVariable, // line #6622:   /// report has been requested and the range has been compiled.
	InstanceVariable, // line #6623:   @optional
	InstanceVariable, // line #6624:   SourceReportCoverage coverage;
	BlankLine,        // line #6625:
	InstanceVariable, // line #6626:   /// Possible breakpoint information for this range, represented as a sorted
	InstanceVariable, // line #6627:   /// list of token positions.  Provided only when the when the
	InstanceVariable, // line #6628:   /// PossibleBreakpoint report has been requested and the range has been
	InstanceVariable, // line #6629:   /// compiled.
	InstanceVariable, // line #6630:   @optional
	InstanceVariable, // line #6631:   List<int> possibleBreakpoints;
	BlankLine,        // line #6632:
	MainConstructor,  // line #6633:   SourceReportRange({
	MainConstructor,  // line #6634:     @required this.scriptIndex,
	MainConstructor,  // line #6635:     @required this.startPos,
	MainConstructor,  // line #6636:     @required this.endPos,
	MainConstructor,  // line #6637:     @required this.compiled,
	MainConstructor,  // line #6638:     this.error,
	MainConstructor,  // line #6639:     this.coverage,
	MainConstructor,  // line #6640:     this.possibleBreakpoints,
	MainConstructor,  // line #6641:   });
	BlankLine,        // line #6642:
	NamedConstructor, // line #6643:   SourceReportRange._fromJson(Map<String, dynamic> json) {
	NamedConstructor, // line #6644:     scriptIndex = json['scriptIndex'];
	NamedConstructor, // line #6645:     startPos = json['startPos'];
	NamedConstructor, // line #6646:     endPos = json['endPos'];
	NamedConstructor, // line #6647:     compiled = json['compiled'];
	NamedConstructor, // line #6648:     error = createServiceObject(json['error'], const ['ErrorRef']);
	NamedConstructor, // line #6649:     coverage =
	NamedConstructor, // line #6650:         _createSpecificObject(json['coverage'], SourceReportCoverage.parse);
	NamedConstructor, // line #6651:     possibleBreakpoints = json['possibleBreakpoints'] == null
	NamedConstructor, // line #6652:         ? null
	NamedConstructor, // line #6653:         : List<int>.from(json['possibleBreakpoints']);
	NamedConstructor, // line #6654:   }
	BlankLine,        // line #6655:
	OtherMethod,      // line #6656:   Map<String, dynamic> toJson() {
	OtherMethod,      // line #6657:     var json = <String, dynamic>{};
	OtherMethod,      // line #6658:     json.addAll({
	OtherMethod,      // line #6659:       'scriptIndex': scriptIndex,
	OtherMethod,      // line #6660:       'startPos': startPos,
	OtherMethod,      // line #6661:       'endPos': endPos,
	OtherMethod,      // line #6662:       'compiled': compiled,
	OtherMethod,      // line #6663:     });
	OtherMethod,      // line #6664:     _setIfNotNull(json, 'error', error?.toJson());
	OtherMethod,      // line #6665:     _setIfNotNull(json, 'coverage', coverage?.toJson());
	OtherMethod,      // line #6666:     _setIfNotNull(json, 'possibleBreakpoints',
	OtherMethod,      // line #6667:         possibleBreakpoints?.map((f) => f)?.toList());
	OtherMethod,      // line #6668:     return json;
	OtherMethod,      // line #6669:   }
	BlankLine,        // line #6670:
	OtherMethod,      // line #6671:   String toString() => '[SourceReportRange ' //
	OtherMethod,      // line #6672:       'scriptIndex: ${scriptIndex}, startPos: ${startPos}, endPos: ${endPos}, ' //
	OtherMethod,      // line #6673:       'compiled: ${compiled}]';
	BlankLine,        // line #6674:
}

var wantVMServiceClass85 = []EntityType{
	Unknown,          // line #6680: {
	OtherMethod,      // line #6681:   static Stack parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #6682:       json == null ? null : Stack._fromJson(json);
	BlankLine,        // line #6683:
	InstanceVariable, // line #6684:   /// A list of frames that make up the synchronous stack, rooted at the message
	InstanceVariable, // line #6685:   /// loop (i.e., the frames since the last asynchronous gap or the isolate's
	InstanceVariable, // line #6686:   /// entrypoint).
	InstanceVariable, // line #6687:   List<Frame> frames;
	BlankLine,        // line #6688:
	InstanceVariable, // line #6689:   /// A list of frames representing the asynchronous path. Comparable to
	InstanceVariable, // line #6690:   /// `awaiterFrames`, if provided, although some frames may be different.
	InstanceVariable, // line #6691:   @optional
	InstanceVariable, // line #6692:   List<Frame> asyncCausalFrames;
	BlankLine,        // line #6693:
	InstanceVariable, // line #6694:   /// A list of frames representing the asynchronous path. Comparable to
	InstanceVariable, // line #6695:   /// `asyncCausalFrames`, if provided, although some frames may be different.
	InstanceVariable, // line #6696:   @optional
	InstanceVariable, // line #6697:   List<Frame> awaiterFrames;
	BlankLine,        // line #6698:
	InstanceVariable, // line #6699:   /// A list of messages in the isolate's message queue.
	InstanceVariable, // line #6700:   List<Message> messages;
	BlankLine,        // line #6701:
	InstanceVariable, // line #6702:   /// Specifies whether or not this stack is complete or has been artificially
	InstanceVariable, // line #6703:   /// truncated.
	InstanceVariable, // line #6704:   bool truncated;
	BlankLine,        // line #6705:
	MainConstructor,  // line #6706:   Stack({
	MainConstructor,  // line #6707:     @required this.frames,
	MainConstructor,  // line #6708:     @required this.messages,
	MainConstructor,  // line #6709:     @required this.truncated,
	MainConstructor,  // line #6710:     this.asyncCausalFrames,
	MainConstructor,  // line #6711:     this.awaiterFrames,
	MainConstructor,  // line #6712:   });
	BlankLine,        // line #6713:
	NamedConstructor, // line #6714:   Stack._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #6715:     frames = List<Frame>.from(
	NamedConstructor, // line #6716:         createServiceObject(json['frames'], const ['Frame']) ?? []);
	NamedConstructor, // line #6717:     asyncCausalFrames = json['asyncCausalFrames'] == null
	NamedConstructor, // line #6718:         ? null
	NamedConstructor, // line #6719:         : List<Frame>.from(
	NamedConstructor, // line #6720:             createServiceObject(json['asyncCausalFrames'], const ['Frame']));
	NamedConstructor, // line #6721:     awaiterFrames = json['awaiterFrames'] == null
	NamedConstructor, // line #6722:         ? null
	NamedConstructor, // line #6723:         : List<Frame>.from(
	NamedConstructor, // line #6724:             createServiceObject(json['awaiterFrames'], const ['Frame']));
	NamedConstructor, // line #6725:     messages = List<Message>.from(
	NamedConstructor, // line #6726:         createServiceObject(json['messages'], const ['Message']) ?? []);
	NamedConstructor, // line #6727:     truncated = json['truncated'];
	NamedConstructor, // line #6728:   }
	BlankLine,        // line #6729:
	OverrideMethod,   // line #6730:   @override
	OverrideMethod,   // line #6731:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #6732:     var json = <String, dynamic>{};
	OverrideMethod,   // line #6733:     json['type'] = 'Stack';
	OverrideMethod,   // line #6734:     json.addAll({
	OverrideMethod,   // line #6735:       'frames': frames.map((f) => f.toJson()).toList(),
	OverrideMethod,   // line #6736:       'messages': messages.map((f) => f.toJson()).toList(),
	OverrideMethod,   // line #6737:       'truncated': truncated,
	OverrideMethod,   // line #6738:     });
	OverrideMethod,   // line #6739:     _setIfNotNull(json, 'asyncCausalFrames',
	OverrideMethod,   // line #6740:         asyncCausalFrames?.map((f) => f?.toJson())?.toList());
	OverrideMethod,   // line #6741:     _setIfNotNull(json, 'awaiterFrames',
	OverrideMethod,   // line #6742:         awaiterFrames?.map((f) => f?.toJson())?.toList());
	OverrideMethod,   // line #6743:     return json;
	OverrideMethod,   // line #6744:   }
	BlankLine,        // line #6745:
	OtherMethod,      // line #6746:   String toString() => '[Stack ' //
	OtherMethod,      // line #6747:       'type: ${type}, frames: ${frames}, messages: ${messages}, ' //
	OtherMethod,      // line #6748:       'truncated: ${truncated}]';
	BlankLine,        // line #6749:
}

var wantVMServiceClass86 = []EntityType{
	Unknown,          // line #6753: {
	OtherMethod,      // line #6754:   static Success parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #6755:       json == null ? null : Success._fromJson(json);
	BlankLine,        // line #6756:
	MainConstructor,  // line #6757:   Success();
	BlankLine,        // line #6758:
	NamedConstructor, // line #6759:   Success._fromJson(Map<String, dynamic> json) : super._fromJson(json);
	BlankLine,        // line #6760:
	OverrideMethod,   // line #6761:   @override
	OverrideMethod,   // line #6762:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #6763:     var json = <String, dynamic>{};
	OverrideMethod,   // line #6764:     json['type'] = 'Success';
	OverrideMethod,   // line #6765:     return json;
	OverrideMethod,   // line #6766:   }
	BlankLine,        // line #6767:
	OtherMethod,      // line #6768:   String toString() => '[Success type: ${type}]';
	BlankLine,        // line #6769:
}

var wantVMServiceClass87 = []EntityType{
	Unknown,          // line #6771: {
	OtherMethod,      // line #6772:   static Timeline parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #6773:       json == null ? null : Timeline._fromJson(json);
	BlankLine,        // line #6774:
	InstanceVariable, // line #6775:   /// A list of timeline events. No order is guaranteed for these events; in
	InstanceVariable, // line #6776:   /// particular, these events may be unordered with respect to their
	InstanceVariable, // line #6777:   /// timestamps.
	InstanceVariable, // line #6778:   List<TimelineEvent> traceEvents;
	BlankLine,        // line #6779:
	InstanceVariable, // line #6780:   /// The start of the period of time in which traceEvents were collected.
	InstanceVariable, // line #6781:   int timeOriginMicros;
	BlankLine,        // line #6782:
	InstanceVariable, // line #6783:   /// The duration of time covered by the timeline.
	InstanceVariable, // line #6784:   int timeExtentMicros;
	BlankLine,        // line #6785:
	MainConstructor,  // line #6786:   Timeline({
	MainConstructor,  // line #6787:     @required this.traceEvents,
	MainConstructor,  // line #6788:     @required this.timeOriginMicros,
	MainConstructor,  // line #6789:     @required this.timeExtentMicros,
	MainConstructor,  // line #6790:   });
	BlankLine,        // line #6791:
	NamedConstructor, // line #6792:   Timeline._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #6793:     traceEvents = List<TimelineEvent>.from(
	NamedConstructor, // line #6794:         createServiceObject(json['traceEvents'], const ['TimelineEvent']) ??
	NamedConstructor, // line #6795:             []);
	NamedConstructor, // line #6796:     timeOriginMicros = json['timeOriginMicros'];
	NamedConstructor, // line #6797:     timeExtentMicros = json['timeExtentMicros'];
	NamedConstructor, // line #6798:   }
	BlankLine,        // line #6799:
	OverrideMethod,   // line #6800:   @override
	OverrideMethod,   // line #6801:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #6802:     var json = <String, dynamic>{};
	OverrideMethod,   // line #6803:     json['type'] = 'Timeline';
	OverrideMethod,   // line #6804:     json.addAll({
	OverrideMethod,   // line #6805:       'traceEvents': traceEvents.map((f) => f.toJson()).toList(),
	OverrideMethod,   // line #6806:       'timeOriginMicros': timeOriginMicros,
	OverrideMethod,   // line #6807:       'timeExtentMicros': timeExtentMicros,
	OverrideMethod,   // line #6808:     });
	OverrideMethod,   // line #6809:     return json;
	OverrideMethod,   // line #6810:   }
	BlankLine,        // line #6811:
	OtherMethod,      // line #6812:   String toString() => '[Timeline ' //
	OtherMethod,      // line #6813:       'type: ${type}, traceEvents: ${traceEvents}, timeOriginMicros: ${timeOriginMicros}, ' //
	OtherMethod,      // line #6814:       'timeExtentMicros: ${timeExtentMicros}]';
	BlankLine,        // line #6815:
}

var wantVMServiceClass88 = []EntityType{
	Unknown,          // line #6819: {
	OtherMethod,      // line #6820:   static TimelineEvent parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #6821:       json == null ? null : TimelineEvent._fromJson(json);
	BlankLine,        // line #6822:
	InstanceVariable, // line #6823:   Map<String, dynamic> json;
	BlankLine,        // line #6824:
	MainConstructor,  // line #6825:   TimelineEvent();
	BlankLine,        // line #6826:
	NamedConstructor, // line #6827:   TimelineEvent._fromJson(this.json);
	BlankLine,        // line #6828:
	OtherMethod,      // line #6829:   Map<String, dynamic> toJson() {
	OtherMethod,      // line #6830:     var result = json == null ? <String, dynamic>{} : Map.of(json);
	OtherMethod,      // line #6831:     result['type'] = 'TimelineEvent';
	OtherMethod,      // line #6832:     return result;
	OtherMethod,      // line #6833:   }
	BlankLine,        // line #6834:
	OtherMethod,      // line #6835:   String toString() => '[TimelineEvent ]';
	BlankLine,        // line #6836:
}

var wantVMServiceClass89 = []EntityType{
	Unknown,          // line #6838: {
	OtherMethod,      // line #6839:   static TimelineFlags parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #6840:       json == null ? null : TimelineFlags._fromJson(json);
	BlankLine,        // line #6841:
	InstanceVariable, // line #6842:   /// The name of the recorder currently in use. Recorder types include, but are
	InstanceVariable, // line #6843:   /// not limited to: Callback, Endless, Fuchsia, Macos, Ring, Startup, and
	InstanceVariable, // line #6844:   /// Systrace. Set to "null" if no recorder is currently set.
	InstanceVariable, // line #6845:   String recorderName;
	BlankLine,        // line #6846:
	InstanceVariable, // line #6847:   /// The list of all available timeline streams.
	InstanceVariable, // line #6848:   List<String> availableStreams;
	BlankLine,        // line #6849:
	InstanceVariable, // line #6850:   /// The list of timeline streams that are currently enabled.
	InstanceVariable, // line #6851:   List<String> recordedStreams;
	BlankLine,        // line #6852:
	MainConstructor,  // line #6853:   TimelineFlags({
	MainConstructor,  // line #6854:     @required this.recorderName,
	MainConstructor,  // line #6855:     @required this.availableStreams,
	MainConstructor,  // line #6856:     @required this.recordedStreams,
	MainConstructor,  // line #6857:   });
	BlankLine,        // line #6858:
	NamedConstructor, // line #6859:   TimelineFlags._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #6860:     recorderName = json['recorderName'];
	NamedConstructor, // line #6861:     availableStreams = List<String>.from(json['availableStreams']);
	NamedConstructor, // line #6862:     recordedStreams = List<String>.from(json['recordedStreams']);
	NamedConstructor, // line #6863:   }
	BlankLine,        // line #6864:
	OverrideMethod,   // line #6865:   @override
	OverrideMethod,   // line #6866:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #6867:     var json = <String, dynamic>{};
	OverrideMethod,   // line #6868:     json['type'] = 'TimelineFlags';
	OverrideMethod,   // line #6869:     json.addAll({
	OverrideMethod,   // line #6870:       'recorderName': recorderName,
	OverrideMethod,   // line #6871:       'availableStreams': availableStreams.map((f) => f).toList(),
	OverrideMethod,   // line #6872:       'recordedStreams': recordedStreams.map((f) => f).toList(),
	OverrideMethod,   // line #6873:     });
	OverrideMethod,   // line #6874:     return json;
	OverrideMethod,   // line #6875:   }
	BlankLine,        // line #6876:
	OtherMethod,      // line #6877:   String toString() => '[TimelineFlags ' //
	OtherMethod,      // line #6878:       'type: ${type}, recorderName: ${recorderName}, availableStreams: ${availableStreams}, ' //
	OtherMethod,      // line #6879:       'recordedStreams: ${recordedStreams}]';
	BlankLine,        // line #6880:
}

var wantVMServiceClass90 = []EntityType{
	Unknown,          // line #6882: {
	OtherMethod,      // line #6883:   static Timestamp parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #6884:       json == null ? null : Timestamp._fromJson(json);
	BlankLine,        // line #6885:
	InstanceVariable, // line #6886:   /// A timestamp in microseconds since epoch.
	InstanceVariable, // line #6887:   int timestamp;
	BlankLine,        // line #6888:
	MainConstructor,  // line #6889:   Timestamp({
	MainConstructor,  // line #6890:     @required this.timestamp,
	MainConstructor,  // line #6891:   });
	BlankLine,        // line #6892:
	NamedConstructor, // line #6893:   Timestamp._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #6894:     timestamp = json['timestamp'];
	NamedConstructor, // line #6895:   }
	BlankLine,        // line #6896:
	OverrideMethod,   // line #6897:   @override
	OverrideMethod,   // line #6898:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #6899:     var json = <String, dynamic>{};
	OverrideMethod,   // line #6900:     json['type'] = 'Timestamp';
	OverrideMethod,   // line #6901:     json.addAll({
	OverrideMethod,   // line #6902:       'timestamp': timestamp,
	OverrideMethod,   // line #6903:     });
	OverrideMethod,   // line #6904:     return json;
	OverrideMethod,   // line #6905:   }
	BlankLine,        // line #6906:
	OtherMethod,      // line #6907:   String toString() => '[Timestamp type: ${type}, timestamp: ${timestamp}]';
	BlankLine,        // line #6908:
}

var wantVMServiceClass91 = []EntityType{
	Unknown,          // line #6911: {
	OtherMethod,      // line #6912:   static TypeArgumentsRef parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #6913:       json == null ? null : TypeArgumentsRef._fromJson(json);
	BlankLine,        // line #6914:
	InstanceVariable, // line #6915:   /// A name for this type argument list.
	InstanceVariable, // line #6916:   String name;
	BlankLine,        // line #6917:
	MainConstructor,  // line #6918:   TypeArgumentsRef({
	MainConstructor,  // line #6919:     @required this.name,
	MainConstructor,  // line #6920:     @required String id,
	MainConstructor,  // line #6921:   }) : super(id: id);
	BlankLine,        // line #6922:
	NamedConstructor, // line #6923:   TypeArgumentsRef._fromJson(Map<String, dynamic> json)
	NamedConstructor, // line #6924:       : super._fromJson(json) {
	NamedConstructor, // line #6925:     name = json['name'];
	NamedConstructor, // line #6926:   }
	BlankLine,        // line #6927:
	OverrideMethod,   // line #6928:   @override
	OverrideMethod,   // line #6929:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #6930:     var json = super.toJson();
	OverrideMethod,   // line #6931:     json['type'] = '@TypeArguments';
	OverrideMethod,   // line #6932:     json.addAll({
	OverrideMethod,   // line #6933:       'name': name,
	OverrideMethod,   // line #6934:     });
	OverrideMethod,   // line #6935:     return json;
	OverrideMethod,   // line #6936:   }
	BlankLine,        // line #6937:
	OtherMethod,      // line #6938:   int get hashCode => id.hashCode;
	BlankLine,        // line #6939:
	OtherMethod,      // line #6940:   operator ==(other) => other is TypeArgumentsRef && id == other.id;
	BlankLine,        // line #6941:
	OtherMethod,      // line #6942:   String toString() =>
	OtherMethod,      // line #6943:       '[TypeArgumentsRef type: ${type}, id: ${id}, name: ${name}]';
	BlankLine,        // line #6944:
}

var wantVMServiceClass92 = []EntityType{
	Unknown,          // line #6948: {
	OtherMethod,      // line #6949:   static TypeArguments parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #6950:       json == null ? null : TypeArguments._fromJson(json);
	BlankLine,        // line #6951:
	InstanceVariable, // line #6952:   /// A name for this type argument list.
	InstanceVariable, // line #6953:   String name;
	BlankLine,        // line #6954:
	InstanceVariable, // line #6955:   /// A list of types.
	InstanceVariable, // line #6956:   ///
	InstanceVariable, // line #6957:   /// The value will always be one of the kinds: Type, TypeRef, TypeParameter,
	InstanceVariable, // line #6958:   /// BoundedType.
	InstanceVariable, // line #6959:   List<InstanceRef> types;
	BlankLine,        // line #6960:
	MainConstructor,  // line #6961:   TypeArguments({
	MainConstructor,  // line #6962:     @required this.name,
	MainConstructor,  // line #6963:     @required this.types,
	MainConstructor,  // line #6964:     @required String id,
	MainConstructor,  // line #6965:   }) : super(id: id);
	BlankLine,        // line #6966:
	NamedConstructor, // line #6967:   TypeArguments._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #6968:     name = json['name'];
	NamedConstructor, // line #6969:     types = List<InstanceRef>.from(
	NamedConstructor, // line #6970:         createServiceObject(json['types'], const ['InstanceRef']) ?? []);
	NamedConstructor, // line #6971:   }
	BlankLine,        // line #6972:
	OverrideMethod,   // line #6973:   @override
	OverrideMethod,   // line #6974:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #6975:     var json = super.toJson();
	OverrideMethod,   // line #6976:     json['type'] = 'TypeArguments';
	OverrideMethod,   // line #6977:     json.addAll({
	OverrideMethod,   // line #6978:       'name': name,
	OverrideMethod,   // line #6979:       'types': types.map((f) => f.toJson()).toList(),
	OverrideMethod,   // line #6980:     });
	OverrideMethod,   // line #6981:     return json;
	OverrideMethod,   // line #6982:   }
	BlankLine,        // line #6983:
	OtherMethod,      // line #6984:   int get hashCode => id.hashCode;
	BlankLine,        // line #6985:
	OtherMethod,      // line #6986:   operator ==(other) => other is TypeArguments && id == other.id;
	BlankLine,        // line #6987:
	OtherMethod,      // line #6988:   String toString() =>
	OtherMethod,      // line #6989:       '[TypeArguments type: ${type}, id: ${id}, name: ${name}, types: ${types}]';
	BlankLine,        // line #6990:
}

var wantVMServiceClass93 = []EntityType{
	Unknown,          // line #7002: {
	OtherMethod,      // line #7003:   static UnresolvedSourceLocation parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #7004:       json == null ? null : UnresolvedSourceLocation._fromJson(json);
	BlankLine,        // line #7005:
	InstanceVariable, // line #7006:   /// The script containing the source location if the script has been loaded.
	InstanceVariable, // line #7007:   @optional
	InstanceVariable, // line #7008:   ScriptRef script;
	BlankLine,        // line #7009:
	InstanceVariable, // line #7010:   /// The uri of the script containing the source location if the script has yet
	InstanceVariable, // line #7011:   /// to be loaded.
	InstanceVariable, // line #7012:   @optional
	InstanceVariable, // line #7013:   String scriptUri;
	BlankLine,        // line #7014:
	InstanceVariable, // line #7015:   /// An approximate token position for the source location. This may change
	InstanceVariable, // line #7016:   /// when the location is resolved.
	InstanceVariable, // line #7017:   @optional
	InstanceVariable, // line #7018:   int tokenPos;
	BlankLine,        // line #7019:
	InstanceVariable, // line #7020:   /// An approximate line number for the source location. This may change when
	InstanceVariable, // line #7021:   /// the location is resolved.
	InstanceVariable, // line #7022:   @optional
	InstanceVariable, // line #7023:   int line;
	BlankLine,        // line #7024:
	InstanceVariable, // line #7025:   /// An approximate column number for the source location. This may change when
	InstanceVariable, // line #7026:   /// the location is resolved.
	InstanceVariable, // line #7027:   @optional
	InstanceVariable, // line #7028:   int column;
	BlankLine,        // line #7029:
	MainConstructor,  // line #7030:   UnresolvedSourceLocation({
	MainConstructor,  // line #7031:     this.script,
	MainConstructor,  // line #7032:     this.scriptUri,
	MainConstructor,  // line #7033:     this.tokenPos,
	MainConstructor,  // line #7034:     this.line,
	MainConstructor,  // line #7035:     this.column,
	MainConstructor,  // line #7036:   });
	BlankLine,        // line #7037:
	NamedConstructor, // line #7038:   UnresolvedSourceLocation._fromJson(Map<String, dynamic> json)
	NamedConstructor, // line #7039:       : super._fromJson(json) {
	NamedConstructor, // line #7040:     script = createServiceObject(json['script'], const ['ScriptRef']);
	NamedConstructor, // line #7041:     scriptUri = json['scriptUri'];
	NamedConstructor, // line #7042:     tokenPos = json['tokenPos'];
	NamedConstructor, // line #7043:     line = json['line'];
	NamedConstructor, // line #7044:     column = json['column'];
	NamedConstructor, // line #7045:   }
	BlankLine,        // line #7046:
	OverrideMethod,   // line #7047:   @override
	OverrideMethod,   // line #7048:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #7049:     var json = <String, dynamic>{};
	OverrideMethod,   // line #7050:     json['type'] = 'UnresolvedSourceLocation';
	OverrideMethod,   // line #7051:     _setIfNotNull(json, 'script', script?.toJson());
	OverrideMethod,   // line #7052:     _setIfNotNull(json, 'scriptUri', scriptUri);
	OverrideMethod,   // line #7053:     _setIfNotNull(json, 'tokenPos', tokenPos);
	OverrideMethod,   // line #7054:     _setIfNotNull(json, 'line', line);
	OverrideMethod,   // line #7055:     _setIfNotNull(json, 'column', column);
	OverrideMethod,   // line #7056:     return json;
	OverrideMethod,   // line #7057:   }
	BlankLine,        // line #7058:
	OtherMethod,      // line #7059:   String toString() => '[UnresolvedSourceLocation type: ${type}]';
	BlankLine,        // line #7060:
}

var wantVMServiceClass94 = []EntityType{
	Unknown,          // line #7063: {
	OtherMethod,      // line #7064:   static Version parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #7065:       json == null ? null : Version._fromJson(json);
	BlankLine,        // line #7066:
	InstanceVariable, // line #7067:   /// The major version number is incremented when the protocol is changed in a
	InstanceVariable, // line #7068:   /// potentially incompatible way.
	InstanceVariable, // line #7069:   int major;
	BlankLine,        // line #7070:
	InstanceVariable, // line #7071:   /// The minor version number is incremented when the protocol is changed in a
	InstanceVariable, // line #7072:   /// backwards compatible way.
	InstanceVariable, // line #7073:   int minor;
	BlankLine,        // line #7074:
	MainConstructor,  // line #7075:   Version({
	MainConstructor,  // line #7076:     @required this.major,
	MainConstructor,  // line #7077:     @required this.minor,
	MainConstructor,  // line #7078:   });
	BlankLine,        // line #7079:
	NamedConstructor, // line #7080:   Version._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #7081:     major = json['major'];
	NamedConstructor, // line #7082:     minor = json['minor'];
	NamedConstructor, // line #7083:   }
	BlankLine,        // line #7084:
	OverrideMethod,   // line #7085:   @override
	OverrideMethod,   // line #7086:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #7087:     var json = <String, dynamic>{};
	OverrideMethod,   // line #7088:     json['type'] = 'Version';
	OverrideMethod,   // line #7089:     json.addAll({
	OverrideMethod,   // line #7090:       'major': major,
	OverrideMethod,   // line #7091:       'minor': minor,
	OverrideMethod,   // line #7092:     });
	OverrideMethod,   // line #7093:     return json;
	OverrideMethod,   // line #7094:   }
	BlankLine,        // line #7095:
	OtherMethod,      // line #7096:   String toString() =>
	OtherMethod,      // line #7097:       '[Version type: ${type}, major: ${major}, minor: ${minor}]';
	BlankLine,        // line #7098:
}

var wantVMServiceClass95 = []EntityType{
	Unknown,          // line #7101: {
	OtherMethod,      // line #7102:   static VMRef parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #7103:       json == null ? null : VMRef._fromJson(json);
	BlankLine,        // line #7104:
	InstanceVariable, // line #7105:   /// A name identifying this vm. Not guaranteed to be unique.
	InstanceVariable, // line #7106:   String name;
	BlankLine,        // line #7107:
	MainConstructor,  // line #7108:   VMRef({
	MainConstructor,  // line #7109:     @required this.name,
	MainConstructor,  // line #7110:   });
	BlankLine,        // line #7111:
	NamedConstructor, // line #7112:   VMRef._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #7113:     name = json['name'];
	NamedConstructor, // line #7114:   }
	BlankLine,        // line #7115:
	OverrideMethod,   // line #7116:   @override
	OverrideMethod,   // line #7117:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #7118:     var json = <String, dynamic>{};
	OverrideMethod,   // line #7119:     json['type'] = '@VM';
	OverrideMethod,   // line #7120:     json.addAll({
	OverrideMethod,   // line #7121:       'name': name,
	OverrideMethod,   // line #7122:     });
	OverrideMethod,   // line #7123:     return json;
	OverrideMethod,   // line #7124:   }
	BlankLine,        // line #7125:
	OtherMethod,      // line #7126:   String toString() => '[VMRef type: ${type}, name: ${name}]';
	BlankLine,        // line #7127:
}

var wantVMServiceClass96 = []EntityType{
	Unknown,          // line #7129: {
	OtherMethod,      // line #7130:   static VM parse(Map<String, dynamic> json) =>
	OtherMethod,      // line #7131:       json == null ? null : VM._fromJson(json);
	BlankLine,        // line #7132:
	InstanceVariable, // line #7133:   /// A name identifying this vm. Not guaranteed to be unique.
	InstanceVariable, // line #7134:   String name;
	BlankLine,        // line #7135:
	InstanceVariable, // line #7136:   /// Word length on target architecture (e.g. 32, 64).
	InstanceVariable, // line #7137:   int architectureBits;
	BlankLine,        // line #7138:
	InstanceVariable, // line #7139:   /// The CPU we are actually running on.
	InstanceVariable, // line #7140:   String hostCPU;
	BlankLine,        // line #7141:
	InstanceVariable, // line #7142:   /// The operating system we are running on.
	InstanceVariable, // line #7143:   String operatingSystem;
	BlankLine,        // line #7144:
	InstanceVariable, // line #7145:   /// The CPU we are generating code for.
	InstanceVariable, // line #7146:   String targetCPU;
	BlankLine,        // line #7147:
	InstanceVariable, // line #7148:   /// The Dart VM version string.
	InstanceVariable, // line #7149:   String version;
	BlankLine,        // line #7150:
	InstanceVariable, // line #7151:   /// The process id for the VM.
	InstanceVariable, // line #7152:   int pid;
	BlankLine,        // line #7153:
	InstanceVariable, // line #7154:   /// The time that the VM started in milliseconds since the epoch.
	InstanceVariable, // line #7155:   ///
	InstanceVariable, // line #7156:   /// Suitable to pass to DateTime.fromMillisecondsSinceEpoch.
	InstanceVariable, // line #7157:   int startTime;
	BlankLine,        // line #7158:
	InstanceVariable, // line #7159:   /// A list of isolates running in the VM.
	InstanceVariable, // line #7160:   List<IsolateRef> isolates;
	BlankLine,        // line #7161:
	InstanceVariable, // line #7162:   /// A list of isolate groups running in the VM.
	InstanceVariable, // line #7163:   List<IsolateGroupRef> isolateGroups;
	BlankLine,        // line #7164:
	InstanceVariable, // line #7165:   /// A list of system isolates running in the VM.
	InstanceVariable, // line #7166:   List<IsolateRef> systemIsolates;
	BlankLine,        // line #7167:
	InstanceVariable, // line #7168:   /// A list of isolate groups which contain system isolates running in the VM.
	InstanceVariable, // line #7169:   List<IsolateGroupRef> systemIsolateGroups;
	BlankLine,        // line #7170:
	MainConstructor,  // line #7171:   VM({
	MainConstructor,  // line #7172:     @required this.name,
	MainConstructor,  // line #7173:     @required this.architectureBits,
	MainConstructor,  // line #7174:     @required this.hostCPU,
	MainConstructor,  // line #7175:     @required this.operatingSystem,
	MainConstructor,  // line #7176:     @required this.targetCPU,
	MainConstructor,  // line #7177:     @required this.version,
	MainConstructor,  // line #7178:     @required this.pid,
	MainConstructor,  // line #7179:     @required this.startTime,
	MainConstructor,  // line #7180:     @required this.isolates,
	MainConstructor,  // line #7181:     @required this.isolateGroups,
	MainConstructor,  // line #7182:     @required this.systemIsolates,
	MainConstructor,  // line #7183:     @required this.systemIsolateGroups,
	MainConstructor,  // line #7184:   });
	BlankLine,        // line #7185:
	NamedConstructor, // line #7186:   VM._fromJson(Map<String, dynamic> json) : super._fromJson(json) {
	NamedConstructor, // line #7187:     name = json['name'];
	NamedConstructor, // line #7188:     architectureBits = json['architectureBits'];
	NamedConstructor, // line #7189:     hostCPU = json['hostCPU'];
	NamedConstructor, // line #7190:     operatingSystem = json['operatingSystem'];
	NamedConstructor, // line #7191:     targetCPU = json['targetCPU'];
	NamedConstructor, // line #7192:     version = json['version'];
	NamedConstructor, // line #7193:     pid = json['pid'];
	NamedConstructor, // line #7194:     startTime = json['startTime'];
	NamedConstructor, // line #7195:     isolates = List<IsolateRef>.from(
	NamedConstructor, // line #7196:         createServiceObject(json['isolates'], const ['IsolateRef']) ?? []);
	NamedConstructor, // line #7197:     isolateGroups = List<IsolateGroupRef>.from(
	NamedConstructor, // line #7198:         createServiceObject(json['isolateGroups'], const ['IsolateGroupRef']) ??
	NamedConstructor, // line #7199:             []);
	NamedConstructor, // line #7200:     systemIsolates = List<IsolateRef>.from(
	NamedConstructor, // line #7201:         createServiceObject(json['systemIsolates'], const ['IsolateRef']) ??
	NamedConstructor, // line #7202:             []);
	NamedConstructor, // line #7203:     systemIsolateGroups = List<IsolateGroupRef>.from(createServiceObject(
	NamedConstructor, // line #7204:             json['systemIsolateGroups'], const ['IsolateGroupRef']) ??
	NamedConstructor, // line #7205:         []);
	NamedConstructor, // line #7206:   }
	BlankLine,        // line #7207:
	OverrideMethod,   // line #7208:   @override
	OverrideMethod,   // line #7209:   Map<String, dynamic> toJson() {
	OverrideMethod,   // line #7210:     var json = <String, dynamic>{};
	OverrideMethod,   // line #7211:     json['type'] = 'VM';
	OverrideMethod,   // line #7212:     json.addAll({
	OverrideMethod,   // line #7213:       'name': name,
	OverrideMethod,   // line #7214:       'architectureBits': architectureBits,
	OverrideMethod,   // line #7215:       'hostCPU': hostCPU,
	OverrideMethod,   // line #7216:       'operatingSystem': operatingSystem,
	OverrideMethod,   // line #7217:       'targetCPU': targetCPU,
	OverrideMethod,   // line #7218:       'version': version,
	OverrideMethod,   // line #7219:       'pid': pid,
	OverrideMethod,   // line #7220:       'startTime': startTime,
	OverrideMethod,   // line #7221:       'isolates': isolates.map((f) => f.toJson()).toList(),
	OverrideMethod,   // line #7222:       'isolateGroups': isolateGroups.map((f) => f.toJson()).toList(),
	OverrideMethod,   // line #7223:       'systemIsolates': systemIsolates.map((f) => f.toJson()).toList(),
	OverrideMethod,   // line #7224:       'systemIsolateGroups':
	OverrideMethod,   // line #7225:           systemIsolateGroups.map((f) => f.toJson()).toList(),
	OverrideMethod,   // line #7226:     });
	OverrideMethod,   // line #7227:     return json;
	OverrideMethod,   // line #7228:   }
	BlankLine,        // line #7229:
	OtherMethod,      // line #7230:   String toString() => '[VM]';
	BlankLine,        // line #7231:
}

var wantVMServiceClassAll = [][]EntityType{
	wantVMServiceClass01,
	wantVMServiceClass02,
	wantVMServiceClass03,
	wantVMServiceClass04,
	wantVMServiceClass05,
	wantVMServiceClass06,
	wantVMServiceClass07,
	wantVMServiceClass08,
	wantVMServiceClass09,
	wantVMServiceClass10,
	wantVMServiceClass11,
	wantVMServiceClass12,
	wantVMServiceClass13,
	wantVMServiceClass14,
	wantVMServiceClass15,
	wantVMServiceClass16,
	wantVMServiceClass17,
	wantVMServiceClass18,
	wantVMServiceClass19,
	wantVMServiceClass20,
	wantVMServiceClass21,
	wantVMServiceClass22,
	wantVMServiceClass23,
	wantVMServiceClass24,
	wantVMServiceClass25,
	wantVMServiceClass26,
	wantVMServiceClass27,
	wantVMServiceClass28,
	wantVMServiceClass29,
	wantVMServiceClass30,
	wantVMServiceClass31,
	wantVMServiceClass32,
	wantVMServiceClass33,
	wantVMServiceClass34,
	wantVMServiceClass35,
	wantVMServiceClass36,
	wantVMServiceClass37,
	wantVMServiceClass38,
	wantVMServiceClass39,
	wantVMServiceClass40,
	wantVMServiceClass41,
	wantVMServiceClass42,
	wantVMServiceClass43,
	wantVMServiceClass44,
	wantVMServiceClass45,
	wantVMServiceClass46,
	wantVMServiceClass47,
	wantVMServiceClass48,
	wantVMServiceClass49,
	wantVMServiceClass50,
	wantVMServiceClass51,
	wantVMServiceClass52,
	wantVMServiceClass53,
	wantVMServiceClass54,
	wantVMServiceClass55,
	wantVMServiceClass56,
	wantVMServiceClass57,
	wantVMServiceClass58,
	wantVMServiceClass59,
	wantVMServiceClass60,
	wantVMServiceClass61,
	wantVMServiceClass62,
	wantVMServiceClass63,
	wantVMServiceClass64,
	wantVMServiceClass65,
	wantVMServiceClass66,
	wantVMServiceClass67,
	wantVMServiceClass68,
	wantVMServiceClass69,
	wantVMServiceClass70,
	wantVMServiceClass71,
	wantVMServiceClass72,
	wantVMServiceClass73,
	wantVMServiceClass74,
	wantVMServiceClass75,
	wantVMServiceClass76,
	wantVMServiceClass77,
	wantVMServiceClass78,
	wantVMServiceClass79,
	wantVMServiceClass80,
	wantVMServiceClass81,
	wantVMServiceClass82,
	wantVMServiceClass83,
	wantVMServiceClass84,
	wantVMServiceClass85,
	wantVMServiceClass86,
	wantVMServiceClass87,
	wantVMServiceClass88,
	wantVMServiceClass89,
	wantVMServiceClass90,
	wantVMServiceClass91,
	wantVMServiceClass92,
	wantVMServiceClass93,
	wantVMServiceClass94,
	wantVMServiceClass95,
	wantVMServiceClass96,
}
