objects = $(shell ls *.rs)
dirs = $(shell ls -d */ | sed -E "s/\/$$//g")

$(objects): FORCE
	@rustc $@
	@./$(@:.rs=)
	@rm $(@:.rs=)

$(dirs): FORCE
	cd $@; cargo run

FORCE: ;