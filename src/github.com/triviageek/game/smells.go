package game

// see http://industriallogic.com/wp-content/uploads/2005/09/smellstorefactorings.pdf
var smells = []Smell{
	{`Alternative Classes with Different Interfaces`,
		`occurs when the interfaces of two classes
are different and yet the classes are quite similar. If you can find the similarities between
the two classes, you can often refactor the classes to make them share a common
interface`},
	{`Combinatorial Explosion`,
		`A subtle form of duplication, this smell exists when numerous
pieces of code do the same thing using different combinations of data or behavior`},
	{`Comments (a.k.a. Deodorant)`,
		` When you feel like writing a comment, first try "to refactor so that the comment becomes superfluous"`},
	{`Conditional Complexity`,
		`Conditional logic is innocent in its infancy, when it’s simple to
understand and contained within a few lines of code. Unfortunately, it rarely ages well.`},
	{`Data Class`,
		`Classes that have fields, getting and setting methods for the fields, and
nothing else. Such classes are dumb data holders and are almost certainly being
manipulated in far too much detail by other classes`},
	{`Data Clumps`,
		`Bunches of data that that hang around together really ought to be made
into their own object. A good test is to consider deleting one of the data values: if you did
this, would the others make any sense? If they don't, it's a sure sign that you have an
object that's dying to be born.`},
	{`Divergent Change`,
		`Occurs when one class is commonly changed in different ways for 
different reasons. Separating these divergent responsibilities decreases the chance that
one change could affect another and lower maintenance costs. `},
	{`Duplicated Code`,
		`The most pervasive and pungent smell in software.
It tends to be either explicit or subtle. This smell can exists in structures or processing steps 
that are outwardly different, yet essentially the same`},
	{`Feature Envy`,
		`Data and behavior that acts on that data belong together. When a method
makes too many calls to other classes to obtain data or functionality, this smell is in
the air. `},
	{`Freeloader (a.k.a. Lazy Class)`,
		`A class that isn’t doing enough to pay for itself should be eliminated`},
	{`Inappropriate Intimacy`,
		`Sometimes classes become far too close and spend too
much time delving into each others’ private parts. We may not be prudes when it comes
to people, but we think our classes should follow strict, puritan rules. Those
classes need to be broken up as lovers were in ancient days`},
	{`Incomplete Library Class`,
		`Occurs when responsibilities emerge in our code that clearly
should be moved to a foreign class, but we are unable or unwilling to modify the
class to accept these new responsibilities.`},
	{`Indecent Exposure`,
		`The smell occurs when methods or classes that ought
not to be visible to clients are publicly visible to them. Exposing such code means that
clients know about code that is unimportant or only indirectly important. This contributes
to the complexity of a design`},
	{`Large Class`,
		` Fowler and Beck note that the presence of too many instance variables
usually indicates that a class is trying to do too much. In general, those classes typically
contain too many responsibilities`},
	{`Long Method`,
		`Two long methods may very well contain duplicated code. Yet if you
break those methods into smaller methods, you can often find ways for the two to share
logic. Fowler and Beck also describe how small methods help explain code. If you don’t
understand what a chunk of code does and you extract that code to a small, well-named
method, it will be easier to understand the original code. `},
	{`Long Parameter List`,
		`Long lists of parameters in a method, though common in
procedural code, are difficult to understand and likely to be volatile. Consider which objects
this method really needs to do its job`},
	{`Message Chains`,
		`Occur when you see a long sequence of method calls or temporary
variables to get some data. This chain makes the code dependent on the relationships
between many potentially unrelated objects`},
	{`Middle Man`,
		` Delegation is good, and one of the key fundamental features of objects. But
too much of a good thing can lead to objects that add no value, simply passing messages
on to another object`},
	{`Oddball Solution`,
		` When a problem is solved one way throughout a system and the same
problem is solved another way in the same system, one of the solutions is an
inconsistent solution. The presence of this smell usually indicates subtly duplicated code`},
	{`Parallel Inheritance Hierarchies`,
		` This is really a special case of Shotgun Surgery - every
time you make a subclass of one class, you have to make a subclass of another`},
	{`Primitive Obsession`,
		`Integers, Strings, doubles, arrays and
other low-level language elements, are generic because many people use them. Classes,
on the other hand, may be as specific as you need them to be, since you create them for
specific purposes. In many cases, classes provide a simpler and more natural way to
model things than primitives. `},
	{`Refused Bequest`,
		` This smell results from inheriting code you don't want. Instead of
tolerating the inheritance, you write code to refuse the code -- which leads to ugly,
confusing code, to say the least.`},
	{`Shotgun Surgery`,
		` This smell is evident when you must change lots of pieces of code in
different places simply to add a new or extended piece of behavior`},
	{`Solution Sprawl`,
		`When code and/or data used in performing a responsibility becomes
spread across numerous classes. This smell often results
from quickly adding a feature to a system without spending enough time simplifying and
consolidating the design to best accommodate the feature.`},
	{`Speculative Generality`,
		`This odor exists when you have generic or abstract code that
isn’t actually needed today. Such code often exists to support future behavior, which may
or may not be necessary in the future`},
	{`Switch Statement`,
		`This smell exists when the same switch statement (or “if…else if…else
if” statement) is duplicated across a system. Such duplicated code reveals a lack of objectorientation
and a missed opportunity to rely on the elegance of polymorphism`},
	{`Temporary Field`,
		` Objects sometimes contain fields that don't seem to be needed all the
time. The rest of the time, the field is empty or contains irrelevant data, which is difficult to
understand. This is often an alternative to Long Parameter List`},
}
